package v0

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/res-mgmt/pkg/api/core/token"
	dbToken "github.com/yoneyan/res-mgmt/pkg/api/store/token/v0"
)

func Authentication(db *gorm.DB, data token.Token) error {
	hDB := dbToken.TokenStore{DB: db}

	err := hDB.Get(token.AccessToken, &data)
	if len(hDB.Token) == 0 {
		return fmt.Errorf("auth failed")
	}
	if err != nil {
		return fmt.Errorf("db error: " + err.Error())
	}

	return nil
}
