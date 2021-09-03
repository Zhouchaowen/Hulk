package generates

import (
	"fmt"
	"github.com/srlemon/gen-id/generator"
	"github.com/srlemon/gen-id/utils"
	"math/rand"
	"reflect"
	"time"
)

type ParamLimit interface {
	GetParamType() string
	IsParent() bool
	GetNext() ParamLimit
}

type RequestConfig struct {
	Param  map[string]ParamLimit
	Header map[string]interface{}
}

// 生成地址
func generateAddress() string {
	g := generator.GeneratorData{}
	return g.GeneratorAddress()
}

// 生成银行卡号
func generatorBankID() string {
	g := generator.GeneratorData{}
	return g.GeneratorBankID()
}

// 生成邮箱
func generatorEmail() string {
	g := generator.GeneratorData{}
	return g.GeneratorEmail()
}

// 生成身份证号
func generatorIDCart() string {
	g := generator.GeneratorData{}
	g.GeneratorIDCart()
	return g.IDCard
}

// 生成手机号码
func generatorPhone() string {
	g := generator.GeneratorData{}
	return g.GeneratorPhone()
}

// 生成IP
func generatorIP() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

// 生成随机获取城市和地址
func generatorProvinceAdnCityRand() string {
	g := generator.GeneratorData{}
	return g.GeneratorProvinceAdnCityRand()
}

func generatorRandDate() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
	return time.Unix(utils.RandInt64(begin.Unix(), end.Unix()), 0)
}

func Generator(config RequestConfig) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range config.Param {
		switch v.GetParamType() {
		case reflect.Int.String():
			t, _ := v.(*IntRule)
			if num, err := generateRangeInt(t.Min, t.Max); err == nil {
				ret[k] = num
			}
		case reflect.String.String():
			t, _ := v.(*StringRule)
			if str, err := generateRangeString(t.Min, t.Max, t.MinLen, t.MaxLen); err == nil {
				ret[k] = str
			}
		case reflect.Float64.String():
			t, _ := v.(*FloatRule)
			if flo, err := generateRangeFloat(t.Min, t.Max, t.Retain); err == nil {
				ret[k] = flo
			}
		case reflect.Map.String():
			return nil
		case reflect.Array.String():
			return nil
		case reflect.Bool.String():
			return nil
		}
	}
	return ret
}
