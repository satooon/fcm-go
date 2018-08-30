package topic_test

import (
	"fmt"
	"github.com/satooon/fcm-go/command"
	"github.com/satooon/fcm-go/command/topic"
	"github.com/urfave/cli"
	"os"
	"testing"
)

var (
	project_root = fmt.Sprintf("%s/src/github.com/satooon/fcm-go", os.Getenv("GOPATH"))
)

func TestTopicAction1(t *testing.T) {
	app := cli.NewApp()
	app.Flags = command.NewFlags()
	app.Commands = []cli.Command{
		topic.NewTopic().Command(),
	}

	args := []string{
		"foo",
		"-c", fmt.Sprintf("%s/serviceAccountKey.json", project_root),
		"--dry-run",
		"topic",
		"--title", "title test",
		"--body", "body test",
		"--name", "example",
	}

	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}
}
