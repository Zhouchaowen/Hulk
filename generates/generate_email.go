package generates

import (
	"github.com/srlemon/gen-id/generator"
	"strings"
)

// 生成邮箱
func generatorEmail() string {
	g := generator.GeneratorData{}
	return g.GeneratorEmail()
}

func generatorNotEmail() string {
	g := generator.GeneratorData{}
	emails := strings.Split(g.GeneratorEmail(), "@")
	return emails[0] + "#" + emails[1]
}

type EmailRule struct {
}

func (s *EmailRule) GetParamType() ParamType {
	return Email
}

func (s *EmailRule) GetNonComplianceCount() int {
	return 0
}

func (s *EmailRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *EmailRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorEmail())
	res = append(res, generatorNotEmail())

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
