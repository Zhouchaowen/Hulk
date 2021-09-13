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

func generateString(s *StringRule) (string, error) {
	if s.Customized != "" || len(s.Customized) != 0 {
		return s.Customized, nil
	}

	res, err := generateRangeString(s.Min, s.Max, s.MinLen, s.MaxLen)
	if err != nil {
		return "", err
	}

	if s.Prefix != "" || len(s.Prefix) != 0 {
		res = res + s.Prefix
	}
	if s.Suffix != "" || len(s.Suffix) != 0 {
		res = res + s.Suffix
	}
	return res, nil
}

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
func generateMinCharRangeString(s *StringRule) (string, error) {
	s.Max = s.Min
	s.Min = minChar
	return generateString(s)
}

// 获取[max，maxChar]字符区间的值
func generateMaxCharRangeString(s *StringRule) (string, error) {
	s.Min = s.Max
	s.Max = maxChar
	return generateString(s)
}

// 获取[minCharLen，minLen]区间长度的值
func generateMinRangeString(s *StringRule) (string, error) {
	s.MaxLen = s.MinLen
	s.MinLen = minCharLen
	return generateString(s)
}

// 获取[maxLen，maxCharLen]区间长度的值
func generateMaxRangeString(s *StringRule) (string, error) {
	s.MinLen = s.MaxLen
	s.MaxLen = maxCharLen
	return generateString(s)
}

func generateNonComplianceString(s *StringRule, idx int) (string, error) {
	switch idx {
	case 0:
		return generateMinCharRangeString(s)
	case 1:
		return generateMaxCharRangeString(s)
	case 2:
		return generateMinRangeString(s)
	case 3:
		return generateMaxRangeString(s)
	}
	return "", nil
}

type StringRule struct {
	Min            int    `json:"min"`
	Max            int    `json:"max"`
	MinLen         int    `json:"min_len"`
	MaxLen         int    `json:"max_len"`
	MustCustomized bool   `json:"must_customized"`
	Customized     string `json:"customized"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
}

func (s *StringRule) GetParamType() ParamType {
	return String
}

func (s *StringRule) GetNonComplianceCount() int {
	if len(s.Customized) != 0 {
		return 0
	}
	return 4
}

func (s *StringRule) GetNonComplianceOtherTypes() []ParamType {
	if s.MustCustomized {
		return []ParamType{}
	}

	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *StringRule) GetParams() []interface{} {
	var res []interface{}
	str, _ := generateString(s)
	res = append(res, str)
	str, _ = generateMinCharRangeString(s)
	res = append(res, str)
	str, _ = generateMaxCharRangeString(s)
	res = append(res, str)
	str, _ = generateMinRangeString(s)
	res = append(res, str)
	str, _ = generateMaxRangeString(s)
	res = append(res, str)
	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
