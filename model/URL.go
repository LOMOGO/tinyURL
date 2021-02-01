package model

import "gorm.io/gorm"

type URLInfo struct {
	gorm.Model
	URL     string `json:"url" binding:"url"`
	URLCode string
}
