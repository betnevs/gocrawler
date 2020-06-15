package models

import (
	"time"
)

type House struct {
	Title    string
	DateTime string
	Images   []string
	Url      string
}

type HouseItem struct {
	ID          int64     `gorm:"column:ID;primary_key"`
	Title       string    `gorm:"column:title"`
	PublishTime time.Time `gorm:"column:publish_time"`
	Images      string    `gorm:"column:images"`
	Url         string    `gorm:"column:url"`
}
