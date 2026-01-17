//go:build linux
// +build linux

package main

import (
	"github.com/ggtony233/random_image/utils"
	"github.com/reujab/wallpaper"
)

func main() {
	path := utils.RandomImagePath(utils.GetJsonPath())
	err := wallpaper.SetFromFile(path)
	if err != nil {
		print(err.Error())
		panic(err)
	}
}
