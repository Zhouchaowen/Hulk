package generates

import "reflect"

type MapRule struct {
	Types map[string]ParamLimit
}

func (s *MapRule) GetParamType() string {
	return reflect.Map.String()
}

func (s *MapRule) IsParent() bool {
	return true
}

func (s *MapRule) GetNext() ParamLimit {
	return nil
}