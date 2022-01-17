package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "great"
	app.Usage = "fight the loneliness"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello Friend!")
		return nil
	}
	app.Run(os.Args)
}

// go build -o great main.go
// ./great --help
// ./great --version
