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

func (s *IdCartRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorIDCart)

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		param, _ := generatorNonCompliance(otherTypes[i])
		res = append(res, param)
	}
	return res
}
