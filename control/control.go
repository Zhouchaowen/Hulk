package control

import (
	"Hulk/models"
	"Hulk/send"
	"Hulk/utils"
	"fmt"
	"path"
)

func loadData(dir string, file string) []map[string]interface{} {
	var data = make([]map[string]interface{}, 0)
	utils.ReadJson(path.Join(dir, file), &data)
	for i, _ := range data {
		for k, v := range data[i] {
			fmt.Println(k, v)
		}
	}
	return data
}

func Run(dir string, file string, config models.InterfaceConfig) {
	data := loadData(dir, file)
	var send = send.HttpRequest{
		Method:      send.POST,
		Url:         "http://10.2.0.153:6021/sg_logstatistics_cmd",
		ContentType: send.ContentTypeJson,
		Header:      map[string]string{},
	}
	send.Init()
	for i, _ := range data {
		send.Send(data[i])
	}
}
