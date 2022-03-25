package models

import "time"

type Drawflow struct {
	Id         uint64
	Body       string
	CreateDate time.Time
}

type CreateDrawflow struct {
	Body string `json:"body"`
}
