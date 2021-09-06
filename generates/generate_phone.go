package generates

import "github.com/srlemon/gen-id/generator"

// 生成手机号码
func generatorPhone() string {
	g := generator.GeneratorData{}
	return g.GeneratorPhone()
}

type PhoneRule struct {
}

func (s *PhoneRule) GetParamType() ParamType {
	return Phone
}

func (s *PhoneRule) GetNonComplianceCount() int {
	return 0
}

func (s *PhoneRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}
