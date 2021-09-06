package generates

import (
	"github.com/srlemon/gen-id/generator"
	"github.com/srlemon/gen-id/metadata"
	"github.com/srlemon/gen-id/utils"
)

// 生成手机号码
func generatorPhone() string {
	g := generator.GeneratorData{}
	return g.GeneratorPhone()
}

func generatorPhoneMin11() string {
	suffix, _ := generateFixedString(minChar0, maxCHar9, 7)
	return metadata.MobilePrefix[utils.RandInt(0, generator.MobilePrefix)] + suffix
}

func generatorPhoneMax11() string {
	str, _ := generateFixedString(minChar0, maxCHar9, 3)
	g := generator.GeneratorData{}
	return g.GeneratorPhone() + str
}

func generatorNotPhone() string {
	str, _ := generateFixedString(minChar0, maxCHar9, 11)
	return str
}

func generatorNotPhoneChar() string {
	char := generateOneChar()
	suffix, _ := generateFixedString(minChar0, maxCHar9, 7)
	return metadata.MobilePrefix[utils.RandInt(0, generator.MobilePrefix)] + string(char) + suffix
}

func generateNonCompliancePhone(phoneRule *PhoneRule, idx int) (string, error) {
	switch idx {
	case 0:
		return generatorPhoneMin11(), nil
	case 1:
		return generatorPhoneMax11(), nil
	case 2:
		return generatorNotPhone(), nil
	case 3:
		return generatorNotPhoneChar(), nil
	}
	return "", nil
}

type PhoneRule struct {
}

func (s *PhoneRule) GetParamType() ParamType {
	return Phone
}

func (s *PhoneRule) GetNonComplianceCount() int {
	return 4
}

func (s *PhoneRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
	}
}

func (s *PhoneRule) GetParams() []interface{} {
	var res []interface{}
	res = append(res, generatorPhone())
	res = append(res, generatorPhoneMin11())
	res = append(res, generatorPhoneMax11())
	res = append(res, generatorNotPhone())
	res = append(res, generatorNotPhoneChar())
	otherTypes := s.GetNonComplianceOtherTypes()
	for i, _ := range otherTypes {
		param, _ := generatorNonCompliance(otherTypes[i])
		res = append(res, param)
	}
	return res
}
