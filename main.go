package main

import (
	"github.com/codegangsta/cli"
	"github.com/maleck13/testapp/cmd"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "testapp"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}

