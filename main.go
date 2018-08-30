package main

import (
	"github.com/satooon/fcm-go/command"
	"github.com/satooon/fcm-go/command/tokens"
	"github.com/satooon/fcm-go/command/topic"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	name    = "fcm-go"
	author  = "satooon"
	version = "0.0.1"
	usage   = "Firebase Cloud Messaging send tool"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Author = author
	app.Version = version
	app.Usage = usage

	app.Flags = command.NewFlags()

	app.Commands = []cli.Command{
		topic.NewTopic().Command(),
		tokens.NewTokens().Command(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
