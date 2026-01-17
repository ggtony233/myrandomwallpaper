//go:build windows
// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/ggtony233/random_image/utils"
)

const (
	SPI_SETDESKWALLPAPER = 0x0014
	SPIF_UPDATEINIFILE   = 0x01
	SPIF_SENDCHANGE      = 0x02
)

func main() {
	// 1️⃣ 从 utils 获取随机图片路径
	path := utils.RandomImagePath(utils.GetJsonPath())
	if path == "null" {
		return
	}

	// 2️⃣ 设置壁纸
	setWallpaper(path)
}

func setWallpaper(path string) {

	user32 := syscall.NewLazyDLL("user32.dll")
	proc := user32.NewProc("SystemParametersInfoW")

	p, _ := syscall.UTF16PtrFromString(path)

	proc.Call(
		uintptr(SPI_SETDESKWALLPAPER),
		0,
		uintptr(unsafe.Pointer(p)),
		uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE),
	)
}
