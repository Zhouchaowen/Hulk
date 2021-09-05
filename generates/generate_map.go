package generates

type MapRule struct {
	Types map[string]ParamLimit
}

func (s *MapRule) GetParamType() ParamType {
	return Map
}

func (s *MapRule) GetNonComplianceCount() int {
	return 0
}

func (s *MapRule) GetNext() ParamLimit {
	return nil
}
