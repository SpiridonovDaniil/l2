package models

import "time"

type Event struct {
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	UserId int       `json:"user_id"`
}

type JsonResult struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}
