package send

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	HTTP   string = "http://"
	HTTPS  string = "https://"
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

type HttpRequest struct {
	Method      string
	Url         string
	ContentType ContentType
	Header      map[string]string
	Param       map[string]interface{}
	req         *http.Request
	client      *http.Client
}

func (s *HttpRequest) Init() {
	reader := bytes.NewReader([]byte(""))

	req, err := http.NewRequest(s.Method, s.Url, reader)
	if err != nil {
		panic(err)
	}
	// 设置Header
	initHeader(req, s.Header, s.ContentType)
	s.req = req
	s.client = http.DefaultClient
}

func (s *HttpRequest) Send(param map[string]interface{}) {
	if s.client == nil {
		panic("this client is nil,you must execute init first")
	}
	// 设置req
	reader, err := SetParam(param, s.ContentType)
	if reader != nil {
		s.req.Body = ioutil.NopCloser(reader)
		s.req.ContentLength = int64(reader.Len())
		snapshot := *reader
		s.req.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return io.NopCloser(&r), nil
		}
		if s.req.GetBody != nil && s.req.ContentLength == 0 {
			s.req.Body = http.NoBody
			s.req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		}
	}

	// 执行请求
	resp, err := s.client.Do(s.req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取请求响应
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(strconv.Itoa(resp.StatusCode) + " " + string(data))
	}
	fmt.Println(string(data))
}

type HttpsRequest struct {
	method string
	header map[string]string
	param  map[string]interface{}
}

func (s *HttpsRequest) Send() {

}
