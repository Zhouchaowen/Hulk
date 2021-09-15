package controllers

import (
	"Hulk/generates"
	"Hulk/middleware"
	"Hulk/models"
	"Hulk/send"
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
	group.GET("/run", ic.Run)
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

func (s *InterfaceConfig) GetConfig(c *gin.Context) {
	idStr := c.Query("id")
	if len(idStr) <= 0 {
		middleware.ResponseError(c, 2000, fmt.Errorf("id is must"))
	}
	id, _ := strconv.Atoi(idStr)
	res := models.GetInterfaceConfigModel(id)
	middleware.ResponseSuccess(c, res)
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

func (s *InterfaceConfig) Run(c *gin.Context) {
	idStr := c.Query("id")
	if len(idStr) <= 0 {
		middleware.ResponseError(c, 2000, fmt.Errorf("id is must"))
	}
	id, _ := strconv.Atoi(idStr)
	var ic = models.InterfaceConfig{}
	ic.GetOneByKey(id)

	var hr = &send.RequestControl{
		Method:      ic.Method,
		Url:         ic.Addr,
		ContentType: ic.ContentType,
		Header:      ic.Header,
	}
	hr.Run("c_param.json")
}
