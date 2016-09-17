package main

import (
	"fmt"
	"os"
	"strings"

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
		cli.BoolFlag{
			Name:  "add-url",
			Usage: "末尾に診断メーカーのURLを付加して出力する。",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Int("shindanID") <= 0 {
			panic(fmt.Sprintf("shindanIDの値が不正です: %d\n", c.Int("shindanID")))
		}
		result, _ := utils.Shindan(c.Int("shindanID"), c.String("username"))
		if c.Bool("add-url") {
			if strings.Index(result, "\n") == -1 {
				result = fmt.Sprintf("%s https://shindanmaker.com/%d", result, c.Int("shindanID"))
			} else {
				result = fmt.Sprintf("%s\nhttps://shindanmaker.com/%d\n", result, c.Int("shindanID"))
			}
		}
		fmt.Println(result)
		return nil
	}

	app.Run(os.Args)
}
