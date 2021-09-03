package generates

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func generateRangeFloat(min, max float64, retain int) (float64, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	rand.Seed(time.Now().Unix())
	t := min + (max-min)*rand.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(retain)+"f", t), 64)
}

type FloatRule struct {
	Min    float64
	Max    float64
	Retain int
}

func (s *FloatRule) GetParamType() string {
	return reflect.Float64.String()
}

func (s *FloatRule) IsParent() bool {
	return false
}

func (s *FloatRule) GetNext() ParamLimit {
	return nil
}
