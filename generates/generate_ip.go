package generates

import (
	"fmt"
	"math/rand"
)

// 生成IP
func generatorIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func generatorIPV4() {

}

func generatorIPV6() {

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
}

func (s *IpRule) GetParamType() ParamType {
	return IP
}

func (s *IpRule) GetNonComplianceCount() int {
	return 2
}

func (s *IpRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}
