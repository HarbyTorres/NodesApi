package models

import "time"

type Drawflow struct {
	Id         uint64
	body       string
	CreateDate time.Time
}
