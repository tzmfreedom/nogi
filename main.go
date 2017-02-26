package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

type Member struct {
	Name          string `toml:"name"`
	Hiragana      string `toml:"hiragana"`
	BloodType     string `toml:"blood_type"`
	Constellation string `toml:"constellation"`
	Height        string `toml:"height"`
	Birthday      string `toml:"birthday"`
}

type MemberFile struct {
	Members []*Member `toml:"members"`
}

const (
	APP_VERSION = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "nogi"

	app.Usage = "Nogizaka 46 Command Line Interface"
	app.Version = APP_VERSION
	app.Commands = []cli.Command{
		{
			Name:    "members",
			Aliases: []string{"m"},
			Usage:   "List Member",
			Flags:   []cli.Flag{},
			Action: func(ctx *cli.Context) error {
				buf, err := ioutil.ReadFile("data/members.toml")
				if err != nil {
					return err
				}
				file := &MemberFile{}
				err = toml.Unmarshal(buf, file)
				for _, member := range file.Members {
					attrs := []string{
						member.Name,
						member.Hiragana,
						member.BloodType,
						member.Birthday,
						member.Constellation,
						member.Height,
					}
					fmt.Println(strings.Join(attrs, "\t"))
				}
				return err
			},
		},
	}
	app.Run(os.Args)
}
