package main

import (
	"github.com/urfave/cli"
	"strconv"
	"os"
)

type ConvCfg struct {

	Font string
	FontSize string
	ImgInLoc string
	ColorRGBA string
	XOffset string
	YOffset string
}

func StartCliApp() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {

		cli.StringFlag {
			Name: "interval, i",
			Value: "1",
			Usage: "设置诗词桌面自动切换时间间隔，单位：小时",
		},

		cli.StringFlag {
			Name: "font, f",
			Value: "fonts/FangZhengKaiTiJianTi-1.ttf",
			Usage: "设置显示的字体",
		},

		cli.StringFlag {
			Name: "fontsize, s",
			Value: "25",
			Usage: "设置字体大小",
		},

		cli.StringFlag {
			Name: "imgloc, l",
			Value: "images/poetry_bg_0.jpg",
			Usage: "指定背景图片位置",
		},

		cli.StringFlag {
			Name: "color, c",
			Value: "(0,0,0,0.8)",
			Usage: "设置字体颜色，格式为 rgba",
		},

		cli.StringFlag {
			Name: "xoffset",
			Value: "+0",
			Usage: "设置文字整体在X轴方向的偏移量",
		},

		cli.StringFlag {
			Name: "yoffset",
			Value: "+0",
			Usage: "设置文字整体在Y轴方向的偏移量",
		},
	}
	
	app.Action = func(cmd *cli.Context) {

		interval, err := strconv.Atoi(cmd.String("interval"))
		ChkErr(err)

		// Getting ready for passing arguments
		var convCfg ConvCfg
		
		convCfg.Font = cmd.String("font")
		convCfg.FontSize = cmd.String("fontsize")
		convCfg.ColorRGBA = cmd.String("color")
		convCfg.ImgInLoc = cmd.String("imgloc")
		convCfg.XOffset = cmd.String("xoffset")
		convCfg.YOffset = cmd.String("yoffset")

		StartSchedule(interval, convCfg)
	}

	err := app.Run(os.Args)
	ChkErr(err)	
}
