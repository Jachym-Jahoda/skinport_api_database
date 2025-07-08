package database

import (
	"gorm.io/gorm"
)

type CsgoItem struct {
	gorm.Model
	Volume90Day  int
	Volume30Day  int
	Name         string
	MarketPage   string
	Average90Day float64
	Average30Day float64
	MinPrice     float64
	MarketPlace  string
	Note         string
}
