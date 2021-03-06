package generates

import "github.com/srlemon/gen-id/generator"

// 生成银行卡号
func generatorBankID(s *BankIdRule) string {
	if s.Customized != "" || len(s.Customized) != 0 {
		return s.Customized
	}
	g := generator.GeneratorData{}

	buf := []byte(g.GeneratorBankID())
	// 替换前缀
	if s.Prefix != "" {
		t := []byte(s.Prefix)
		for i := 0; i < len(t); i++ {
			buf[i] = t[i]
		}
	}
	// 替换后缀
	if s.Suffix != "" {
		t := []byte(s.Suffix)
		bankIdLen := len(buf)
		j := 0
		for i := bankIdLen - len(t); i < bankIdLen; i++ {
			buf[i] = t[j]
			j++
		}
	}
	return string(buf)
}

type BankIdRule struct {
	MustCustomized bool   `json:"must_customized"`
	Customized     string `json:"customized"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
}

func (s *BankIdRule) GetParamType() ParamType {
	return BankID
}

func (s *BankIdRule) GetNonComplianceCount() int {
	if len(s.Customized) != 0 {
		return 0
	}
	return 0
}

func (s *BankIdRule) GetNonComplianceOtherTypes() []ParamType {
	if s.MustCustomized {
		return []ParamType{}
	}
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *BankIdRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorBankID(s))

	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		if param, err := generatorNonCompliance(otherTypes[i]); err != nil && param != nil {
			res = append(res, param)
		}
	}
	return res
}
