package token

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	ID                      = 0
	UserToken               = 10
	UserTokenAndAccessToken = 11
	AccessToken             = 12
	ExpiredTime             = 13
	AdminToken              = 20
	AddToken                = 100
	UpdateToken             = 101
	UpdateAll               = 110
)

type Token struct {
	gorm.Model
	ExpiredAt   time.Time `json:"expired_at"`
	UserToken   string    `json:"user_token"`
	TmpToken    string    `json:"tmp_token"`
	AccessToken string    `json:"access_token"`
	Debug       string    `json:"debug"`
}

type Result struct {
	Token []Token `json:"token"`
}

type ResultTmpToken struct {
	Token string `json:"token"`
}

type ResultDatabase struct {
	Err   error
	Token []Token
}
