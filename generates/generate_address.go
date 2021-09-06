package generates

import "github.com/srlemon/gen-id/generator"

// 生成地址
func generateAddress() string {
	g := generator.GeneratorData{}
	return g.GeneratorAddress()
}

// 生成随机获取城市和地址
func generatorProvinceAdnCityRand() string {
	g := generator.GeneratorData{}
	return g.GeneratorProvinceAdnCityRand()
}

type AddressRule struct {
}

func (s *AddressRule) GetParamType() ParamType {
	return Address
}

func (s *AddressRule) GetNonComplianceCount() int {
	return 0
}

func (s *AddressRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *AddressRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generateAddress())

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
