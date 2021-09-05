package generates

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	maxFloat = float64(1 << 63)
	minFloat = -float64(1 << 63)
	retain   = 2
)

func generateRangeFloat(min, max float64, retain int) (float64, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	t := min + (max-min)*rand.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(retain)+"f", t), 64)
}

func generateDefaultFloat() (float64, error) {
	return generateRangeFloat(minFloat, maxFloat, retain)
}

type FloatRule struct {
	Min    float64
	Max    float64
	Retain int
}

func (s *FloatRule) GetParamType() ParamType {
	return Float64
}

func (s *FloatRule) GetNonComplianceCount() int {
	return 0
}

func (s *FloatRule) GetNext() ParamLimit {
	return nil
}
