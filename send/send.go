package send

import (
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
	method      string
	url         string
	contentType ContentType
	header      map[string]string
	param       map[string]interface{}
	req         *http.Request
	client      *http.Client
}

func (s *HttpRequest) Init() {
	// 设置bodyFrom
	reader, err := SetParam(s.param, s.contentType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(s.method, s.url, reader)
	if err != nil {
		panic(err)
	}
	// 设置Header
	initHeader(req, s.header, s.contentType)
	s.req = req
	s.client = http.DefaultClient
}

func (s *HttpRequest) Send(param map[string]interface{}) {
	if s.client == nil {
		panic("this client is nil,you must execute init first")
	}
	// 设置req
	reader, err := SetParam(param, s.contentType)
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

	resp, err := s.client.Do(s.req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取请求
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(strconv.Itoa(resp.StatusCode) + " " + string(rbody))
	}
	fmt.Println(string(rbody))
}

type HttpsRequest struct {
	method string
	header map[string]string
	param  map[string]interface{}
}

func (s *HttpsRequest) Send() {

}
