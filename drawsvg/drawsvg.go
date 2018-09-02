package drawsvg

import (
	"fmt"
	"strconv"
)

//Canvas 画布信息
type Canvas struct {
	viewBox    viewBox
	reSize     int
	longResize int
	laResize   int
}

//ViewBox ViewBox
type viewBox struct {
	widthOne  int
	heightOne int
	widthTwo  int
	heightTwo int
}

//Circle 画圆点
type circle struct {
	radius int
	stroke string
	fill   string
}

//SetCircle 设置circle形状
func (svg *SVG) SetCircle(radius int, stroke, fill string) {
	if radius < 0 {
		svg.circle.radius = 1
	} else {
		svg.circle.radius = radius
	}
	if stroke == "" {
		svg.circle.stroke = "blue"
	} else {
		svg.circle.stroke = stroke
	}
	if fill == "" {
		svg.circle.fill = "transparent"
	} else {
		svg.circle.fill = fill
	}
}

//PolyLine 画直线
type PolyLine struct {
	PathInfo    []IPath
	points      string
	Fill        string
	Stroke      string
	StrokeWidth int    `json:"stroke-width"`
	FillOpacity string `json:"fill-opacity"`
	Aid         string //额外信息1
	Alt         string //额外信息2
}

//path指令，如果不传则默认直线
var directive = []string{
	"line",  //直线
	"wheel", //带半圆
}

//IPath IPath
type IPath struct {
	Long string
	Lat  string
	ID   int
	Name string

	//下面仅对PATH有效
	Directive string //指令
}

//Path 路径
type Path struct {
	PathInfo    []IPath
	d           string
	Fill        string
	Stroke      string
	StrokeWidth int    `json:"stroke-width"`
	FillOpacity string `json:"fill-opacity"`
	Aid         string //额外信息1
	Alt         string //额外信息2
}

//Output 输出
type Output struct {
	Floder   string
	FileName string
}

//SVG svg数据结构
type SVG struct {
	Canvas   Canvas   `json:"canvas"` //svg画布与基础参数
	circle   circle   //画点
	PolyLine PolyLine `json:"polyLine"` //画线
	Path     Path     `json:"path"`     //画线段
	Output   Output   `json:"output"`   //输出
}

//Create 创建一个SVG
func Create() *SVG {
	var svg SVG
	svg.Canvas.reSize = 1
	return &svg
}

//SetPathList 设置path
func (svg *SVG) SetPathList(pathLists []IPath) bool {
	svg.Path.PathInfo = pathLists
	return true
}

//SetPolyList 设置polyList
func (svg *SVG) SetPolyList(polyLine []IPath) bool {
	svg.PolyLine.PathInfo = polyLine
	return true
}

//Draw 画操作
func (svg *SVG) Draw() string {
	//画线段
	svg.DrawSVG()

	content := ""
	if svg.Path.d != "" {
		content += svg.Path.String()
	}
	if svg.PolyLine.points != "" {
		content += svg.PolyLine.String()
	}
	if len(svg.Path.PathInfo) > 0 {
		for _, v := range svg.Path.PathInfo {
			if v.Directive != "wheel" {
				longStr, latiStr := svg.dataConversion(v.Lat, v.Long)
				content += svg.circle.String(latiStr, longStr, v.Name, v.ID)
			}
		}
	}
	svg.Canvas.viewBox.widthTwo += 100
	svg.Canvas.viewBox.heightTwo += 100
	return svg.String(content)
}

func (c *viewBox) String() string {
	w1 := strconv.Itoa(c.widthOne)
	h1 := strconv.Itoa(c.heightOne)
	w2 := strconv.Itoa(c.widthTwo)
	h2 := strconv.Itoa(c.heightTwo)
	return w1 + "," + h1 + "," + w2 + "," + h2
}

func (c *circle) String(long, la, name string, id int) string {
	return fmt.Sprintf("<circle cx=\"%s\" cy=\"%s\" r=\"%d\" stroke=\"%s\" fill=\"%s\" alt=\"%s\" aid=\"%d\" />", la, long, c.radius, c.stroke, c.fill, name, id)
}

func (c *PolyLine) String() string {
	return fmt.Sprintf("<polyline points=\"%s\" stroke=\"%s\" fill=\"%s\" stroke-width=\"%d\" fill-opacity=\"%s\" alt=\"%s\" aid=\"%s\" />", c.points, c.Stroke, c.Fill, c.StrokeWidth, c.FillOpacity, c.Alt, c.Aid)
}

//DrawSVG 画D
func (svg *SVG) DrawSVG() {
	pathInfoCount := len(svg.Path.PathInfo)
	if pathInfoCount > 0 {
		path := ""
		for k, v := range svg.Path.PathInfo {
			prefix := "L "
			if k == 0 {
				prefix = "M"
			}
			if v.Directive == "wheel" {
				prefix = "Q "
			}
			longStr, latiStr := svg.dataConversion(v.Lat, v.Long)
			path += prefix + longStr + " " + latiStr + " "

			if v.Directive == "wheel" && k+1 <= pathInfoCount {
				longStr, latiStr := svg.dataConversion(svg.Path.PathInfo[k+1].Lat, svg.Path.PathInfo[k+1].Long)
				path += longStr + " " + latiStr + " "
			}
		}
		svg.Path.d = path
	}

	if len(svg.PolyLine.PathInfo) > 0 {
		points := ""
		for k, v := range svg.PolyLine.PathInfo {
			prefix := ", "
			if k == 0 {
				prefix = ""
			}
			longStr, latiStr := svg.dataConversion(v.Lat, v.Long)

			points += prefix + longStr + " " + latiStr + " "
		}
		svg.PolyLine.points = points
	}
}

func (c *Path) String() string {
	return fmt.Sprintf("<path d=\"%s\" stroke=\"%s\" fill=\"%s\" stroke-width=\"%d\" fill-opacity=\"%s\" alt=\"%s\" aid=\"%s\" />", c.d, c.Stroke, c.Fill, c.StrokeWidth, c.FillOpacity, c.Alt, c.Aid)
}

//String svg
func (svg *SVG) String(content string) string {
	return fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"%s\" version=\"1.1\"  width=\"100%%\" height=\"100%%\">%s</svg>", svg.Canvas.viewBox.String(), content)
}

//dataConversion 数据转换
func (svg *SVG) dataConversion(data1, data2 string) (string, string) {
	f1, _ := strconv.ParseFloat(data1, 32)
	res1 := int((f1 - float64(svg.Canvas.laResize)) * float64(svg.Canvas.reSize))

	f2, _ := strconv.ParseFloat(data2, 32)
	res2 := int((f2 - float64(svg.Canvas.longResize)) * float64(svg.Canvas.reSize))

	if res1 > svg.Canvas.viewBox.widthTwo {
		svg.Canvas.viewBox.widthTwo = res1
	}
	if res2 > svg.Canvas.viewBox.heightTwo {
		svg.Canvas.viewBox.heightTwo = res2
	}
	str1 := strconv.Itoa(res1)
	str2 := strconv.Itoa(res2)
	return str1, str2
}
