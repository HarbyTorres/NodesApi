package services

import (
	"apinodos/src/models"
	"encoding/json"
	"fmt"
	"strconv"
)

func CodeGenerator(draw models.DrawflowMap) string {
	var nodes []models.Node
	nodes = processDrawflow(draw)

	for i := 0; i < len(nodes); i++ {
		switch nodes[i].Class {
		case "number":
			fmt.Println("number")
		case "add":
			fmt.Println("add")
		}
	}
	return "Ok"
}

func processDrawflow(draw models.DrawflowMap) []models.Node {
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
	return nodes
}
