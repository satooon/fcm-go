package topic

import (
	"firebase.google.com/go/messaging"
	"github.com/satooon/fcm-go/app"
	"github.com/urfave/cli"
	"log"
)

type topic struct{}

func NewTopic() *topic {
	return &topic{}
}

func (t *topic) Command() cli.Command {
	return cli.Command{
		Name:  "topic",
		Usage: "Send topic",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "Topic name",
			},
			cli.StringFlag{
				Name:  "title",
				Usage: "Notification title",
			},
			cli.StringFlag{
				Name:  "body",
				Usage: "Notification body",
			},
		},
		Action: t.action,
	}
}

func (t *topic) action(c *cli.Context) error {
	app, err := app.NewApp(c.GlobalString("config"))
	if err != nil {
		return err
	}
	client, err := app.MessagingClient()
	if err != nil {
		return err
	}

	msg := &messaging.Message{
		Topic: c.String("name"),
		Notification: &messaging.Notification{
			Title: c.String("title"),
			Body:  c.String("body"),
		},
	}

	res, err := client.Send(app.Context(), msg, c.GlobalBool("dry-run"))
	if err != nil {
		return err
	}
	log.Printf("Successfully sent topic message: %v\n", res)
	return nil
}
