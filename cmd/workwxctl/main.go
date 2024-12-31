package main

import (
	"os"

	"github.com/jmek087/go-workwx-client/v2/cmd/workwxctl/commands"
)

func main() {
	app := commands.InitApp()
	// ignore errors
	_ = app.Run(os.Args)
}
