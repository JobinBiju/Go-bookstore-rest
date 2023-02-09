package models

import (
	"time"
)

type Book struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"desc"`
	Published   time.Time `json:"pub"`
}