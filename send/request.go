package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

func SetParam(param map[string]interface{}, contentType ContentType) (*bytes.Reader, error) {
	switch contentType {
	case ContentTypeFrom:
		data := url.Values{}
		for k, v := range param {
			if reflect.TypeOf(v).Kind() != reflect.String {
				b, _ := json.Marshal(v)
				data.Set(k, string(b))
				continue
			}
			data.Set(k, v.(string))
		}
		return bytes.NewReader([]byte(data.Encode())), nil
	case ContentTypeJson:
		bytesData, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(bytesData), nil
	}
	return nil, fmt.Errorf("no content type")
}
