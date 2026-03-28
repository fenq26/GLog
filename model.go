package main

import "time"

type Log struct {
	Service  string    `json:"service"`
	Level    string    `json:"level"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Owner    string    `json:"owner"`
}