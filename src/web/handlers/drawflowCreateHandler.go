package handlers

import (
	"apinodos/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type SaveDrawflow struct {
}

func (*SaveDrawflow) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	//var node models.Node
	var draw models.DrawflowMap

	var data map[string]interface{}

	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&data)

	err := json.Unmarshal([]byte(), &draw)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n json object:::: %v", draw)
	fmt.Println(data)
	_ = json.NewEncoder(w).Encode(data)
}

func CreateSaveDrawflow() *SaveDrawflow {
	return &SaveDrawflow{}
}
