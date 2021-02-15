package v0

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/res-mgmt/pkg/api/core/item"
	"log"
	"time"
)

type ItemStore struct {
	DB    *gorm.DB
	Items []item.Item
}

func (h *ItemStore) Create(t *item.Item) error {
	return h.DB.Create(&t).Error
}

func (h *ItemStore) Delete(t *item.Item) error {
	return h.DB.Delete(t).Error
}

func (h *ItemStore) Update(base uint, input *item.Item) error {
	var err error = nil

	if item.UpdateAll == base {
		err = h.DB.Model(&item.Item{Model: gorm.Model{ID: input.ID}}).Update(item.Item{
			Model:   gorm.Model{ID: input.ID},
			Name:    input.Name,
			TypeID:  input.TypeID,
			OwnerID: input.OwnerID,
			NOC:     input.NOC,
			Comment: input.Comment,
		}).Error
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n %s", time.Now(), err)
	}
	return err
}

func (h *ItemStore) Get(base uint, input *item.Item) error {
	var err error = nil

	if base == item.ID { //ID
		err = h.DB.First(&h.Items, input.ID).Error
	} else if base == item.Name { //Name
		err = h.DB.Where("name = ?", item.Name).Find(&h.Items).Error
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n", time.Now())
	}

	return err
}

func (h *ItemStore) GetAll() error {
	err := h.DB.Find(&h.Items).Error
	if err != nil {
		return err
	}

	return nil
}
