package main

import (
	"os/exec"
	"fmt"
)


func SetWallPaper(picLoc string, interval int) {

	setWlCmd := exec.Command("bin/wallpaper.exe", "webview_ui/images/out_0.jpg")
	_, err := setWlCmd.Output()
	ChkErr(err)

	fmt.Println("\n设置壁纸成功！\n\n将在 ", interval, " 小时后重新触发，您可以最小化以等待 ...")
}
