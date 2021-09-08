package generates

import (
	"Hulk/utils"
	"log"
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
	var s = &BankIdRule{}
	t.Log(generatorBankID(s))
}

func TestGeneratorEmail(t *testing.T) {
	var s = &EmailRule{
		Prefix: "zcw",
		Suffix: "126.cn",
	}
	t.Log(generatorEmail(s))
}

func TestGeneratorIDCart(t *testing.T) {
	var s = &IdCartRule{}
	t.Log(generatorIDCart(s))
}

func TestGeneratorPhone(t *testing.T) {
	t.Log(generatorPhone())
}

func TestGeneratorIP(t *testing.T) {
	var s = &IpRule{
		isIpV4: true,
	}
	t.Log(generatorIP(s))
}

func TestGeneratorProvinceAdnCityRand(t *testing.T) {
	t.Log(generatorProvinceAdnCityRand())
}

func TestGeneratorRandDate(t *testing.T) {
	t.Log(generatorRandTime())
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
	g4 := AddressRule{}
	g5 := EmailRule{}
	g6 := BankIdRule{}
	g7 := IdCartRule{}
	g8 := PhoneRule{}
	g9 := IpRule{}
	g10 := TimeRule{}
	g11 := MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
		},
	}
	g12 := ArrayRule{
		Len:  2,
		Type: &g4,
	}
	ic := RequestConfig{
		Param: map[string]ParamLimit{
			"name":     &g1,
			"age":      &g2,
			"score":    &g3,
			"address":  &g4,
			"email":    &g5,
			"BankId":   &g6,
			"IDCart":   &g7,
			"Phone":    &g8,
			"IP":       &g9,
			"birthday": &g10,
			"map":      &g11,
			"arr":      &g12,
		},
	}
	_ = ic

	var p = map[string]ParamLimit{
		"zdnsuser": &StringRule{
			Customized: "94ff8576408972af80af994142e43323",
		},
		"resource_type": &StringRule{
			Customized: "top_category",
		},
		"attrs": &MapRule{
			Types: map[string]ParamLimit{
				"limit": &IntRule{
					Min: 1,
					Max: 10,
				},
			},
		},
	}
	Generator("/Users/zdns/Desktop/Hulk", p)
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

func TestSpreadParams(t *testing.T) {
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
	g4 := AddressRule{}
	g5 := EmailRule{}
	g6 := BankIdRule{}
	g7 := IdCartRule{}
	g8 := PhoneRule{}
	g9 := IpRule{}
	g10 := TimeRule{}
	g11 := MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
		},
	}
	g12 := ArrayRule{
		Len:  2,
		Type: &g11,
	}
	ic := RequestConfig{
		Param: map[string]ParamLimit{
			"name":     &g1,
			"age":      &g2,
			"score":    &g3,
			"address":  &g4,
			"email":    &g5,
			"BankId":   &g6,
			"IDCart":   &g7,
			"Phone":    &g8,
			"IP":       &g9,
			"birthday": &g10,
			"map":      &g11,
			"arr":      &g12,
		},
	}
	params := SpreadParams(ic.Param)
	for i, _ := range params {
		log.Println(params[i])
	}

}
