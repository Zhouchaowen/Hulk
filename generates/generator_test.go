package generates

import (
	"Hulk/utils"
	"bytes"
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
	_ = EmailRule{}
	_ = BankIdRule{}
	_ = IdCartRule{}
	_ = PhoneRule{}
	_ = IpRule{}
	_ = TimeRule{}
	_ = MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
		},
	}
	_ = ArrayRule{
		Len:  2,
		Type: &g4,
	}

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
	_ = ArrayRule{
		Len:  2,
		Type: &g4,
	}
	var buf bytes.Buffer
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
	_ = EmailRule{}
	_ = BankIdRule{}
	_ = IdCartRule{}
	_ = PhoneRule{}
	_ = IpRule{}
	_ = TimeRule{}
	_ = MapRule{
		Types: map[string]ParamLimit{
			"name":  &g1,
			"age":   &g2,
			"score": &g3,
		},
	}
	_ = ArrayRule{
		Len:  2,
		Type: &g4,
	}
	_ = `{
    "id":1,
    "agreement":"http",
    "name":"zcw",
    "addr":"http:1.1.1.1",
    "method":"GET",
    "request_config":{
        "int":{
            "param_type":2,
            "param":{
                "max":122,
                "min":33
            }
        },
        "float":{
            "param_type":3,
            "param":{
                "max":122,
                "min":33,
                "retain":2
            }
        },
        "string":{
            "param_type":4,
            "param":{
                "max":122,
                "min":33,
                "maxLen":5,
                "minLen":3
            }
        },
        "array":{
            "param_type":5,
            "param":{
                "len":3,
                "type":{
                    "param_type":4,
                    "param":{
                        "max":122,
                        "min":33,
                        "maxLen":5,
                        "minLen":3
                    }
                }
            }
        },
        "email":{
            "param_type":7,
            "param":{
                "Customized":"123",
                "Prefix":"12",
                "Suffix":"12"
            }
        },
        "address":{
            "param_type":8,
            "param":{}
        },
        "bankId":{
            "param_type":9,
            "param":{
                "Customized":"123",
                "Prefix":"12",
                "Suffix":"12"
            }
        },
        "idCart":{
            "param_type":10,
            "param":{
                "Customized":"123"
            }
        },
        "phone":{
            "param_type":11,
            "param":{}
        },
        "ip":{
            "param_type":12,
            "param":{
                "Customized":"10.2.0.10",
                "Prefix":"12",
                "Suffix":"12"
            }
        },
        "time":{
            "param_type":13,
            "param":{}
        }
    }
}`
}
