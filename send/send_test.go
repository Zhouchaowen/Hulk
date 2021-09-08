package send

import (
	"encoding/json"
	"log"
	"testing"
)

func TestHttpRequest_Send(t *testing.T) {
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
