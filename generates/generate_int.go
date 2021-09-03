package generates

import (
	"errors"
	"math/rand"
	"reflect"
)

func generateRangeBool() bool {
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

func generateRangeInt(min, max int) (int, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	randNum := rand.Intn(max-min) + min
	return randNum, nil
}

type IntRule struct {
	Min int
	Max int
}

func (s *IntRule) GetParamType() string {
	return reflect.Int.String()
}

func (s *IntRule) IsParent() bool {
	return false
}

func (s *IntRule) GetNext() ParamLimit {
	return nil
}
