package gooky

import (
	"net/http"
)

type Request struct {
	HttpRequest *http.Request
}
