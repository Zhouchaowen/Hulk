package generates

import "github.com/srlemon/gen-id/generator"

// 生成身份证号
func generatorIDCart() string {
	g := generator.GeneratorData{}
	g.GeneratorIDCart()
	return g.IDCard
}

type IdCartRule struct {
}

func (s *IdCartRule) GetParamType() ParamType {
	return IDCart
}

func (s *IdCartRule) GetNonComplianceCount() int {
	return 0
}

func (s *IdCartRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}
