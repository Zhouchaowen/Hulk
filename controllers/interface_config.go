package controllers

import (
	"Hulk/middleware"
	"Hulk/models"
	"github.com/gin-gonic/gin"
	"log"
)

type InterfaceConfig struct{}

func InterfaceConfigRegister(group *gin.RouterGroup) {
	ic := InterfaceConfig{}
	group.POST("/add", ic.AddConfig)
}

func (s *InterfaceConfig) AddConfig(c *gin.Context) {
	params := &models.InterfaceConfig{}
	c.ShouldBindJSON(params)
	paramLimits, err := params.GenParamLimitMap()
	if err != nil {
		panic(err)
	}
	log.Println(paramLimits)
	log.Println(params.Insert())
	middleware.ResponseSuccess(c, params)
}
