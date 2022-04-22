package handlers

import (
	"apinodos/src/models"
	"apinodos/src/services"
	"encoding/json"
	"net/http"
)

type DrawflowHandler struct {
}

func (*DrawflowHandler) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	//var node models.Node
	var drawflowSvs services.DrawflowSvc
	var draw models.CreateDrawflow
	encoder := json.NewDecoder(r.Body)
	err := encoder.Decode(&draw)
	if err != nil {
		panic(err)
	}
	response, err := drawflowSvs.Create(draw)
	if err != nil {
		panic(err)
	}
	_ = json.NewEncoder(w).Encode(string(response))
}

func (*DrawflowHandler) GetDrawflows(w http.ResponseWriter, r *http.Request) {
	//var node models.Node
	var drawflowSvs services.DrawflowSvc

	response, err := drawflowSvs.GetAll()
	if err != nil {
		panic(err)
	}
	_ = json.NewEncoder(w).Encode(string(response))
}

func CreateDrawflowHandler() *DrawflowHandler {
	return &DrawflowHandler{}
}
