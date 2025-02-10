package models

import "gorm.io/gorm"

type ResponseIP struct {
	IP          string `json:"ip"`
	CountryName string `json:"country_name"`
}

type IP struct {
	gorm.Model
	IP          string `json:"ip" gorm:"ip"`
	CountryName string `json:"country_name" gorm:"country_name"`
}
