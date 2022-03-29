package main

import "apinodos/src/web/handlers"

func main() {

	drawflow := handlers.CreateSaveDrawflow()

	mux := Routes(drawflow)
	server := NewServer(mux)
	server.listen()

}
