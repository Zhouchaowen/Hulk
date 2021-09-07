package control

import (
	"Hulk/utils"
	"fmt"
	"path"
)

func loadData(dir string, file string) map[string]interface{} {
	var data = make(map[string]interface{})
	utils.ReadJson(path.Join(dir, file), &data)
	for k, v := range data {
		fmt.Println(k, v)
	}
	return data
}
