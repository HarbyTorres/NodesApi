package main

import (
	"apinodos/src/repository"
	"apinodos/src/web/handlers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(dfh *handlers.SaveDrawflow, cdg *handlers.CodeGenerator, gdf *repository.GetDrawflow) *chi.Mux {
	mux := chi.NewMux()
	/*
		mux.Use(
			middleware.Logger,
			middleware.Recoverer,
		)
	*/
	mux.Get("/hello", HelloWorld)
	mux.Post("/drawflow", dfh.SaveDrawflow)
	mux.Post("/code", cdg.GenerateCodeHandler)
	mux.Get("/drawflow", gdf.GetFile)
	return mux
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	res := map[string]interface{}{"message": "Hello World"}

	_ = json.NewEncoder(w).Encode(res)

}
