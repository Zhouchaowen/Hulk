package generates

import (
	"testing"
)

func TestGenerateAddress(t *testing.T) {
	t.Log(generateAddress())
}

func BenchmarkGenerateAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateAddress()
	}
}

func TestGeneratorBankID(t *testing.T) {
	t.Log(generatorBankID())
}

func TestGeneratorEmail(t *testing.T) {
	t.Log(generatorEmail())
}

func TestGeneratorIDCart(t *testing.T) {
	t.Log(generatorIDCart())
}

func TestGeneratorPhone(t *testing.T) {
	t.Log(generatorPhone())
}

func TestGeneratorIP(t *testing.T) {
	t.Log(generatorIP())
}

func TestGeneratorProvinceAdnCityRand(t *testing.T) {
	t.Log(generatorProvinceAdnCityRand())
}

func TestGeneratorRandDate(t *testing.T) {
	t.Log(generatorRandDate())
}

func TestGenerator(t *testing.T) {
	g1 := StringRule{
		Max:    100,
		Min:    50,
		MaxLen: 12,
		MinLen: 5,
	}
	g2 := IntRule{
		Min: 18,
		Max: 130,
	}
	g3 := FloatRule{
		Min:    5000,
		Max:    100000,
		Retain: 2,
	}
	ic := RequestConfig{
		Param: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"money": &g3,
		},
	}
	t.Log(Generator(ic))
}
