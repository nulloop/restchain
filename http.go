package restchain

import (
	"net/http"

	"github.com/go-chi/chi"
)

type PostRouter interface {
	Post(path string, handler http.HandlerFunc)
}

func NewPostRouter() PostRouter {
	return chi.NewRouter()
}
