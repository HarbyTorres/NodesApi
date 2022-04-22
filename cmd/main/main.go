package main

import (
	"apinodos/src/repository"
	"apinodos/src/web/handlers"
)

func main() {

	drawflow := handlers.CreateSaveDrawflow()
	code := handlers.CreateCodeGenerator()
	drawflowget := repository.CreateGetDrawflow()
	mux := Routes(drawflow, code, drawflowget)
	server := NewServer(mux)
	server.listen()

}
