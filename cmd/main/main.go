package main

import (
	"apinodos/src/web/handlers"
)

func main() {

	drawflow := handlers.CreateDrawflowHandler()
	code := handlers.CreateCodeGenerator()
	mux := Routes(drawflow, code)
	server := NewServer(mux)
	server.listen()

}
