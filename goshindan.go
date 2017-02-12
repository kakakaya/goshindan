package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kakakaya/goshindan/utils"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Authors = []cli.Author{cli.Author{
		Name: "kakakaya", Email: "kakakaya+git@gmail.com",
	}}
	app.Name = "goshindan(御診断)"
	app.Usage = "shindanmaker.comの診断結果を取得したりするツール。"
	app.Version = "2.0.0"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "shindan",
			Aliases: []string{"s", "sd", "診断"},
			Usage:   "診断メーカーでの診断を行い、結果を印字する。",
			Description: "shindanmaker.comへのアクセスを行い、結果を印字する" +
				"もし診断IDやユーザー名が指定されない場合、デフォルト値で実行される。(see help.)",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "shindan-id, s",
					Value: 509717, // ねどこ
					Usage: "診断メーカーの診断IDを指定する。 もし <https://shindanmaker.com/509717> を使いたい場合、" +
						"509717を指定する",
				},
				cli.StringFlag{
					Name:  "username, u",
					Value: "Gopher a.k.a SandBag", // デフォルト名
					Usage: "診断に使うユーザー名を指定する。",
				},
				cli.BoolFlag{
					Name:  "append-url, add-url",
					Usage: "このオプションを指定すると、診断に使ったURLを末尾に付加して印字する。",
				},
			},

			Action: func(c *cli.Context) error {
				var shindanId = c.Int("shindan-id")
				if shindanId <= 0 {
					return fmt.Errorf("shindanIDの値が不正です: %d\n", shindanId)
				}
				result, _ := utils.Shindan(shindanId, c.String("username"))
				if c.Bool("append-url") {
					if strings.Index(result, "\n") == -1 {
						result = fmt.Sprintf("%s https://shindanmaker.com/%d", result, shindanId)
					} else {
						result = fmt.Sprintf("%s\nhttps://shindanmaker.com/%d\n", result, shindanId)
					}
				}
				fmt.Println(result)
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	app.Run(os.Args)
}
