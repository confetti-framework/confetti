package http

import (
	"net/http"
)

type Kernel interface {
	Handle(request *http.Request) http.ResponseWriter
}
