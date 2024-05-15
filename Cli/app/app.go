package app

import "github.com/urfave/cli"

// Returns the Cli Application ready to be executed
func Generate() (app *cli.App) {

	app = cli.NewApp()
	app.Name = "CLI Aplication"
	app.Usage = "Search IPs and names of Ethernet Servers"

	return
}
