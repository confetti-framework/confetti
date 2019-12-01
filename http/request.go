package http

import (
	"io/ioutil"
	"lanvard/foundation"
	"net/http"
)

type Request struct {
	App      foundation.Application
	Original http.Request
	body     []byte
}

func NewRequest(app foundation.Application, request http.Request) Request {
	return Request{App: app, Original: request}
}

func (r Request) Content() string {
	if r.body == nil {
		bodyBytes, err := ioutil.ReadAll(r.Original.Body)
		r.body = bodyBytes
		if err != nil {
			panic(err)
		}
	}

	return string(r.body)
}

func (r Request) SetContent(content string) Request {
	r.body = []byte(content)

	return r
}

func (r Request) GetMethod() string {
	return r.Original.Method
}
