package presenter

import (
	"bytes"
)

type Response struct {
	Code              int         `json:"code"`
	Data              interface{} `json:"data,omitempty"`
	Stream            *bytes.Reader
	StreamFileName    string `json:"stream_file_name,omitempty"`
	StreamContentType string `json:"stream_content_type,omitempty"`
	Errors            error  `json:"errors,omitempty"`
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithStream(fileBytes *[]byte, fileName string, contentType string) *Response {
	reader := bytes.NewReader(*fileBytes)
	r.Stream = reader
	r.StreamFileName = fileName
	r.StreamContentType = contentType
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithError(err error) *Response {
	r.Errors = err
	return r
}
