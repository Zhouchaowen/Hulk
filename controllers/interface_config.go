package controllers

import (
	"Hulk/generates"
	"Hulk/middleware"
	"Hulk/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type InterfaceConfig struct{}

func InterfaceConfigRegister(group *gin.RouterGroup) {
	ic := InterfaceConfig{}
	group.POST("", ic.AddConfig)
	group.GET("", ic.GetConfig)
	group.GET("/create", ic.CreateData)
}

func (s *InterfaceConfig) AddConfig(c *gin.Context) {
	params := &models.InterfaceConfig{}
	c.ShouldBindJSON(params)
	paramLimits, err := params.GenParamLimitMap()
	if err != nil {
		panic(err)
	}
	params.Insert()
	log.Println(paramLimits)
	middleware.ResponseSuccess(c, params)
}

func (s *InterfaceConfig) CreateData(c *gin.Context) {
	idStr := c.Query("id")
	if len(idStr) <= 0 {
		middleware.ResponseError(c, 2000, fmt.Errorf("id is must"))
	}
	id, _ := strconv.Atoi(idStr)
	var ic = models.InterfaceConfig{}
	ic.GetOneByKey(id)
	param, err := ic.GenParamLimitMap()
	if err != nil {
		middleware.ResponseError(c, 2001, fmt.Errorf("param err"))
	}
	generates.Generator("/Users/zdns/Desktop/Hulk", param)
	middleware.ResponseSuccess(c, "ok")
}

func (s *InterfaceConfig) GetConfig(c *gin.Context) {
	idStr := c.Query("id")
	if len(idStr) <= 0 {
		middleware.ResponseError(c, 2000, fmt.Errorf("id is must"))
	}
	id, _ := strconv.Atoi(idStr)
	res := models.GetInterfaceConfigModel(id)
	middleware.ResponseSuccess(c, res)
}
