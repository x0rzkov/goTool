# goTool

a useful tool for golang

## install

```
go get github.com/goTool
```

## fileserver

### usage
```
package main

import fileserver "github.com/zmisgod/goTool/fileServer"

func main() {
	port := 12345
	host := ""
	showTem := "tem"
	template := ""
	basePath := "./../pictureSpider/"
	validDownload := []string{"svg", "jpg", "png", "jpeg", "gif"}
	dir := fileserver.Create( port, basePath, host, showTem, template, validDownload)
	dir.CreateServer()
}
```
- `port` 文件系统会开启一个http server，这就是http server 的端口号
- `basePath` 基础目录，开启http server 后会打开你设置的`basePath`的文件夹的目录
- `host` 设置http server的`host`
- `showTem` 设置输出的类型，可选`json`与`template`，`json`会以`json`格式输出，`template`则会以网页的形式输出
- `template` 如果你觉得我写的html不是很好看或者有其他的需要，你可以复写默认的html。
- `validDownload` 默认可以打开的文件类型

## contact me

[@zmisgod](https://weibo.com/zmisgod)

[zmis.me](https://zmis.me)

