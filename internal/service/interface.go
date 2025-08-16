package service

import "net/http"

type Servicer interface {
	Create(*http.Request) ([]byte, int)
	Read(*http.Request) ([]byte, int)
	Update(*http.Request) ([]byte, int)
	Delete(*http.Request) ([]byte, int)
}
