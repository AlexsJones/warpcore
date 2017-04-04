package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/AlexsJones/warpcore/core"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Warpcore"

	var provider string
	var filepath string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "provider",
			Usage:       "providers supported: aws|google",
			Destination: &provider,
		},
		cli.StringFlag{
			Name:        "path",
			Usage:       "filepath for gtf files",
			Destination: &filepath,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("provider") == "aws" {
			log.Println("Initalising aws provider")
		} else if c.String("provider") == "google" {
			log.Println("Initalising google cloud provider")
		} else if len(c.String("provider")) == 0 {
			fmt.Println("Please set a provider e.g. aws")
			os.Exit(1)
		} else if len(c.String("provider")) >= 0 {
			fmt.Println("Unsupported provider")
			os.Exit(1)
		}
		return nil
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

	fmt.Println("Using filepath:" + filepath)

	_ = core.NewContext(filepath)
}
