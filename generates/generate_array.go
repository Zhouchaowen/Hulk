package generates

type ArrayRule struct {
	Len  int
	Type ParamLimit
}

func (s *ArrayRule) GetParamType() ParamType {
	return Array
}

func (s *ArrayRule) GetNonComplianceCount() int {
	return 0
}

func (s *ArrayRule) GetNonComplianceParamTypes() []ParamType {
	return []ParamType{}
}
