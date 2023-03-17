package v1

import "net/http"

type API struct {
	Path   string
	Method string
	Func   http.HandlerFunc
}

type ActDeleted struct {
	Deleted bool `json:"deleted"`
}

func NewAPI(path string, method string, handlerFunc http.HandlerFunc) API {
	return API{path, method, handlerFunc}
}
