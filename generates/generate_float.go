package generates

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	maxFloat  = float64(1 << 31)
	minFloat  = -float64(1 << 31)
	retain    = 2
	maxRetain = 10
)

func generateRangeFloat(min, max float64, retain int) (float64, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	t := min + (max-min)*rand.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(retain)+"f", t), 64)
}

func generateDefaultFloat() (float64, error) {
	return generateRangeFloat(minFloat, maxFloat, retain)
}

// 获取[minFloat，min]区间的值
func generateMinRangeFloat(min float64, retain int) (float64, error) {
	return generateRangeFloat(minFloat, min, retain)
}

// 获取[max，maxFloat]区间的值
func generateMaxRangeFloat(max float64, retain int) (float64, error) {
	return generateRangeFloat(max, maxFloat, retain)
}

// 获取[max，max]区间小数点后保留 maxRetain 的值
func generateRetainRangeFloat(min, max float64) (float64, error) {
	return generateRangeFloat(min, max, maxRetain)
}

func generateNonComplianceFloat(floatRule *FloatRule, idx int) (float64, error) {
	switch idx {
	case 0:
		return generateMinRangeFloat(floatRule.Min, floatRule.Retain)
	case 1:
		return generateMaxRangeFloat(floatRule.Max, floatRule.Retain)
	case 2:
		return generateRetainRangeFloat(floatRule.Min, floatRule.Max)
	}
	return 0, nil
}

type FloatRule struct {
	Min    float64
	Max    float64
	Retain int
}

func (s *FloatRule) GetParamType() ParamType {
	return Float64
}

func (s *FloatRule) GetNonComplianceCount() int {
	return 3
}

func (s *FloatRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		String,
	}
}

func (s *FloatRule) GetParams() []interface{} {
	var res []interface{}
	fla, _ := generateRangeFloat(s.Min, s.Max, s.Retain)
	res = append(res, fla)
	fla, _ = generateMinRangeFloat(s.Min, s.Retain)
	res = append(res, fla)
	fla, _ = generateMaxRangeFloat(s.Max, s.Retain)
	res = append(res, fla)
	fla, _ = generateRetainRangeFloat(s.Min, s.Max)
	res = append(res, fla)

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		param, _ := generatorNonCompliance(otherTypes[i])
		res = append(res, param)
	}
	return res
}
