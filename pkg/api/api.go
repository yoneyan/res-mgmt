package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	item "github.com/yoneyan/res-mgmt/pkg/api/core/item/v0"
	token "github.com/yoneyan/res-mgmt/pkg/api/core/token/v0"
	"github.com/yoneyan/res-mgmt/pkg/api/core/tool/config"
	"github.com/yoneyan/res-mgmt/pkg/api/store"
	"log"
	"net/http"
	"strconv"
	"time"
)

func RestAPI() {
	router := gin.Default()
	router.Use(cors)

	// DBの呼び出し
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}

	defer db.Close()

	hItem := item.NewItemHandler(db)
	hToken := token.NewTokenHandler(db)

	go token.TokenRemove()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//
			// Itemuser_id
			//
			// Add
			v1.POST("/item", hItem.Add)
			// Delete
			v1.DELETE("/item/:id", hItem.Delete)
			// Update
			v1.PUT("/item/:id", hItem.Update)
			// Get
			v1.GET("/item", hItem.GetAll)
			v1.GET("/item/:id", hItem.Get)
			//
			// Token
			//
			// Generate
			v1.POST("/token/generate", hToken.Generate)
			// Add
			v1.GET("/token/init", hToken.GenerateInit)
			// get token for user
			v1.GET("/token", hToken.Generate)
			// delete
			v1.DELETE("/token", hToken.Delete)
		}
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Conf.Controller.Port), router))
}

func cors(c *gin.Context) {

	//c.Header("Access-Control-Allow-Headers", "Accept, Content-ID, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-ID", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
