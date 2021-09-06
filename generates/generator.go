package generates

import (
	utils2 "Hulk/utils"
	"encoding/json"
	"fmt"
	"path"
	"strconv"
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
	Time
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
	Time:    "time",
}

type ParamLimit interface {
	GetParamType() ParamType
	GetNonComplianceCount() int
	GetNonComplianceOtherTypes() []ParamType
}

type RequestConfig struct {
	Param  map[string]ParamLimit
	Header map[string]interface{}
}

// 生成合规参数
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
		case Time:
			ret[k] = generatorRandTime()
		}
	}
	return ret
}

// 生成类型默认参数
func generatorNonCompliance(paramType ParamType) (interface{}, error) {
	switch paramType {
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

// 生成其它类型不合规参数
func generatorNonComplianceOtherTypeParam(paramLimit ParamLimit, paramType ParamType) interface{} {
	switch paramLimit.GetParamType() {
	case Bool:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Int:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Float64:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case String:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Array:
	case Map:
	case Email:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Address:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case BankID:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case IDCart:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Phone:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case IP:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	case Time:
		if res, err := generatorNonCompliance(paramType); err == nil {
			return res
		}
	}
	return nil
}

// 生成本类型不合规的参数
func generatorNonComplianceParam(paramLimit ParamLimit, idx int) interface{} {
	switch paramLimit.GetParamType() {
	case Int:
		t, _ := paramLimit.(*IntRule)
		if res, err := generateNonComplianceInt(t, idx); err == nil {
			return res
		}
	case Float64:
		t, _ := paramLimit.(*FloatRule)
		if res, err := generateNonComplianceFloat(t, idx); err == nil {
			return res
		}
	case String:
		t, _ := paramLimit.(*StringRule)
		if res, err := generateNonComplianceString(t, idx); err == nil {
			return res
		}
	}
	return nil
}

func getNonComplianceParam(config map[string]ParamLimit) []map[string]interface{} {
	var paramValue = make([]ParamLimit, len(config))
	var paramKey = make([]string, len(config))
	var num = 0
	for k, v := range config {
		paramValue[num] = v
		paramKey[num] = k
		num++
	}
	var res = make([]map[string]interface{}, 0)
	for i := 0; i < len(paramKey); i++ {
		val := paramValue[i]
		key := paramKey[i]
		for j := 0; j < val.GetNonComplianceCount(); j++ {
			var ret = make(map[string]interface{}, len(config))
			for k, v := range config {
				if key == k {
					ret[k] = generatorNonComplianceParam(v, j)
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
			res = append(res, ret)
		}
	}
	return res
}

func getNonComplianceOtherTypeParam(config map[string]ParamLimit) []map[string]interface{} {
	var paramValue = make([]ParamLimit, len(config))
	var paramKey = make([]string, len(config))
	var num = 0
	for k, v := range config {
		paramValue[num] = v
		paramKey[num] = k
		num++
	}
	var res = make([]map[string]interface{}, 0)
	for i := 0; i < len(paramKey); i++ {
		val := paramValue[i]
		key := paramKey[i]
		paramType := val.GetNonComplianceOtherTypes()
		for j := 0; j < len(paramType); j++ {
			var ret = make(map[string]interface{}, len(config))
			for k, v := range config {
				if key == k {
					ret[k] = generatorNonComplianceOtherTypeParam(v, paramType[j])
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
			res = append(res, ret)
		}
	}
	return res
}

func Generator(dir string, config map[string]ParamLimit) map[string]interface{} {
	//param := generatorParams(config)
	//fileParamName := path.Join(dir, "param.json")
	//if err := utils2.WriteJson(fileParamName, param); err != nil {
	//
	//}
	//ncParams := getNonComplianceParam(config)
	//fileParamName := path.Join(dir, "nc_param.json")
	//if err := utils2.WriteJson(fileParamName, ncParams); err != nil {
	//
	//}
	//
	ncOtherParams := getNonComplianceOtherTypeParam(config)
	fileParamName := path.Join(dir, "nc_other_param.json")
	if err := utils2.WriteJson(fileParamName, ncOtherParams); err != nil {

	}
	return nil
}
