package generates

import (
	"encoding/json"
	"fmt"
	"github.com/srlemon/gen-id/generator"
	"github.com/srlemon/gen-id/utils"
	"math/rand"
	"strconv"
	"time"
)

type ParamType int

const (
	Invalid ParamType = iota
	Bool
	Int
	Float64
	String
	Array
	Map
	Email
	Address
	BankID
	IDCart
	Phone
	IP
)

func (k ParamType) String() string {
	if int(k) < len(ParamTypeNames) {
		return ParamTypeNames[k]
	}
	return "paramType" + strconv.Itoa(int(k))
}

var ParamTypeNames = []string{
	Invalid: "invalid",
	Bool:    "bool",
	Int:     "int",
	Float64: "float64",
	String:  "string",
	Array:   "array",
	Map:     "map",
	Email:   "email",
	Address: "address",
	BankID:  "bankId",
	IDCart:  "idCart",
	Phone:   "phone",
	IP:      "ip",
}

type ParamLimit interface {
	GetParamType() ParamType
	GetNonComplianceCount() int
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

func generatorParams(config map[string]ParamLimit) map[string]interface{} {
	var ret = make(map[string]interface{}, len(config))
	for k, v := range config {
		switch v.GetParamType() {
		case Int:
			t, _ := v.(*IntRule)
			if num, err := generateRangeInt(t.Min, t.Max); err == nil {
				ret[k] = num
			}
		case String:
			t, _ := v.(*StringRule)
			if str, err := generateRangeString(t.Min, t.Max, t.MinLen, t.MaxLen); err == nil {
				ret[k] = str
			}
		case Float64:
			t, _ := v.(*FloatRule)
			if flo, err := generateRangeFloat(t.Min, t.Max, t.Retain); err == nil {
				ret[k] = flo
			}
		case Map:
			t, _ := v.(*MapRule)
			ret[k] = generatorParams(t.Types)
		case Array:
			t, _ := v.(*ArrayRule)
			arr := make([]interface{}, t.Len)
			for i := 0; i < t.Len; i++ {
				idx := strconv.Itoa(i)
				var chr = map[string]ParamLimit{
					idx: t.Type,
				}
				tt := generatorParams(chr)
				arr[i] = tt[idx]
			}
			ret[k] = arr
		case Bool:
			ret[k] = generateRangeBool()
		case Email:
			ret[k] = generatorEmail()
		case Address:
			ret[k] = generateAddress()
		case BankID:
			ret[k] = generatorBankID()
		case IDCart:
			ret[k] = generatorIDCart()
		case IP:
			ret[k] = generatorIP()
		case Phone:
			ret[k] = generatorPhone()
		}
	}
	return ret
}

func generatorNonCompliance(idx int) (interface{}, error) {
	switch ParamType(idx) {
	case Bool:
		return generateRangeBool(), nil
	case Int:
		return generateDefaultInt()
	case Float64:
		return generateDefaultFloat()
	case String:
		return generateDefaultString()
	}
	return nil, nil
}

func generatorNonComplianceParam(paramLimit ParamLimit, idx int) interface{} {
	switch paramLimit.GetParamType() {
	case Int:
		if res, err := generatorNonCompliance(idx); err == nil {
			return res
		}
	case String:
		return nil
	}
	return nil
}

func generatorNonComplianceTypeParam(paramLimit ParamLimit, idx int) interface{} {
	switch paramLimit.GetParamType() {
	case Int:
		t, _ := paramLimit.(*IntRule)
		if res, err := generateNonComplianceInt(t, idx); err == nil {
			return res
		}
	case String:
		return nil
	}
	return nil
}

func Generator(path string, config map[string]ParamLimit) map[string]interface{} {
	var paramValue = make([]ParamLimit, len(config))
	var paramKey = make([]string, len(config))
	var num = 0
	for k, v := range config {
		paramValue[num] = v
		paramKey[num] = k
		num++
	}

	for i := 0; i < len(paramKey); i++ {
		val := paramValue[i]
		key := paramKey[i]
		for j := 0; j < val.GetNonComplianceCount(); j++ {
			var ret = make(map[string]interface{}, len(config))
			for k, v := range config {
				if key == k {
					ret[k] = generatorNonComplianceTypeParam(v, j)
				} else {
					var chr = map[string]ParamLimit{
						"key": v,
					}
					tt := generatorParams(chr)
					ret[k] = tt["key"]
				}
			}
			b, _ := json.Marshal(ret)
			fmt.Println(string(b))

		}
	}
	return nil
}
