package main

import (
	"os/exec"
	"runtime"
)


func GenPic(poeStr string, convCfg ConvCfg) string {

	if runtime.GOOS == "windows" {

		var convCmdSlice []string

		// Construct convert cmd slice
		convCmdSlice = append(convCmdSlice, convCfg.ImgInLoc)
		convCmdSlice = append(convCmdSlice, "-font")
		convCmdSlice = append(convCmdSlice, convCfg.Font)
		convCmdSlice = append(convCmdSlice, "-gravity")
		convCmdSlice = append(convCmdSlice, "center")
		convCmdSlice = append(convCmdSlice, "-fill")
		convCmdSlice = append(convCmdSlice, "rgba" + convCfg.ColorRGBA)
		convCmdSlice = append(convCmdSlice, "-pointsize")
		convCmdSlice = append(convCmdSlice, convCfg.FontSize)
		convCmdSlice = append(convCmdSlice, "-annotate")
		convCmdSlice = append(convCmdSlice, convCfg.XOffset + convCfg.YOffset)
		convCmdSlice = append(convCmdSlice, poeStr)
		convCmdSlice = append(convCmdSlice, "webview_ui/images/out_0.jpg")

		convCmd := exec.Command("bin/convert.exe", convCmdSlice...)
		_, err := convCmd.Output()
		ChkErr(err)
	}

	return "webview_ui/images/out_0.jpg"
}

