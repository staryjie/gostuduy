package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var language string
	var recusive bool

	app := cli.NewApp()
	app.Name = "opless"
	app.Version = "0.1.1"
	app.Author = "StaryJie"

	app.Flags = []cli.Flag{  // 参数选项处理
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "select an language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name: "recusive, r",
			Usage: "recusive for the greeting",
			Destination: &recusive,
		},
	}
	app.Action = func(c *cli.Context) error {
		var cmd string

		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println("cmd is", cmd)
		}

		fmt.Println("language is", language)
		fmt.Println("recusive is", recusive)
		return nil
	}
	app.Run(os.Args)
}