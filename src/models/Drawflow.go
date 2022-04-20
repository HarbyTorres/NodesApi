package models

import (
	"time"
)

type Drawflow struct {
	Id         uint64
	Body       DrawflowMap
	CreateDate time.Time
}

type CreateDrawflow struct {
	Body DrawflowMap `json:"body"`
}

type DrawflowMap struct {
	Drawflow struct {
		Home struct {
			Data map[string]interface{} `json:"data"`
		} `json:"Home"`
	} `json:"drawflow"`
}
