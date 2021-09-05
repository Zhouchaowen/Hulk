package generates

import (
	"Hulk/utils"
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
		Max:    122,
		Min:    95,
		MaxLen: 12,
		MinLen: 5,
	}
	g2 := IntRule{
		Min: 18,
		Max: 130,
	}
	g3 := FloatRule{
		Min:    50,
		Max:    100,
		Retain: 2,
	}
	g4 := MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
		},
	}
	g5 := ArrayRule{
		Len:  2,
		Type: &g4,
	}
	ic := RequestConfig{
		Param: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
			"map":   &g4,
			"arr":   &g5,
		},
	}
	Generator("", ic.Param)
}

func TestWriteCsv(t *testing.T) {
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
	g4 := MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"money": &g3,
		},
	}
	g5 := ArrayRule{
		Len:  2,
		Type: &g4,
	}
	ic := RequestConfig{
		Param: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"money": &g3,
			"map":   &g4,
			"arr":   &g5,
		},
	}
	var buf = make([]interface{}, 100)
	for i := 0; i < 2; i++ {
		buf[i] = Generator("", ic.Param)
	}

	err := utils.WriteJson("D:\\GOPROJECTS\\src\\Hulk\\data.json", buf)
	if err != nil {
		panic(err)
	}
}

func TestReadCsv(t *testing.T) {
	var buf = make([]map[string]interface{}, 100)
	err := utils.ReadJson("D:\\GOPROJECTS\\src\\Hulk\\data.json", &buf)
	if err != nil {
		panic(err)
	}
	t.Log(buf[0])
}