package item

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name    string `json:"name"`
	Type    uint   `json:"type"`
	Comment string `json:"comment"`
}
