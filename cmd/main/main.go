package main

import "apinodos/src/web/handlers"

func main() {

	drawflow := handlers.CreateSaveDrawflow()
	code := handlers.CreateCodeGenerator()
	mux := Routes(drawflow, code)
	server := NewServer(mux)
	server.listen()

}
