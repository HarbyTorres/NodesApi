package handlers

import (
	"apinodos/internal/database"
	"apinodos/src/models"
	"apinodos/src/services"
	"encoding/json"
	"net/http"
)

type SaveDrawflow struct {
}

func (*SaveDrawflow) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	//var node models.Node
	var drawflowSvs services.DrawflowSvc
	var draw models.CreateDrawflow
	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&draw)
	payload := drawflowSvs.Create(draw)

	response, err := database.NewQuery(payload)
	if err != nil {
		panic(err)
	}
	_ = json.NewEncoder(w).Encode(response)
}

func CreateSaveDrawflow() *SaveDrawflow {
	return &SaveDrawflow{}
}
