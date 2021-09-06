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
