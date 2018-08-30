package tokens

import (
	"encoding/json"
	"errors"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/satooon/fcm-go/app"
	"github.com/urfave/cli"
	"log"
	"sync"
)

type tokens struct{}

func NewTokens() *tokens {
	return &tokens{}
}

func (t *tokens) Command() cli.Command {
	return cli.Command{
		Name:  "tokens",
		Usage: "Send tokens",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "tokens",
				Usage: "Registration tokens. (example: --tokens XXXX --tokens YYYY --tokens ZZZZ)",
			},
			cli.StringFlag{
				Name:  "title",
				Usage: "Notification title",
			},
			cli.StringFlag{
				Name:  "body",
				Usage: "Notification body",
			},
			cli.IntFlag{
				Name:  "max-conn",
				Usage: "Notification max connection",
				Value: 5,
			},
		},
		Action: t.action,
	}
}

func (t *tokens) action(c *cli.Context) error {
	app, err := app.NewApp(c.GlobalString("config"))
	if err != nil {
		return err
	}
	client, err := app.MessagingClient()
	if err != nil {
		return err
	}

	tokens := c.StringSlice("tokens")
	sm := new(sync.Map)

	ch := make(chan int, c.Int("max-conn"))
	wg := sync.WaitGroup{}

	for _, t := range tokens {
		ch <- 1
		wg.Add(1)
		go func(s string) {
			defer func() {
				<-ch
				wg.Done()
			}()

			msg := &messaging.Message{
				Token: s,
				Notification: &messaging.Notification{
					Title: c.String("title"),
					Body:  c.String("body"),
				},
			}
			res, err := client.Send(app.Context(), msg, c.GlobalBool("dry-run"))
			sm.Store(s, map[string]interface{}{
				"response": res,
				"error":    err,
			})
		}(t)
	}
	wg.Wait()

	res := map[string]map[string]interface{}{}
	is_error := false
	sm.Range(func(k interface{}, v interface{}) bool {
		m := v.(map[string]interface{})
		res[k.(string)] = m
		if m["error"] != nil {
			is_error = true
		}
		return true
	})
	j, _ := json.MarshalIndent(res, "", "  ")
	o := fmt.Sprintf("Finish sent tokens message: \n%v\n", (string)(j))
	if is_error {
		return errors.New(o)
	}
	log.Print(o)
	return nil
}
