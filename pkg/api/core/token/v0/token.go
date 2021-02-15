package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/res-mgmt/pkg/api/core/common"
	"github.com/yoneyan/res-mgmt/pkg/api/core/token"
	"github.com/yoneyan/res-mgmt/pkg/api/core/tool/config"
	"github.com/yoneyan/res-mgmt/pkg/api/core/tool/hash"
	toolToken "github.com/yoneyan/res-mgmt/pkg/api/core/tool/token"
	dbToken "github.com/yoneyan/res-mgmt/pkg/api/store/token/v0"
	"log"
	"net/http"
	"strings"
	"time"
)

type TokenHandler struct {
	db *gorm.DB
}

func NewTokenHandler(database *gorm.DB) *TokenHandler {
	return &TokenHandler{db: database}
}

func (t *TokenHandler) GenerateInit(c *gin.Context) {
	userToken := c.Request.Header.Get("USER_TOKEN")
	log.Println("userToken: " + userToken)
	tmpToken, _ := toolToken.Generate(2)
	hDB := dbToken.TokenStore{DB: t.db}
	err := hDB.Create(&token.Token{ExpiredAt: time.Now().Add(1 * time.Hour), UserToken: userToken, TmpToken: tmpToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
	} else {
		log.Println("Time: " + time.Now().String() + " IP: " + c.ClientIP())
		c.JSON(http.StatusOK, &token.ResultTmpToken{Token: tmpToken})
	}
}

func (t *TokenHandler) Generate(c *gin.Context) {
	userToken := c.Request.Header.Get("USER_TOKEN")
	hashPass := c.Request.Header.Get("HASH_PASS")
	userName := c.Request.Header.Get("Name")
	hDB := dbToken.TokenStore{DB: t.db}
	err := hDB.Get(token.UserToken, &token.Token{UserToken: userToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}

	if config.Conf.Controller.UserName != userName {
		c.JSON(http.StatusUnauthorized, common.Error{Error: "invalid: username or password"})
		return
	}

	log.Println(hDB.Token)

	if hash.Generate(config.Conf.Controller.Password+hDB.Token[0].TmpToken) != strings.ToUpper(hashPass) {
		log.Println("pass(server): " + config.Conf.Controller.Password)
		log.Println("TmpToken: " + hDB.Token[0].TmpToken)
		log.Println("hash(server): " + hash.Generate(config.Conf.Controller.Password+hDB.Token[0].TmpToken))
		log.Println("hash(client): " + hashPass)
		c.JSON(http.StatusUnauthorized, common.Error{Error: "invalid: username or password"})
		return
	}
	accessToken, _ := toolToken.Generate(2)
	err = hDB.Update(token.AddToken, &token.Token{Model: gorm.Model{ID: hDB.Token[0].Model.ID},
		ExpiredAt: time.Now().Add(24 * time.Hour), AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
	} else {
		tmp := []token.Token{{AccessToken: accessToken}}
		c.JSON(http.StatusOK, &token.Result{Token: tmp})
	}
}

func (t *TokenHandler) Delete(c *gin.Context) {
	userToken := c.Request.Header.Get("USER_TOKEN")
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	hDB := dbToken.TokenStore{DB: t.db}
	err := hDB.Get(token.UserTokenAndAccessToken, &token.Token{UserToken: userToken, AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}
	if err = hDB.Delete(&token.Token{Model: gorm.Model{ID: hDB.Token[0].ID}}); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, token.Result{})
}
