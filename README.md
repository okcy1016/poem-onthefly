# poetry-desktop v1.1

欢迎使用诗词桌面！

程序使用 Go 语言开发，[Webview][6] 部分使用了 [Element UI][7] 。 

目前仅支持 Windows 。

# 简介
能够自动、定时地从 [github][1] 上随机获取一首唐、宋诗或宋词，并将文本整合在背景图片上，最后自动设置为壁纸 。

# 更新日志
### v1.1.1 2018.10.07
修复了由于错误判断了网络可用性而导致无法启动的问题。

### v1.1 - 2018.05.18
更新了使用 [webview][6] 渲染的、由 [Element UI][7] 组构的 GUI 界面 。

# 使用效果
![alt text][3]
![alt text][4]
![alt text][5]

# 下载
1. 请前往 [Releases][0] 页面。

# 自行编译时的注意事项
1. 先下载 [Releases][0] 中的包解压。
2. 请分别编译本 git 仓库根目录与 cli/poetry_cli 目录下的文件，将得到的 poetry-desktop.exe 覆盖 zip 文件解压目录的同名文件，将得到的 poetry_cli.exe 覆盖 zip 文件解压目录的 bin/poetry_cli.exe 文件。

# 致谢
感谢以下项目的贡献者：

1. 中华古诗词数据库  
https://github.com/chinese-poetry/chinese-poetry

2. A simple, fast, and fun package for building command line apps in Go  
https://github.com/urfave/cli

3. Tiny cross-platform webview library for C/C++/Golang  
https://github.com/zserge/webview

4. A Vue.js 2.0 UI Toolkit for Web  
https://github.com/ElemeFE/element

5. ImageMagick 7  
https://github.com/ImageMagick/ImageMagick

6. Set the desktop wallpaper on Windows  
https://github.com/sindresorhus/win-wallpaper

# To do
- [ ] 修正目录结构
- [ ] 更换 UI

# How to contribute 
![alt text][8]
(The picture above is from [gin][9] homepage.)


# Welcome your pull requests!

[0]: https://github.com/okcy1016/poetry-desktop/releases
[1]: https://github.com/chinese-poetry/chinese-poetry
[2]: https://yadi.sk/d/RUOf2iUF3WTFJM
[3]: https://github.com/okcy1016/poem-onthefly/raw/old/screenshots/Screenshot%20from%202018-05-18%2016-50-37.png
[4]: https://github.com/okcy1016/poem-onthefly/raw/old/screenshots/show_case_0.png
[5]: https://github.com/okcy1016/poem-onthefly/raw/old/screenshots/show_case_1.png
[6]: https://github.com/zserge/webview
[7]: https://github.com/ElemeFE/element
[8]: https://github.com/okcy1016/poetry-desktop/raw/master/screenshots/Screenshot%20from%202018-06-23%2021-10-04.png
[9]: https://gin-gonic.github.io/gin/
