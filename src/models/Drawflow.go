package models

import (
	"time"
)

type Drawflow struct {
	Id         uint64
	Name       string
	Body       DrawflowMap
	CreateDate time.Time
}

type CreateDrawflow struct {
	Name  string      `json:"name"`
	Body  DrawflowMap `json:"body"`
	DType []string    `json:"dgraph.type,omitempty"`
}

type DrawflowMap struct {
	Drawflow struct {
		Home struct {
			Data map[string]interface{} `json:"data"`
		} `json:"Home"`
	} `json:"drawflow"`
}
