package generates

type MapRule struct {
	Types map[string]ParamLimit `json:"types"`
}

func (s *MapRule) GetParamType() ParamType {
	return Map
}

func (s *MapRule) GetNonComplianceCount() int {
	return 0
}

func (s *MapRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{}
}
