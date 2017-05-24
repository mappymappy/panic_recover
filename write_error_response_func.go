package recover

import (
	"net/http"
)

type WriteErrorResponseFunc func(http.ResponseWriter, *http.Request, interface{})

func DefaultWriteErrorResponseFunc(rw http.ResponseWriter, req *http.Request, panicError interface{}) {
	// use default content-type text/plain;char-set=utf8
	rw.WriteHeader(http.StatusInternalServerError)
}
