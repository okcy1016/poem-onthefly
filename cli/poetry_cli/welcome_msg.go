package main

import (
	"fmt"
//	"os"
//	"bufio"
)


func WelcomeMsg() {

	fmt.Println("\n欢迎使用诗词桌面！\n\n本系统能够自动从网路上随机获取一首唐、宋诗或宋词，并将文本整合在背景图片上，最后自动设置为壁纸。\n\n要查看帮助请在当前位置打开cmd命令行输入: poetry-desktop --help\n\n本项目位于: https://github.com/okcy1016/poetry-desktop\n\n感谢以下项目的贡献者：\n\n1、中华古诗词数据库\nhttps://github.com/chinese-poetry/chinese-poetry\n\n2、ImageMagick 7\nhttps://github.com/ImageMagick/ImageMagick\n\n3、Set the desktop wallpaper on Windows\nhttps://github.com/sindresorhus/win-wallpaper\n\n请按回车键以进入程序主循环 ...\n")

	// Waiting for user to press Enter ...
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
}
