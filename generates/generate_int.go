package generates

import (
	"errors"
	"math/rand"
)

const (
	maxInt             = 1<<31 - 1
	minInt             = -1<<31 - 1
	nonComplianceCount = 3
)

func generateRangeBool() bool {
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

// 获取[min，max]区间的值
func generateRangeInt(min, max int) (int, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	randNum := rand.Intn(max-min) + min
	return randNum, nil
}

// 默认[minInt,maxInt]区间的值
func generateDefaultInt() (int, error) {
	return generateRangeInt(minInt, maxInt)
}

// 获取[minInt，min]区间的值
func generateMinRangeInt(min int) (int, error) {
	return generateRangeInt(minInt, min)
}

// 获取[max，maxInt]区间的值
func generateMaxRangeInt(max int) (int, error) {
	return generateRangeInt(max, maxInt)
}

// 获取赋负值[minInt，0]
func generateNegativeInt() (int, error) {
	return generateRangeInt(minInt, 0)
}

func generateNonComplianceInt(intRule *IntRule, idx int) (int, error) {
	switch idx {
	case 0:
		return generateMinRangeInt(intRule.Min)
	case 1:
		return generateMaxRangeInt(intRule.Max)
	case 2:
		return generateNegativeInt()
	}
	return 0, nil
}

type IntRule struct {
	Min int
	Max int
}

func (s *IntRule) GetParamType() ParamType {
	return Int
}

func (s *IntRule) GetNonComplianceCount() int {
	return nonComplianceCount
}

func (s *IntRule) GetNext() ParamLimit {
	return nil
}
