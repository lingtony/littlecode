package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

var (
	serverChoice int
	err          error
)

func main() {
	app := cli.NewApp()
	app.Name = "gossh"
	app.Usage = "for ssh to servers"
	app.Action = func(c *cli.Context) error {
		fmt.Println("1.jumpserver\n2.ansible")
		fmt.Scanf("%d", &serverChoice)
		switch serverChoice {
		case 1:
			cmd := exec.Command("/usr/bin/ssh", "-p 34185", "-a", "42.62.69.161")
			_, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}

		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
