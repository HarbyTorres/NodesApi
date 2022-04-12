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

	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&draw)

	fmt.Printf("\n\n json object:::: %v", draw)
	_ = json.NewEncoder(w).Encode(draw)
}

func CreateSaveDrawflow() *SaveDrawflow {
	return &SaveDrawflow{}
}
