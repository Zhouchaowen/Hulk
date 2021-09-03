package models

import "Hulk/generates"

type InterfaceConfig struct {
	Id            int
	Agreement     string
	Name          string
	Addr          string
	Method        string
	RequestConfig generates.RequestConfig
	Response      map[string]interface{}
}
