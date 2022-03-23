package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	mux.Get("/hello", HelloWorld)

	return mux
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-Type", "application/json")
	w.Header().Set("done-by", "harby")

	res := map[string]interface{}{"message": "Hello World"}

	_ = json.NewEncoder(w).Encode(res)

}
