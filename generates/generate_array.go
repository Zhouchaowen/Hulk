package generates

type ArrayRule struct {
	Len  int        `json:"len"`
	Type ParamLimit `json:"type"`
}

func (s *ArrayRule) GetParamType() ParamType {
	return Array
}

func (s *ArrayRule) GetNonComplianceCount() int {
	return 0
}

func (s *ArrayRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{}
}
