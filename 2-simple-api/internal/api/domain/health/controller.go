package health

import (
	"github.com/go-chi/chi"
	"net/http"
)

func GetController() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getHandler)
	return r
}

func getHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
