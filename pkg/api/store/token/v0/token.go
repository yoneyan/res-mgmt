package v0

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/res-mgmt/pkg/api/core/token"
	"log"
	"time"
)

type TokenStore struct {
	DB    *gorm.DB
	Token []token.Token
}

func (h *TokenStore) Create(t *token.Token) error {
	return h.DB.Create(&t).Error
}

func (h *TokenStore) Delete(t *token.Token) error {
	return h.DB.Delete(t).Error
}

func (h *TokenStore) Update(base uint, t *token.Token) error {
	var err error = nil

	if token.AddToken == base {
		err = h.DB.Model(&token.Token{Model: gorm.Model{ID: t.ID}}).Update(token.Token{Model: gorm.Model{},
			ExpiredAt: t.ExpiredAt, AccessToken: t.AccessToken}).Error
	} else if token.UpdateToken == base {
		err = h.DB.Model(&token.Token{Model: gorm.Model{ID: t.ID}}).Update("expired_at", t.ExpiredAt).Error
	} else if token.UpdateAll == base {
		err = h.DB.Model(&token.Token{Model: gorm.Model{ID: t.ID}}).Update(token.Token{
			ExpiredAt: t.ExpiredAt, UserToken: t.UserToken, TmpToken: t.TmpToken,
			AccessToken: t.AccessToken, Debug: t.Debug}).Error
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n %s", time.Now(), err)
	}
	return err
}

func (h *TokenStore) Get(base int, input *token.Token) error {
	var err error = nil

	if base == token.UserToken {
		err = h.DB.Where("user_token = ? AND expired_at > ?", input.UserToken, time.Now()).Find(&h.Token).Error
	} else if base == token.AccessToken {
		err = h.DB.Where("access_token = ? AND expired_at > ?", input.UserToken, time.Now()).Find(&h.Token).Error
	} else if base == token.UserTokenAndAccessToken {
		err = h.DB.Where("user_token = ? AND access_token = ? AND expired_at > ?",
			input.UserToken, input.AccessToken, time.Now()).Find(&h.Token).Error
	} else if base == token.ExpiredTime {
		err = h.DB.Where("expired_at < ? ", time.Now()).Find(&h.Token).Error
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n", time.Now())
	}
	return err

}

func (h *TokenStore) GetAll() error {
	err := h.DB.Find(&h.Token).Error
	if err != nil {
		return err
	}

	return nil
}
