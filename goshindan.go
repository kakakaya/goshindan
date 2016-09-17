package main

import (
	"fmt"
	"os"

	"github.com/kakakaya/goshindan/utils"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goshindan(誤診断)"
	app.Usage = "shindanmaker.comの診断結果を取得するツールです。"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "shindanID, s",
			Value: 0,
			Usage: "set shindanmaker/<shindanID>.",
		},
		cli.StringFlag{
			Name:  "username, u",
			Value: "Gopher a.k.a SandBag",
			Usage: "診断したい名前を入れて下さい",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Int("shindanID") <= 0 {
			panic(fmt.Sprintf("shindanIDの値が不正です: %d\n", c.Int("shindanID")))
		}
		result, _ := utils.Shindan(c.Int("shindanID"), c.String("username"))
		fmt.Println(result)
		return nil
	}

	app.Run(os.Args)
}
