package services

import (
	"apinodos/src/models"
	"encoding/json"
	"fmt"
	"strconv"
)

type CodeGenerator2 struct {
}

func CodeGenerator(draw models.DrawflowMap) string {
	var nodes []models.Node
	nodes = processDrawflowMap(draw)

	res := "def func():\n"

	for i := 0; i < len(nodes); i++ {
		switch nodes[i].Class {
		case "number":
			res += numCodeGenerator(nodes[i])
			break
		case "if":
			fmt.Println("if")
			break
		default:
			res += oppCodeGenerator(nodes[i])
		}
	}
	fmt.Println(res)
	return res
}

func oppCodeGenerator(node models.Node) string {
	num1 := node.Inputs.Input1.Connections[0].Node
	num2 := node.Inputs.Input2.Connections[0].Node
	var opp string
	switch node.Class {
	case "add":
		opp = "+"
	case "subt":
		opp = "-"
	case "divide":
		opp = "/"
	case "multiply":
		opp = "*"
	}
	return fmt.Sprintf("\tnum%v = num%v %s num%v\n", node.ID, num1, opp, num2)
}

func numCodeGenerator(node models.Node) string {
	return fmt.Sprintf("\tnum%v = %v\n", node.ID, node.Data["num"])
}

func processDrawflowMap(draw models.DrawflowMap) []models.Node {
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
