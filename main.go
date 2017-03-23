package main

import (
	"fmt"
	"github.com/AlexsJones/warpcore/core"
	"github.com/urfave/cli"
	"os"
	"sort"
)

func main() {

	_ = core.NewContext()

	app := cli.NewApp()
	app.Name = "Warpcore"

	app.Commands = []cli.Command{
		{
			Name:    "inspect",
			Aliases: []string{"i"},
			Usage:   "Inspect terraform",
			Action: func(c *cli.Context) error {
				fmt.Println("...")
				return nil
			},
		},
		{
			Name:    "engage",
			Aliases: []string{"e"},
			Usage:   "Test terraform deployment",
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
