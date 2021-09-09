package models

import (
	"Hulk/common"
	"Hulk/generates"
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
