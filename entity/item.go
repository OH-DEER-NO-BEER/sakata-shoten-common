package entity

import "time"

type Item struct {
	Id    int
	Name  string
	Price int
	Time  time.Time
}

type Items []Item
