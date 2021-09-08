package generates

import "github.com/srlemon/gen-id/generator"

// 生成身份证号
func generatorIDCart(s *IdCartRule) string {
	if s.Customized != "" || len(s.Customized) != 0 {
		return s.Customized
	}

	g := generator.GeneratorData{}
	g.GeneratorIDCart()
	return g.IDCard
}

type IdCartRule struct {
	Customized string
}

func (s *IdCartRule) GetParamType() ParamType {
	return IDCart
}

func (s *IdCartRule) GetNonComplianceCount() int {
	return 0
}

func (s *IdCartRule) GetNonComplianceOtherTypes() []ParamType {
	if s.Customized != "" || len(s.Customized) != 0 {
		return []ParamType{}
	}
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
