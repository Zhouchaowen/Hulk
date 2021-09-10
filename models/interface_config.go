package models

import (
	"Hulk/common"
	"Hulk/db"
	"Hulk/generates"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type InterfaceConfig struct {
	Id            int                    `json:"id"`
	Agreement     string                 `json:"agreement"`
	Name          string                 `json:"name"`
	Addr          string                 `json:"addr"`
	Method        string                 `json:"method"`
	RequestConfig map[string]interface{} `json:"request_config"`
	Response      map[string]interface{} `json:"response"`
	Header        map[string]interface{} `json:"header"`
}

type InterfaceConfigModel struct {
	Id            int    `json:"id"`
	Agreement     string `json:"agreement"`
	Name          string `json:"name"`
	Addr          string `json:"addr"`
	Method        string `json:"method"`
	RequestConfig string `json:"request_config"`
	Response      string `json:"response"`
	Header        string `json:"header"`
}

func (s *InterfaceConfig) BindValidParam(c *gin.Context) error {
	return common.DefaultGetValidParams(c, s)
}

func (s *InterfaceConfig) GenParamLimitMap() (map[string]generates.ParamLimit, error) {
	var paramLimit = map[string]generates.ParamLimit{}
	for k, v := range s.RequestConfig {
		if config, ok := v.(map[string]interface{}); ok {
			param, err := generates.MapToParamLimitObject(config)
			if err != nil {
				return nil, err
			}
			paramLimit[k] = param
		}
	}
	return paramLimit, nil
}

func (s *InterfaceConfig) Insert() int64 {
	var ifm = InterfaceConfigModel{
		Id:        s.Id,
		Addr:      s.Addr,
		Agreement: s.Agreement,
		Name:      s.Name,
		Method:    s.Method,
	}
	rc, _ := json.Marshal(s.RequestConfig)
	ifm.RequestConfig = string(rc)
	r, _ := json.Marshal(s.Response)
	ifm.Response = string(r)
	h, _ := json.Marshal(s.Header)
	ifm.Header = string(h)
	result := db.Db.Create(&ifm)
	log.Println(result)
	return result.RowsAffected
}
