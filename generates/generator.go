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
	ParamType ParamType  `json:"param_type"`
	Param     ParamLimit `json:"param"`
}

// 通过 map 构造 ParamLimit 对象
func MapToParamLimitObject(config map[string]interface{}) (ParamLimit, error) {
	num, err := strconv.Atoi(fmt.Sprintf("%.0f", config["param_type"]))
	if err != nil {
		return nil, err
	}
	switch ParamType(num) {
	case Bool:
	case Int:
		var intRule = &IntRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, intRule)
		if err != nil {
			return nil, err
		}
		return intRule, nil
	case Float64:
		var floatRule = &FloatRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, floatRule)
		if err != nil {
			return nil, err
		}
		return floatRule, nil
	case String:
		var stringRule = &StringRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, stringRule)
		if err != nil {
			return nil, err
		}
		return stringRule, nil
	case Array:
		var arrayRule = &ArrayRule{}
		if param, ok := config["param"].(map[string]interface{}); ok {
			for k, v := range param {
				if k == "len" {
					len, err := strconv.Atoi(fmt.Sprintf("%.0f", v))
					if err != nil {
						return nil, err
					}
					arrayRule.Len = len
				} else {
					if child, ok := v.(map[string]interface{}); ok {
						val, err := MapToParamLimitObject(child)
						if err != nil {
							return nil, err
						}
						arrayRule.Type = val
					}
				}
			}
		}
		return arrayRule, nil
	case Map:
		var mapRule = &MapRule{}
		if param, ok := config["param"].(map[string]interface{}); ok {
			if types, ok := param["types"]; ok {
				mapRule.Types = make(map[string]ParamLimit, len(types.(map[string]interface{})))
				for k, v := range types.(map[string]interface{}) {
					if child, ok := v.(map[string]interface{}); ok {
						val, err := MapToParamLimitObject(child)
						if err != nil {
							return nil, err
						}
						mapRule.Types[k] = val
					}
				}
			}
		}
		return mapRule, nil
	case Email:
		var emailRule = &EmailRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, emailRule)
		if err != nil {
			return nil, err
		}
		return emailRule, nil
	case Address:
		var addressRule = &AddressRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, addressRule)
		if err != nil {
			return nil, err
		}
		return addressRule, nil
	case BankID:
		var bankIdRule = &BankIdRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, bankIdRule)
		if err != nil {
			return nil, err
		}
		return bankIdRule, nil
	case IDCart:
		var idCartRule = &IdCartRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, idCartRule)
		if err != nil {
			return nil, err
		}
		return idCartRule, nil
	case Phone:
		var phoneRule = &PhoneRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, phoneRule)
		if err != nil {
			return nil, err
		}
		return phoneRule, nil
	case IP:
		var ipRule = &IpRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, ipRule)
		if err != nil {
			return nil, err
		}
		return ipRule, nil
	case Time:
		var timeRule = &TimeRule{}
		buf, err := json.Marshal(config["param"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buf, timeRule)
		if err != nil {
			return nil, err
		}
		return timeRule, nil
	}
	return nil, fmt.Errorf("no such type %d", num)
}

// 生成合规参数
func generatorParams(config map[string]ParamLimit) map[string]interface{} {
	var ret = make(map[string]interface{}, len(config))
	for k, v := range config {
		switch v.GetParamType() {
		case Int:
			t, _ := v.(*IntRule)
			if num, err := generateInt(t); err == nil {
				ret[k] = num
			}
		case String:
			t, _ := v.(*StringRule)
			if str, err := generateString(t); err == nil {
				ret[k] = str
			}
		case Float64:
			t, _ := v.(*FloatRule)
			if flo, err := generateFloat(t); err == nil {
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
			t, _ := v.(*EmailRule)
			ret[k] = generatorEmail(t)
		case Address:
			ret[k] = generateAddress()
		case BankID:
			t, _ := v.(*BankIdRule)
			ret[k] = generatorBankID(t)
		case IDCart:
			t, _ := v.(*IdCartRule)
			ret[k] = generatorIDCart(t)
		case IP:
			t, _ := v.(*IpRule)
			ret[k] = generatorIP(t)
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
	case Phone:
		t, _ := paramLimit.(*PhoneRule)
		if res, err := generateNonCompliancePhone(t, idx); err == nil {
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

		// 生成Map
		if t, ok := val.(*MapRule); ok {
			chi := getNonComplianceParam(t.Types)
			for l, _ := range chi {
				ma := generatorParams(config)
				ma[key] = chi[l]
				res = append(res, ma)
			}
			continue
		}

		// 生成Array
		if t, ok := val.(*ArrayRule); ok {
			var chr = map[string]ParamLimit{
				"key": t.Type,
			}
			chi := getNonComplianceParam(chr)
			for l, _ := range chi {
				var arr = make([]interface{}, t.Len)
				for k := 0; k < t.Len; k++ {
					arr[k] = chi[l]["key"]
				}
				ma := generatorParams(config)
				ma[key] = arr
				res = append(res, ma)
			}
			continue
		}

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

		// 生成Map
		if t, ok := val.(*MapRule); ok {
			chi := getNonComplianceOtherTypeParam(t.Types)
			for l, _ := range chi {
				ma := generatorParams(config)
				ma[key] = chi[l]
				res = append(res, ma)
			}
			continue
		}

		// 生成Array
		if t, ok := val.(*ArrayRule); ok {
			var chr = map[string]ParamLimit{
				"key": t.Type,
			}
			chi := getNonComplianceOtherTypeParam(chr)
			for l, _ := range chi {
				var arr = make([]interface{}, t.Len)
				for k := 0; k < t.Len; k++ {
					arr[k] = chi[l]["key"]
				}
				ma := generatorParams(config)
				ma[key] = arr
				res = append(res, ma)
			}
			continue
		}

		paramTypes := val.GetNonComplianceOtherTypes()
		for j := 0; j < len(paramTypes); j++ {
			var ret = make(map[string]interface{}, len(config))
			for k, v := range config {
				if key == k {
					ret[k] = generatorNonComplianceOtherTypeParam(v, paramTypes[j])
				} else {
					var chr = map[string]ParamLimit{
						"key": v,
					}
					tt := generatorParams(chr)
					ret[k] = tt["key"]
				}
			}
			res = append(res, ret)
		}
	}
	return res
}

func Generator(dir string, config map[string]ParamLimit) map[string]interface{} {
	//param := generatorParams(config)
	//fileParamName := path.Join(dir, "param.json")
	//var data = make([]map[string]interface{}, 1)
	//data[0] = param
	//if err := utils2.WriteJson(fileParamName, data); err != nil {
	//
	//}
	//ncParams := getNonComplianceParam(config)
	//for i, _ := range ncParams {
	//	b, _ := json.Marshal(ncParams[i])
	//	fmt.Println(string(b))
	//}
	//fmt.Println(len(ncParams))
	//fileParamName := path.Join(dir, "nc_param.json")
	//if err := utils2.WriteJson(fileParamName, ncParams); err != nil {
	//
	//}

	ncOtherParams := getNonComplianceOtherTypeParam(config)
	for i, _ := range ncOtherParams {
		b, _ := json.Marshal(ncOtherParams[i])
		fmt.Println(string(b))
	}
	fmt.Println(len(ncOtherParams))
	fileParamName := path.Join(dir, "nc_other_param.json")
	if err := utils2.WriteJson(fileParamName, ncOtherParams); err != nil {

	}
	return nil
}

type ParamNode struct {
	Key  string        `json:"key"`
	List []interface{} `json:"list"`
}

func SpreadParams(config map[string]ParamLimit) []ParamNode {
	var ret = make([]ParamNode, 0)
	for k, v := range config {
		switch v.GetParamType() {
		case Int:
			t, _ := v.(*IntRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case String:
			t, _ := v.(*StringRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case Float64:
			t, _ := v.(*FloatRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case Map: // 坑
			t, _ := v.(*MapRule)
			children := SpreadParams(t.Types)
			for i, _ := range children {
				node := ParamNode{
					Key:  k + "." + children[i].Key,
					List: children[i].List,
				}
				ret = append(ret, node)
			}
		case Array: // 坑
			t, _ := v.(*ArrayRule)
			for i := 0; i < t.Len; i++ {
				var chr = map[string]ParamLimit{
					"": t.Type,
				}
				children := SpreadParams(chr)
				for j, _ := range children {
					node := ParamNode{
						Key:  k + "." + strconv.Itoa(i) + children[j].Key,
						List: children[j].List,
					}
					ret = append(ret, node)
				}
			}
		case Bool:

		case Email:
			t, _ := v.(*EmailRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case Address:
			t, _ := v.(*AddressRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case BankID:
			t, _ := v.(*BankIdRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case IDCart:
			t, _ := v.(*IdCartRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case IP:
			t, _ := v.(*IpRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case Phone:
			t, _ := v.(*PhoneRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		case Time:
			t, _ := v.(*TimeRule)
			node := ParamNode{
				Key:  k,
				List: t.GetParams(),
			}
			ret = append(ret, node)
		}
	}
	return ret
}
