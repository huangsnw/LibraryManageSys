package models

import "time"

type Result struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    time.Time `json:"data"`
}
