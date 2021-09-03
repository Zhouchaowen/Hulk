package generates

import (
	"bytes"
	"errors"
	"math/rand"
	"reflect"
)

func generateRangeString(min, max int, minLen, maxLen int) (string, error) {
	if max <= min || maxLen <= minLen {
		return "", errors.New("max cannot be less than min")
	}
	var buf bytes.Buffer
	n := rand.Intn(maxLen-minLen) + minLen
	for i := 0; i < n; i++ {
		b := rand.Intn(max-min) + min
		buf.WriteByte(byte(b))
	}
	return buf.String(), nil
}

func generateOneChar() byte {
	return byte(rand.Intn(128))
}

type StringRule struct {
	Min    int
	Max    int
	MinLen int
	MaxLen int
}

func (s *StringRule) GetParamType() string {
	return reflect.String.String()
}

func (s *StringRule) IsParent() bool {
	return false
}

func (s *StringRule) GetNext() ParamLimit {
	return nil
}
