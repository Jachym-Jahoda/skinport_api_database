package database

import (
	"gorm.io/gorm"
)

type CsgoItem struct {
	gorm.Model
	Volume90Day  int    `gorm:"uniqueIndex:volume_90_day"`
	Volume30Day  int    `gorm:"uniqueIndex:volume_30_day"`
	Name         string `gorm:"uniqueIndex:name"`
	MarketPage   string
	Average90Day float64 `gorm:"uniqueIndex:average_90_day"`
	Average30Day float64 `gorm:"uniqueIndex:average_30_day"`
	MinPrice     float64 `gorm:"uniqueIndex:min_price"`
	Note         string
}
