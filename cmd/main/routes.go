package main

import (
	"apinodos/src/web/handlers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(dfh *handlers.SaveDrawflow, cdg *handlers.CodeGenerator) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	mux.Get("/hello", HelloWorld)
	mux.Post("/drawflow", dfh.SaveDrawflow)
	mux.Post("/code", cdg.GenerateCodeHandler)
	return mux
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	res := map[string]interface{}{"message": "Hello World"}

	_ = json.NewEncoder(w).Encode(res)

}
