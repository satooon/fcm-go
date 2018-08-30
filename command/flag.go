package command

import "github.com/urfave/cli"

func NewFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "Firebase service account key json path",
		},
		cli.BoolFlag{
			Name: "dry-run",
		},
	}
}
