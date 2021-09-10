package generates

import (
	"errors"
	"math/rand"
	"strconv"
)

const (
	maxInt                = 1<<31 - 1
	minInt                = -1<<31 - 1
	nonComplianceIntCount = 3
)

func generateRangeBool() bool {
	if rand.Intn(100)%2 == 0 {
		return true
	}
	return false
}

type BoolRule struct {
}

func (s *BoolRule) GetParamType() ParamType {
	return Bool
}

func (s *BoolRule) GetNonComplianceCount() int {
	return 0
}

func (s *BoolRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Int,
		Float64,
		String,
	}
}

// 获取[min，max]区间的值
func generateInt(s *IntRule) (int, error) {
	if s.Customized != "" || len(s.Customized) != 0 {
		return strconv.Atoi(s.Customized)
	}
	return generateRangeInt(s.Min, s.Max)
}

func generateRangeInt(min, max int) (int, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	randNum := rand.Intn(max-min) + min
	return randNum, nil
}

// 默认[minInt,maxInt]区间的值
func generateDefaultInt() (int, error) {
	var s = &IntRule{
		Min: minInt,
		Max: maxInt,
	}
	return generateInt(s)
}

// 获取[minInt，min]区间的值
func generateMinRangeInt(s *IntRule) (int, error) {
	s.Max = s.Min
	s.Min = minInt
	return generateInt(s)
}

// 获取[max，maxInt]区间的值
func generateMaxRangeInt(s *IntRule) (int, error) {
	s.Min = s.Max
	s.Max = maxInt
	return generateInt(s)
}

// 获取赋负值[minInt，0]
func generateNegativeInt() (int, error) {
	var s = &IntRule{
		Min: minInt,
		Max: 0,
	}
	return generateInt(s)
}

func generateNonComplianceInt(s *IntRule, idx int) (int, error) {
	switch idx {
	case 0:
		return generateMinRangeInt(s)
	case 1:
		return generateMaxRangeInt(s)
	case 2:
		return generateNegativeInt()
	}
	return 0, nil
}

type IntRule struct {
	Min        int
	Max        int
	Customized string
}

func (s *IntRule) GetParamType() ParamType {
	return Int
}

func (s *IntRule) GetNonComplianceCount() int {
	return nonComplianceIntCount
}

func (s *IntRule) GetNonComplianceOtherTypes() []ParamType {
	if s.Customized != "" || len(s.Customized) != 0 {
		return []ParamType{}
	}
	return []ParamType{
		Bool,
		Float64,
		String,
	}
}

func (s *IntRule) GetParams() []interface{} {
	var res []interface{}
	num, _ := generateInt(s)
	res = append(res, num)
	num, _ = generateMinRangeInt(s)
	res = append(res, num)
	num, _ = generateMaxRangeInt(s)
	res = append(res, num)
	num, _ = generateNegativeInt()
	res = append(res, num)

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		param, _ := generatorNonCompliance(otherTypes[i])
		res = append(res, param)
	}
	return res
}
