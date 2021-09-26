package send

import (
	"Hulk/utils"
	"fmt"
	"log"
	"path"
)

type RequestControl struct {
	uid         string
	count       int
	Method      string
	Url         string
	ContentType int
	Header      map[string]string
}

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

func (s *RequestControl) Run(file string) {
	data := loadData("/Users/zdns/Desktop/Hulk", file)
	send := NewHttpRequest(s.Method, s.Url, s.ContentType, s.Header)
	for i, _ := range data {
		result, err := send.Send(data[i])
		if err != nil {
			log.Printf("request:%s", err)
		}
		_ = result
	}
}
