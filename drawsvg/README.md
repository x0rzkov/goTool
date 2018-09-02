## DrawSVG

## usage
```
package main

import (
	"fmt"

	"github.com/zmisgod/goTool/drawsvg"
)

func main() {
	svg := drawsvg.Create()
	svg.SetCircle(2, "#42526e", "transparent")

	var path drawsvg.Path
	path.Fill = "transparent"
	path.Stroke = "blue"
	path.StrokeWidth = 2
	path.FillOpacity = "0.4"
	path.Aid = "121"
	path.Alt = "121212"
	for i := 0; i < 3; i++ {
		var ipath drawsvg.IPath
		if i == 0 {
			ipath.Lat = "100"
			ipath.Long = "105"
			ipath.ID = 11111
			ipath.Name = "station1"
		} else if i == 1 {
			ipath.Lat = "120"
			ipath.Long = "125"
			ipath.Directive = "wheel"
			ipath.ID = 222222
			ipath.Name = "station2"
		} else if i == 2 {
			ipath.Lat = "150"
			ipath.Long = "90"
			ipath.ID = 33333
			ipath.Name = "station3"
		}
		path.PathInfo = append(path.PathInfo, ipath)
	}
	svg.Path = path
	content := svg.Draw()
	fmt.Println(content)
}
```

## result
<img src="https://github.com/zmisgod/goTool/blob/master/img/drawsvg/1.svg">


<img src="https://github.com/zmisgod/goTool/blob/master/img/drawsvg/1.jpg">

## contact me

[@zmisgod](https://weibo.com/zmisgod)

[zmis.me](https://zmis.me)