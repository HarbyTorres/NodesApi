package models

import "time"

type Drawflow struct {
	Id         uint64
	Body       map[string]interface{}
	CreateDate time.Time
}

type CreateDrawflow struct {
	Body map[string]interface{} `json:"body"`
}
