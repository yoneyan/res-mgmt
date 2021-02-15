package item

import "github.com/jinzhu/gorm"

const (
	ID        = 0
	Name      = 1
	ManageID  = 2
	UpdateAll = 110
)

type Item struct {
	gorm.Model
	Name     string `json:"name"`
	ManageID string `json:"manage_id"`
	TypeID   uint   `json:"type_id"`
	OwnerID  uint   `json:"owner_id"`
	NOC      string `json:"noc"`
	Comment  string `json:"comment"`
}

type Result struct {
	Items []Item `json:"items"`
}
