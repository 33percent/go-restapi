package model

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tasks struct {
	gorm.Model
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
}
