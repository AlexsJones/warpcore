package main

import (
  "github.com/urfave/cli"
  "os"
  "fmt"
  "sort"
  "github.com/AlexsJones/warpcore/core"
)

func main() {

  context := core.NewContext()

  app := cli.NewApp()
  app.Name = "Warpcore"

  app.Commands = []cli.Command{
    {
      Name:    "inspect",
      Aliases: []string{"i"},
      Usage:   "Inspect terraform",
      Action:  func(c *cli.Context) error {
        fmt.Println("...")
        return nil
      },
    },
    {
      Name:    "engage",
      Aliases: []string{"e"},
      Usage:   "Test terraform deployment",
      Action:  func(c *cli.Context) error {
        fmt.Println("...")
        return nil
      },
    },
  }

 sort.Sort(cli.CommandsByName(app.Commands))

 app.Run(os.Args)
}
