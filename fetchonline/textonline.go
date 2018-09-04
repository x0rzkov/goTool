package fetchonline

import (
	"fmt"
	"os"
)

//FetchOnline 从远程下载
type FetchOnline struct {
	TargetURL      string
	DownloadFloder string
}

//SetDownloadFloder 设置下载目录
func (t *FetchOnline) SetDownloadFloder(path string) {

}

//StartServer 开启下载服务器
func (t *FetchOnline) StartServer() {

}

//CheckFloder 检查文件夹是否存在
func (t *FetchOnline) CheckFloder() (bool, error) {
	checkFloderNotExists, err := checkPathIsNotExists(t.DownloadFloder)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if checkFloderNotExists {
		err := os.MkdirAll(t.DownloadFloder, 0777)
		if err != nil {
			return false, err
		}
		fmt.Printf("create floder %s successful\n", t.DownloadFloder)
		return true, nil
	}
	return false, err
}

//Download 下载
func (t *FetchOnline) Download() {

}

//CheckQueue 检查redis队列
func (t *FetchOnline) CheckQueue() {

}

//HandleQueue 处理队列数据
func (t *FetchOnline) HandleQueue() {

}
