## Draw SVG

### url
https://api.zmis.me/v1/crh/line.generate

### method
POST

### ContentType
application/json

### Request Body Struct
```
//CHSR China High Speed Rewaiy
type CHSR struct {
	resize int    //缩放大小
	list   []List //线路列表
	circle Circle //站点设置 (圆)
}

//Circle info
type Circle struct {
	radis       int    //圆半径
	color       string //颜色
	border_width int  //border宽度
}

//List 线路
type List struct {
	train_id   int     //id
	train_name string  //此线路名称
	type      int      //类型，暂时无用，同意设置为1
	max_group  int     //此线路有多少分支
	color     string    //画线的颜色
	width     int      //画线的宽度
	station   []Station //站点位置 
}

//Station 地理位置信息
type Station struct {
	id          int    //id
	station_name string //站点名称
	longtitude  string //坐标x
	latitude    string //坐标y 
	type        int    //这个是第几个分支，从0开始，0，1，2
	directive   int   //是否弯曲，弯曲会找前一个数据进行相应的弯曲，不弯曲0， 弯曲为1
}
```

### Response Struct
```
type Response struct
{
    code int //成功200， 失败!=200
    data string //返回svg数据
    msg string //消息，紧关注code !=200时的消息
}
```

### Request Body (example)
```
{"resize":1, "circle": {
	"radis":4,
	"color": "#000",
	"border_width": 6
},"list":[{"train_id":7,"train_name":"上海地铁7号线", "type": 1, "max_group":3,"color":"#FF7F00","width":8,"station":[
	{
		"id":1,
		"station_name":"站1",
		"longtitude":"1603.29",
		"latitude":"224.90",
		"type":0,
		"directive":0
	},
	{
		"id":2,
		"station_name":"站2",
		"longtitude":"1603.29",
		"latitude":"259.48",
		"type":0,
		"directive":0
	},
	{
		"id":3,
		"station_name":"站3",
		"longtitude":"1603.29",
		"latitude":"294.58",
		"type":0,
		"directive":0
	},
	{
		"id":4,
		"station_name":"站4",
		"longtitude":"1603.29",
		"latitude":"329.68",
		"type":0,
		"directive":0
	},
	{
		"id":5,
		"station_name":"站5",
		"longtitude":"1603.29",
		"latitude":"362.83",
		"type":0,
		"directive":0
	},
	{
		"id":6,
		"station_name":"站6",
		"longtitude":"1603.29",
		"latitude":"397.41",
		"type":0,
		"directive":0
	},
	{
		"id":7,
		"station_name":"站7",
		"longtitude":"1603.29",
		"latitude":"431.86",
		"type":0,
		"directive":0
	},
	{
		"id":8,
		"station_name":"站8",
		"longtitude":"1603.29",
		"latitude":"466.31",
		"type":0,
		"directive":0
	},
	{
		"id":9,
		"station_name":"站9",
		"longtitude":"1603.29",
		"latitude":"500.89",
		"type":0,
		"directive":0
	},
	{
		"id":10,
		"station_name":"站10",
		"longtitude":"1603.29",
		"latitude":"535.34",
		"type":0,
		"directive":0
	},
	{
		"id":11,
		"station_name":"站11",
		"longtitude":"1603.29",
		"latitude":"569.92",
		"type":0,
		"directive":0
	},
	{
		"id":12,
		"station_name":"站12",
		"longtitude":"1603.29",
		"latitude":"604.37",
		"type":0,
		"directive":0
	},
	{
		"id":13,
		"station_name":"站13",
		"longtitude":"1603.29",
		"latitude":"638.82",
		"type":0,
		"directive":0
	},
	{
		"id":14,
		"station_name":"站14",
		"longtitude":"1603.29",
		"latitude":"673.27",
		"type":0,
		"directive":0
	},
	{
		"id":15,
		"station_name":"站15",
		"longtitude":"1603.29",
		"latitude":"707.72",
		"type":0,
		"directive":0
	},
	{
		"id":16,
		"station_name":"站16",
		"longtitude":"1603.29",
		"latitude":"764.92",
		"type":0,
		"directive":0
	},
	{
		"id":17,
		"station_name":"站17",
		"longtitude":"1651.26",
		"latitude":"815.10",
		"type":0,
		"directive":0
	},
	{
		"id":18,
		"station_name":"站18",
		"longtitude":"1706.38",
		"latitude":"870.35",
		"type":0,
		"directive":0
	},
	{
		"id":19,
		"station_name":"站19",
		"longtitude":"1760.85",
		"latitude":"926.77",
		"type":0,
		"directive":0
	},
	{
		"id":20,
		"station_name":"站20",
		"longtitude":"1763.71",
		"latitude":"1021.15",
		"type":0,
		"directive":0
	},
	{
		"id":21,
		"station_name":"站21",
		"longtitude":"1762.54",
		"latitude":"1105.00",
		"type":0,
		"directive":0
	},
	{
		"id":22,
		"station_name":"站22",
		"longtitude":"1763.71",
		"latitude":"1240.85",
		"type":0,
		"directive":0
	},
	{
		"id":23,
		"station_name":"站23",
		"longtitude":"1763.71",
		"latitude":"1339.00",
		"type":0,
		"directive":0
	},
	{
		"id":24,
		"station_name":"站24",
		"longtitude":"1763.71",
		"latitude":"1449.00",
		"type":0,
		"directive":0
	},
	{
		"id":23,
		"station_name":"站23支线1",
		"longtitude":"1763.71",
		"latitude":"1339.00",
		"type":1,
		"directive":0
	},
	{
		"id":25,
		"station_name":"站25支线2",
		"longtitude":"1363.71",
		"latitude":"926.77",
		"type":1,
		"directive":0
	},
	{
		"id":25,
		"station_name":"站25支线2",
		"longtitude":"1363.71",
		"latitude":"926.77",
		"type":2,
		"directive":0
	},
	{
		"id":26,
		"station_name":"站26曲线1",
		"longtitude":"1163.71",
		"latitude":"826.77",
		"type":2,
		"directive":1
	},
	{
		"id":27,
		"station_name":"站27支线2",
		"longtitude":"1363.71",
		"latitude":"626.77",
		"type":2,
		"directive":0
	}
]}]}
```

### Response (example)
```
{
    "code": 200,
    "data": "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0,0,1793,1531\" version=\"1.1\"  width=\"100%\" height=\"100%\"><g><path d=\"M1533 206 L 1533 241 L 1533 276 L 1533 311 L 1533 344 L 1533 379 L 1533 413 L 1533 448 L 1533 482 L 1533 517 L 1533 551 L 1533 586 L 1533 620 L 1533 655 L 1533 689 L 1533 746 L 1581 797 L 1636 852 L 1690 908 L 1693 1003 L 1692 1087 L 1693 1222 L 1693 1321 L 1693 1431 \" stroke=\"#FF7F00\" fill=\"transparent\" stroke-width=\"8\" fill-opacity=\"0.4\" alt=\"上海地铁7号线\" aid=\"7\" /><circle cx=\"1533\" cy=\"206\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"1\" /><circle cx=\"1533\" cy=\"241\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"2\" /><circle cx=\"1533\" cy=\"276\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"3\" /><circle cx=\"1533\" cy=\"311\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"4\" /><circle cx=\"1533\" cy=\"344\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"5\" /><circle cx=\"1533\" cy=\"379\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"6\" /><circle cx=\"1533\" cy=\"413\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"7\" /><circle cx=\"1533\" cy=\"448\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"8\" /><circle cx=\"1533\" cy=\"482\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"9\" /><circle cx=\"1533\" cy=\"517\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"10\" /><circle cx=\"1533\" cy=\"551\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"11\" /><circle cx=\"1533\" cy=\"586\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"12\" /><circle cx=\"1533\" cy=\"620\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"13\" /><circle cx=\"1533\" cy=\"655\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"14\" /><circle cx=\"1533\" cy=\"689\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"15\" /><circle cx=\"1533\" cy=\"746\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"16\" /><circle cx=\"1581\" cy=\"797\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"17\" /><circle cx=\"1636\" cy=\"852\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"18\" /><circle cx=\"1690\" cy=\"908\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"19\" /><circle cx=\"1693\" cy=\"1003\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"20\" /><circle cx=\"1692\" cy=\"1087\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"21\" /><circle cx=\"1693\" cy=\"1222\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"22\" /><circle cx=\"1693\" cy=\"1321\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"23\" /><circle cx=\"1693\" cy=\"1431\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"24\" /></g><g><path d=\"M1693 1321 L 1293 908 \" stroke=\"#FF7F00\" fill=\"transparent\" stroke-width=\"8\" fill-opacity=\"0.4\" alt=\"上海地铁7号线\" aid=\"7\" /><circle cx=\"1693\" cy=\"1321\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"23\" /><circle cx=\"1293\" cy=\"908\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"25\" /></g><g><path d=\"M1293 908 Q 856 1145 1093 808 L 1293 608 \" stroke=\"#FF7F00\" fill=\"transparent\" stroke-width=\"8\" fill-opacity=\"0.4\" alt=\"上海地铁7号线\" aid=\"7\" /><circle cx=\"1293\" cy=\"908\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"25\" /><circle cx=\"1093\" cy=\"808\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"26\" /><circle cx=\"1293\" cy=\"608\" r=\"4\" stroke=\"#000\" fill=\"#fff\" alt=\"\" aid=\"27\" /></g></svg>",
    "msg": "ok"
}
```
