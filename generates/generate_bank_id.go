package generates

import "github.com/srlemon/gen-id/generator"

// 生成银行卡号
func generatorBankID() string {
	g := generator.GeneratorData{}
	return g.GeneratorBankID()
}

type BankIdRule struct {
}

func (s *BankIdRule) GetParamType() ParamType {
	return BankID
}

func (s *BankIdRule) GetNonComplianceCount() int {
	return 0
}

func (s *BankIdRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *BankIdRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorBankID())

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
