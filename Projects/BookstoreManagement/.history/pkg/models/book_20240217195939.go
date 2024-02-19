package models

import (
	"github.com/jinzhu/gorm"
	"../../pkg/config/"
)

var db "gorm.DB"

type Book struct {
	gorm.model 
	Name string `gorm:json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}