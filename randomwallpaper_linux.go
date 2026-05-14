//go:build linux
// +build linux

package main

import (
	"flag"
	"os"

	"github.com/ggtony233/random_image/utils"
	"github.com/reujab/wallpaper"
)

func main() {
	Flag := flag.Bool("t", false, "only filepath")
	flag.Parse()
	path := utils.RandomImagePath(utils.GetJsonPath())
	if *Flag {
		os.Stdout.WriteString(path)
		os.Exit(0)
	}
	err := wallpaper.SetFromFile(path)
	if err != nil {
		print(err.Error())
		panic(err)
	}
}
