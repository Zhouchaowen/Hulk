package send

import (
	"encoding/json"
	"log"
	"testing"
)

func TestHttpRequestPost_Send(t *testing.T) {
	var req = HttpRequest{
		Method:      "POST",
		Url:         "http://10.2.0.153:6021/sg_logstatistics_cmd",
		ContentType: ContentTypeJson,
		Param: map[string]interface{}{
			"zdnsuser":      "b957754b403d261580c870fa78a4e0a1",
			"resource_type": "top_category",
			"attrs": map[string]interface{}{
				"limit": 2,
			},
		},
	}
	var p = map[string]interface{}{
		"zdnsuser":      "94ff8576408972af80af994142e43323",
		"resource_type": "top_category",
		"attrs": map[string]interface{}{
			"limit": 1,
		},
	}
	b, _ := json.Marshal(p)
	log.Printf(string(b))
	req.Init()
	req.Send(p)
}

func TestHttpRequestGet_Send(t *testing.T) {
	var req = HttpRequest{
		Method:      GET,
		Url:         "http://10.2.0.153:6040/sg_policymanager",
		ContentType: ContentTypeFrom,
	}
	var p = map[string]interface{}{
		"zdnsuser":      "fmVKVZ-HTaW2NDvGcs_K4A",
		"resource_type": "white_list",
		"_onlycount":    true,
		"orderby":       "create_time desc",
	}
	b, _ := json.Marshal(p)
	log.Printf(string(b))
	req.Init()
	req.Send(p)
}

//func TestHttpRequestGet_Send(t *testing.T) {
//	var req = HttpRequest{
//		Method: GET,
//		Url:    "http://imianba.cn/details.html",
//	}
//
//	var p = map[string]interface{}{
//		"id": "d6ca54c",
//	}
//
//	b, _ := json.Marshal(p)
//	log.Printf(string(b))
//	req.Init()
//	req.Send(p)
//}
