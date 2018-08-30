package app

import (
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type App struct {
	app *firebase.App
	opt option.ClientOption
	ctx context.Context
	cfg *firebase.Config
}

func NewApp(config string) (*App, error) {
	opt := option.WithCredentialsFile(config)
	ctx := context.Background()
	cfg := &firebase.Config{}
	app, err := firebase.NewApp(ctx, cfg, opt)
	if err != nil {
		return nil, err
	}
	return &App{
		app: app,
		opt: opt,
		ctx: ctx,
		cfg: cfg,
	}, err
}

func (a *App) Firebase() *firebase.App {
	return a.app
}

func (a *App) ClientOption() option.ClientOption {
	return a.opt
}

func (a *App) Context() context.Context {
	return a.ctx
}

func (a *App) Config() *firebase.Config {
	return a.cfg
}

func (a *App) MessagingClient() (*MessagingClient, error) {
	c, err := a.app.Messaging(a.ctx)
	if err != nil {
		return nil, err
	}
	return &MessagingClient{c}, nil
}

type MessagingClient struct {
	*messaging.Client
}

func (c *MessagingClient) Send(ctx context.Context, msg *messaging.Message, dry bool) (string, error) {
	if dry {
		return c.Client.SendDryRun(ctx, msg)
	} else {
		return c.Client.Send(ctx, msg)
	}
}
