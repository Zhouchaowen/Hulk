package models

import (
	"Hulk/common"
	"Hulk/generates"
	"github.com/gin-gonic/gin"
)

type InterfaceConfig struct {
	Id            int                     `json:"id"`
	Agreement     string                  `json:"agreement"`
	Name          string                  `json:"name"`
	Addr          string                  `json:"addr"`
	Method        string                  `json:"method"`
	RequestConfig generates.RequestConfig `json:"request_config"`
	Response      map[string]interface{}  `json:"response"`
}

func (s *InterfaceConfig) BindValidParam(c *gin.Context) error {
	return common.DefaultGetValidParams(c, s)
}
