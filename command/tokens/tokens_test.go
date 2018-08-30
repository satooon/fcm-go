package tokens_test

import (
	"fmt"
	"github.com/satooon/fcm-go/command"
	"github.com/satooon/fcm-go/command/tokens"
	"github.com/urfave/cli"
	"os"
	"testing"
)

var (
	project_root = fmt.Sprintf("%s/src/github.com/satooon/fcm-go", os.Getenv("GOPATH"))
	test_token   = "111fqAfhqV6JaI:APA91bGp-k6kf_L7qKe4svSg3U_rYMWoUpPSN0WPiSskKgP7ASn2_DZbTlI7ogD4AbnlBkR_rv4pe8Ek3RmrmXO0M3Xt3Mk24oSnedPmEnxnj2xWlIIV98ziB8dtk40jaAtrwmLTn_AR"
)

func TestTokensAction1(t *testing.T) {
	app := cli.NewApp()
	app.Flags = command.NewFlags()
	app.Commands = []cli.Command{
		tokens.NewTokens().Command(),
	}

	args := []string{
		"foo",
		"-c", fmt.Sprintf("%s/serviceAccountKey.json", project_root),
		"--dry-run",
		"tokens",
		"--title", "title test",
		"--body", "body test",
		"--tokens", "test",
	}

	if err := app.Run(args); err == nil {
		t.Fatal("expected to receive error from Run, got none")
	}
}

func TestTokensAction2(t *testing.T) {
	app := cli.NewApp()
	app.Flags = command.NewFlags()
	app.Commands = []cli.Command{
		tokens.NewTokens().Command(),
	}

	args := []string{
		"foo",
		"-c", fmt.Sprintf("%s/serviceAccountKey.json", project_root),
		"--dry-run",
		"tokens",
		"--title", "title test",
		"--body", "body test",
		"--tokens", test_token,
	}

	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}
}
