package repository

import (
	"apinodos/src/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DrawflowMuttations struct {
}

func (d *DrawflowMuttations) SaveDrawflow(draw models.CreateDrawflow) *strings.Reader {

	jsonbody, err := json.Marshal(draw.Body.Drawflow.Home.Data)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day(),
	)

	data := fmt.Sprintf("{\"query\":\"mutation MyMutation {\\r\\n  addDrawflow(input: {body: \\\"%v\\\", createdAt: \\\"%v\\\", name: \\\"%v\\\"}) {\\r\\n    numUids\\r\\n  }\\r\\n}\",\"variables\":{}}",
		jsonbody,
		date,
		draw.Name,
	)
	fmt.Println(data)

	payload := strings.NewReader(data)

	return payload
}

func (d *DrawflowMuttations) GetDrawflows() *strings.Reader {

	payload := strings.NewReader("{\"query\":\"query MyQuery {\\r\\n  queryDrawflow {\\r\\n    body\\r\\n    createdAt\\r\\n    id\\r\\n    name\\r\\n  }\\r\\n}\\r\\n\",\"variables\":{}}")

	return payload
}

func CreateDrawflowMuttations() *DrawflowMuttations {
	return &DrawflowMuttations{}
}
