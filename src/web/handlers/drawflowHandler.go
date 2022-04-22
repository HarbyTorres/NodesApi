package handlers

import (
	"apinodos/src/models"
	"apinodos/src/services"
	"encoding/json"
	"net/http"
)

type DrawflowHandler struct {
	drawflowSvs services.DrawflowSvc
}

func (d *DrawflowHandler) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	var draw models.CreateDrawflow
	encoder := json.NewDecoder(r.Body)
	err := encoder.Decode(&draw)
	if err != nil {
		panic(err)
	}
	response, err := d.drawflowSvs.Create(draw)
	if err != nil {
		panic(err)
	}
	_ = json.NewEncoder(w).Encode(response)
}

func (d *DrawflowHandler) GetDrawflows(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	response, err := d.drawflowSvs.GetAll()
	if err != nil {
		panic(err)
	}
	//res := string(response)
	//out, _ := json.Marshal(res)

	_ = json.NewEncoder(w).Encode(string(response))
}

func CreateDrawflowHandler() *DrawflowHandler {
	return &DrawflowHandler{}
}
