package models

import (
	"Hulk/common"
	"Hulk/db"
	"Hulk/generates"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

func (s *InterfaceConfig) toModel() InterfaceConfigModel {
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
	return ifm
}

func (s *InterfaceConfig) Insert() int64 {
	var ifm = s.toModel()

	result := db.Db.Model(&ifm).Where("id = ?", ifm.Id)
	if result.Error != nil {
		result = db.Db.Create(&ifm)
	} else {
		result = db.Db.Model(&ifm).Updates(&ifm)
		if result.Error != nil {
			return 0
		}
	}

	return result.RowsAffected
}

func (s *InterfaceConfig) GetOneByKey(id int) {
	var res = GetInterfaceConfigModel(id)

	s.Id = res.Id
	s.Agreement = res.Agreement
	s.Name = res.Name
	s.Addr = res.Addr
	s.Method = res.Method
	var rc = make(map[string]interface{})
	json.Unmarshal([]byte(res.RequestConfig), &rc)
	s.RequestConfig = rc

	var r = make(map[string]interface{})
	json.Unmarshal([]byte(res.Response), &r)
	s.Response = r

	var h = make(map[string]interface{})
	json.Unmarshal([]byte(res.Header), &h)
	s.Header = h
}

func GetInterfaceConfigModel(id int) *InterfaceConfigModel {
	var res = &InterfaceConfigModel{}
	db.Db.First(res, id)
	return res
}
