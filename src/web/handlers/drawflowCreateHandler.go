package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SaveDrawflow struct {
}

func (*SaveDrawflow) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	//var node models.Node

	var data map[string]interface{}

	encoder := json.NewDecoder(r.Body)
	fmt.Print(encoder.Decode(&data))

	home := data["drawflow"]

	fmt.Println(home)
	_ = json.NewEncoder(w).Encode(home)
}

func CreateSaveDrawflow() *SaveDrawflow {
	return &SaveDrawflow{}
}
