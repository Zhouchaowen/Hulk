package send

import (
	"net/http"
)

const (
	HttpHeaderAccept      string = "Accept"
	HttpHeaderContentType string = "Content-Type"
	HttpHeaderUserAgent   string = "User-Agent"
	HttpHeaderDate        string = "Date"
)

type ContentType int

const (
	ContentTypeJson ContentType = iota
	ContentTypeFrom
	ContentTypeStream
	ContentTypeXml
	ContentTypeText
)

func (k ContentType) String() string {
	if int(k) < len(ContentTypeNames) {
		return ContentTypeNames[k]
	}
	return ContentTypeNames[0]
}

var ContentTypeNames = []string{
	ContentTypeJson:   "application/json; charset=UTF-8",
	ContentTypeFrom:   "application/x-www-form-urlencoded; charset=UTF-8",
	ContentTypeStream: "application/octet-stream; charset=UTF-8",
	ContentTypeXml:    "application/xml; charset=UTF-8",
	ContentTypeText:   "application/text; charset=UTF-8",
}

func initHeader(r *http.Request, param map[string]string, contentType ContentType) {
	r.Header.Set("Content-Type", contentType.String())
	for k, v := range param {
		r.Header.Set(k, v)
	}
}
