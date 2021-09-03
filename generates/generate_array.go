package generates

import "reflect"

type ArrayRule struct {
	Len int
	Type ParamLimit
}

func (s *ArrayRule) GetParamType() string {
	return reflect.Array.String()
}

func (s *ArrayRule) IsParent() bool {
	return false
}

func (s *ArrayRule) GetNext() ParamLimit {
	return nil
}
