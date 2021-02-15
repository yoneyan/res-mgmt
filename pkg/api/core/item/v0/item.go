package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	auth "github.com/yoneyan/res-mgmt/pkg/api/core/auth/v0"
	"github.com/yoneyan/res-mgmt/pkg/api/core/common"
	"github.com/yoneyan/res-mgmt/pkg/api/core/item"
	"github.com/yoneyan/res-mgmt/pkg/api/core/token"
	dbItem "github.com/yoneyan/res-mgmt/pkg/api/store/item/v0"
	"log"
	"net/http"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(database *gorm.DB) *ItemHandler {
	return &ItemHandler{db: database}
}

func (t *ItemHandler) Add(c *gin.Context) {
	var input item.Item
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	err := c.BindJSON(&input)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, common.Error{Error: err.Error()})
		return
	}

	if err = check(input); err != nil {
		c.JSON(http.StatusBadRequest, common.Error{Error: err.Error()})
		return
	}

	err = auth.Authentication(t.db, token.Token{AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Error{Error: err.Error()})
		return
	}

	hDB := dbItem.ItemStore{DB: t.db}

	if err = hDB.Create(&item.Item{
		Name: input.Name, TypeID: input.TypeID, OwnerID: input.OwnerID, NOC: input.NOC, Comment: input.Comment,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, item.Item{})
	}
}

func (t *ItemHandler) Delete(c *gin.Context) {
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	err := auth.Authentication(t.db, token.Token{AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Error{Error: err.Error()})
		return
	}

	hDB := dbItem.ItemStore{DB: t.db}

	if err = hDB.Delete(&item.Item{ManageID: c.Param("id")}); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, item.Item{})
	}
}

func (t *ItemHandler) Update(c *gin.Context) {
	var input item.Item
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	err := auth.Authentication(t.db, token.Token{AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Error{Error: err.Error()})
		return
	}

	hDB := dbItem.ItemStore{DB: t.db}
	err = hDB.Get(item.ManageID, &item.Item{ManageID: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}

	replaceItem := hDB.Items[0]
	err = replace(input, replaceItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Error{Error: err.Error()})
		return
	}

	err = hDB.Update(item.UpdateAll, &replaceItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, item.Item{})
	}
}

func (t *ItemHandler) Get(c *gin.Context) {
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	err := auth.Authentication(t.db, token.Token{AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Error{Error: err.Error()})
		return
	}

	hDB := dbItem.ItemStore{DB: t.db}
	err = hDB.Get(item.ManageID, &item.Item{ManageID: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, hDB.Items)
}

func (t *ItemHandler) GetAll(c *gin.Context) {
	accessToken := c.Request.Header.Get("ACCESS_TOKEN")

	err := auth.Authentication(t.db, token.Token{AccessToken: accessToken})
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Error{Error: err.Error()})
		return
	}

	hDB := dbItem.ItemStore{DB: t.db}
	err = hDB.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error{Error: err.Error()})
		return
	}

	log.Println(hDB.Items)

	c.JSON(http.StatusOK, item.Result{Items: hDB.Items})
}
