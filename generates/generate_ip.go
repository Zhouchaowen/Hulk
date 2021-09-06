package generates

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成IP
func generatorIP() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

type IpRule struct {
}

func (s *IpRule) GetParamType() ParamType {
	return IP
}

func (s *IpRule) GetNonComplianceCount() int {
	return 0
}

func (s *IpRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}
