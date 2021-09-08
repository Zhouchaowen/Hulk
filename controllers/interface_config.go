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
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	log.Println(params)
}
