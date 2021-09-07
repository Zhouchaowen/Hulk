package send

import "testing"

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
		"zdnsuser":      "b957754b403d261580c870fa78a4e0a1",
		"resource_type": "top_category",
		"attrs": map[string]interface{}{
			"limit": 1,
		},
	}
	req.Init()
	req.Send(p)
}
