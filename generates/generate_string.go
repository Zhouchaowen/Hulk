package generates

import (
	"bytes"
	"errors"
	"math/rand"
)

const (
	minChar    = 33  // !
	maxChar    = 126 // ~
	minChara   = 'a'
	maxCharz   = 'z'
	minCharA   = 'A'
	maxCharZ   = 'Z'
	minChar0   = '0'
	maxCHar9   = '9'
	minCharLen = 0
	maxCharLen = 255
)

func generateRangeString(min, max int, minLen, maxLen int) (string, error) {
	if max <= min || maxLen <= minLen || min < minChar || max > maxChar {
		return "", errors.New("max cannot be less than min or min cannot be less 33 or max cannot be greater than 126")
	}
	var buf bytes.Buffer
	n := rand.Intn(maxLen-minLen) + minLen
	for i := 0; i < n; i++ {
		b := rand.Intn(max-min) + min
		buf.WriteByte(byte(b))
	}
	return buf.String(), nil
}

func generateFixedString(min, max int, n int) (string, error) {
	if max <= min || min < minChar || max > maxChar {
		return "", errors.New("max cannot be less than min or min cannot be less 33 or max cannot be greater than 126")
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		b := rand.Intn(max-min) + min
		buf.WriteByte(byte(b))
	}
	return buf.String(), nil
}

func generateOneChar() byte {
	return byte(rand.Intn(maxChar-minChar) + minChar)
}

func generateDefaultString() (string, error) {
	return generateRangeString(minChara, maxCharz, minCharLen, maxCharLen)
}

// 获取[minChar，min]字符区间的值
func generateMinCharRangeString(min, minLen, maxLen int) (string, error) {
	return generateRangeString(minChar, min, minLen, maxLen)
}

// 获取[max，maxChar]字符区间的值
func generateMaxCharRangeString(max int, minLen, maxLen int) (string, error) {
	return generateRangeString(max, maxChar, minLen, maxLen)
}

// 获取[minCharLen，minLen]区间长度的值
func generateMinRangeString(min, max, minLen int) (string, error) {
	return generateRangeString(min, max, minCharLen, minLen)
}

// 获取[maxLen，maxCharLen]区间长度的值
func generateMaxRangeString(min, max, maxLen int) (string, error) {
	return generateRangeString(min, max, maxLen, maxCharLen)
}

func generateNonComplianceString(s *StringRule, idx int) (string, error) {
	switch idx {
	case 0:
		return generateMinCharRangeString(s.Min, s.MinLen, s.MaxLen)
	case 1:
		return generateMaxCharRangeString(s.Max, s.MinLen, s.MaxLen)
	case 2:
		return generateMinRangeString(s.Min, s.Max, s.MinLen)
	case 3:
		return generateMaxRangeString(s.Min, s.Max, s.MaxLen)
	}
	return "", nil
}

type StringRule struct {
	Min    int
	Max    int
	MinLen int
	MaxLen int
}

func (s *StringRule) GetParamType() ParamType {
	return String
}

func (s *StringRule) GetNonComplianceCount() int {
	return 4
}

func (s *StringRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *StringRule) GetParams() []interface{} {
	var res []interface{}
	str, _ := generateRangeString(s.Min, s.Max, s.MinLen, s.MaxLen)
	res = append(res, str)
	str, _ = generateMinCharRangeString(s.Min, s.MinLen, s.MaxLen)
	res = append(res, str)
	str, _ = generateMaxCharRangeString(s.Max, s.MinLen, s.MaxLen)
	res = append(res, str)
	str, _ = generateMinRangeString(s.Min, s.Max, s.MinLen)
	res = append(res, str)
	str, _ = generateMaxRangeString(s.Min, s.Max, s.MaxLen)
	res = append(res, str)
	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
