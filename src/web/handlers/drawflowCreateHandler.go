package handlers

import (
	"apinodos/src/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type SaveDrawflow struct {
}

func (*SaveDrawflow) SaveDrawflow(w http.ResponseWriter, r *http.Request) {
	//var node models.Node
	var draw models.DrawflowMap

	encoder := json.NewDecoder(r.Body)
	encoder.Decode(&draw)

	var nodes []models.Node
	for i := 1; i <= len(draw.Drawflow.Home.Data); i++ {

		var node models.Node
		jsonbody, err := json.Marshal(draw.Drawflow.Home.Data[strconv.Itoa(i)])
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(jsonbody, &node); err != nil {
			panic(err)
		}

		//nodeDe, _ := draw.Drawflow.Home.Data[strconv.Itoa(i)].(models.Node)
		nodes = append(nodes, node)
	}

	_ = json.NewEncoder(w).Encode(nodes)
}

func CreateSaveDrawflow() *SaveDrawflow {
	return &SaveDrawflow{}
}
