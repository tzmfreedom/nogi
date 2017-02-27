package main

import (
	"fmt"
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

type Song struct {
	No   int    `toml:"no"`
	Name string `toml:"name"`
}

type SongFile struct {
	Songs []*Song `toml:"songs"`
}

const (
	APP_VERSION      = "0.0.1"
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
				buf, err := dataMembersTomlBytes()
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
		{
			Name:    "songs",
			Aliases: []string{"s"},
			Usage:   "List Song",
			Flags:   []cli.Flag{},
			Action: func(ctx *cli.Context) error {
				buf, err := dataSongsTomlBytes()
				if err != nil {
					return err
				}
				file := &SongFile{}
				err = toml.Unmarshal(buf, file)
				for _, song := range file.Songs {
					attrs := []string{
						fmt.Sprint(song.No),
						song.Name,
					}
					fmt.Println(strings.Join(attrs, "\t"))
				}
				return err
			},
		},
	}
	app.Run(os.Args)
}
