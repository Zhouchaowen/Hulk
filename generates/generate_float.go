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

func generateFloat(s *FloatRule) (float64, error) {
	if s.Customized != "" || len(s.Customized) != 0 {
		return strconv.ParseFloat(s.Customized, 64)
	}
	return generateRangeFloat(s.Min, s.Max, s.Retain)
}

func generateRangeFloat(min, max float64, retain int) (float64, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	t := min + (max-min)*rand.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(retain)+"f", t), 64)
}

func generateDefaultFloat() (float64, error) {
	var s = &FloatRule{
		Min:    minFloat,
		Max:    maxFloat,
		Retain: retain,
	}
	return generateFloat(s)
}

// 获取[minFloat，min]区间的值
func generateMinRangeFloat(s *FloatRule) (float64, error) {
	s.Max = s.Min
	s.Min = minFloat
	return generateFloat(s)
}

// 获取[max，maxFloat]区间的值
func generateMaxRangeFloat(s *FloatRule) (float64, error) {
	s.Min = s.Max
	s.Max = maxFloat
	return generateFloat(s)
}

// 获取[max，max]区间小数点后保留 maxRetain 的值
func generateRetainRangeFloat(s *FloatRule) (float64, error) {
	s.Retain = maxRetain
	return generateFloat(s)
}

func generateNonComplianceFloat(s *FloatRule, idx int) (float64, error) {
	switch idx {
	case 0:
		return generateMinRangeFloat(s)
	case 1:
		return generateMaxRangeFloat(s)
	case 2:
		return generateRetainRangeFloat(s)
	}
	return 0, nil
}

type FloatRule struct {
	Min        float64
	Max        float64
	Retain     int
	Customized string
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
	fla, _ := generateFloat(s)
	res = append(res, fla)
	fla, _ = generateMinRangeFloat(s)
	res = append(res, fla)
	fla, _ = generateMaxRangeFloat(s)
	res = append(res, fla)
	fla, _ = generateRetainRangeFloat(s)
	res = append(res, fla)

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		param, _ := generatorNonCompliance(otherTypes[i])
		res = append(res, param)
	}
	return res
}
