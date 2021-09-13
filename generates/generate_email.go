package generates

import (
	"github.com/srlemon/gen-id/generator"
	"strings"
)

// 生成邮箱
func generatorEmail(s *EmailRule) string {
	if s.Customized != "" || len(s.Customized) != 0 {
		return s.Customized
	}

	g := generator.GeneratorData{}
	emails := strings.Split(g.GeneratorEmail(), "@")

	buf := emails[0] + "@" + emails[1]
	if s.Prefix != "" || len(s.Prefix) != 0 {
		buf = s.Prefix + "@" + emails[1]
	}

	if s.Suffix != "" || len(s.Suffix) != 0 {
		buf = emails[0] + "@" + s.Suffix
	}
	return buf
}

func generatorNotEmail() string {
	g := generator.GeneratorData{}
	emails := strings.Split(g.GeneratorEmail(), "@")
	return emails[0] + "#" + emails[1]
}

// 注意生成条件互斥，Customized优先级最高，Prefix，Suffix最终只执行最后一个，
type EmailRule struct {
	MustCustomized bool `json:"must_customized"`
	Customized     string
	Prefix         string
	Suffix         string
}

func (s *EmailRule) GetParamType() ParamType {
	return Email
}

func (s *EmailRule) GetNonComplianceCount() int {
	if len(s.Customized) != 0 {
		return 0
	}
	return 0
}

func (s *EmailRule) GetNonComplianceOtherTypes() []ParamType {
	if s.MustCustomized {
		return []ParamType{}
	}
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *EmailRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorEmail(s))
	res = append(res, generatorNotEmail())

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
