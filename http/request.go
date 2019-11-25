package http

import (
	"io/ioutil"
	"lanvard/foundation"
	"net/http"
)

type RequestStruct struct {
	App      foundation.Application
	original http.Request
	body     []byte
}

func Request(app foundation.Application, request http.Request) RequestStruct {
	return RequestStruct{App: app, original: request}
}

func (r RequestStruct) Original() http.Request {
	return r.original
}

func (r RequestStruct) Content() string {
	if r.body == nil {
		bodyBytes, err := ioutil.ReadAll(r.original.Body)
		r.body = bodyBytes
		if err != nil {
			panic(err)
		}
	}

	return string(r.body)
}

func (r RequestStruct) SetContent(content string) RequestStruct {
	r.body = []byte(content)

	return r
}
