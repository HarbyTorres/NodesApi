package repository

import (
	"strings"
)

type DrawflowMuttations struct {
}

func CreateDrawflowMuttations() *DrawflowMuttations {
	return &DrawflowMuttations{}
}

func (d *DrawflowMuttations) SaveDrawflow() *strings.Reader {
	payload := strings.NewReader("{\"query\":\"mutation MyMutation {\\r\\n  addDrawflow(input: {body: \\\"drawflowtest3\\\", crateDate: \\\"2022-04-22\\\", name: \\\"drawformback2\\\"}) {\\r\\n    numUids\\r\\n  }\\r\\n}\",\"variables\":{}}")

	return payload
}
