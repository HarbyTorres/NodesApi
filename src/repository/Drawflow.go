package repository

import (
	"apinodos/src/models"
	"strings"
)

type DrawflowMuttations struct {
}

func (d *DrawflowMuttations) SaveDrawflow(draw models.CreateDrawflow) *strings.Reader {

	payload := strings.NewReader("{\"query\":\"mutation MyMutation {\\r\\n  addDrawflow(input: {body: \\\"drawflowtest3\\\", crateDate: \\\"2022-04-22\\\", name: \\\"drawformback2\\\"}) {\\r\\n    numUids\\r\\n  }\\r\\n}\",\"variables\":{}}")

	return payload
}

func (d *DrawflowMuttations) GetDrawflows() *strings.Reader {

	payload := strings.NewReader("{\"query\":\"query MyQuery {\\r\\n  queryDrawflow {\\r\\n    body\\r\\n    crateDate\\r\\n    id\\r\\n    name\\r\\n  }\\r\\n}\\r\\n\",\"variables\":{}}")

	return payload
}

func CreateDrawflowMuttations() *DrawflowMuttations {
	return &DrawflowMuttations{}
}
