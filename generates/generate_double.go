package generates

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateRangeDouble(min, max float64, retain int) (float64, error) {
	if max <= min {
		return 0, errors.New("max cannot be less than min")
	}
	rand.Seed(time.Now().Unix())
	t := min + (max-min)*rand.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(retain)+"f", t), 64)
}

func GenerateRangeBool() bool {
	rand.Seed(time.Now().Unix())
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

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
