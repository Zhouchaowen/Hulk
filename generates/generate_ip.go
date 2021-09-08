package generates

import (
	"fmt"
	"math/rand"
)

// 生成IP
func generatorIP(s *IpRule) string {
	if s.Customized != "" || len(s.Customized) != 0 {
		return s.Customized
	}

	var ip string
	if s.isIpV4 {
		ip = generatorIPV4()
	} else {
		ip = generatorIPV6()
	}

	if s.Prefix != "" || len(s.Prefix) != 0 {
		ip = ip + s.Prefix
	}
	if s.Suffix != "" || len(s.Suffix) != 0 {
		ip = ip + s.Suffix
	}
	return ip
}

func generatorIPV4() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func generatorIPV6() string {
	return ""
}

func generatorNotIP() string {
	num, _ := generateRangeInt(255, 255*2)
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), num, rand.Intn(255), rand.Intn(255))
}

func generatorNotIPChar() string {
	num := generateOneChar()
	return fmt.Sprintf("%d.%s.%d.%d", rand.Intn(255), string(num), rand.Intn(255), rand.Intn(255))
}

func generateNonComplianceIp(ipRule *IpRule, idx int) string {
	switch idx {
	case 0:
		return generatorNotIP()
	case 1:
		return generatorNotIPChar()
	}
	return ""
}

type IpRule struct {
	Customized string
	isIpV4     bool
	Prefix     string
	Suffix     string
}

func (s *IpRule) GetParamType() ParamType {
	return IP
}

func (s *IpRule) GetNonComplianceCount() int {
	return 2
}

func (s *IpRule) GetNonComplianceOtherTypes() []ParamType {
	if s.Customized != "" || len(s.Customized) != 0 {
		return []ParamType{}
	}
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *IpRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorIP(s))
	res = append(res, generatorNotIP())
	res = append(res, generatorNotIPChar())

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
