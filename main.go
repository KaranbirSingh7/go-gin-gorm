package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karanbirsingh7/go-gin-gorm/internal/configuration"
	"github.com/karanbirsingh7/go-gin-gorm/internal/db"
)

// pointer to original database connection on startup
var DB *db.Database

func main() {
	// connect to DB
	conn, err := db.New()
	if err != nil {
		panic(err.Error())
	}
	DB = &conn

	if err := DB.MigrateSchemas(&configuration.RecoveryScenario{}); err != nil {
		panic(err.Error())
	}

	// setup router
	r := setupRouter()

	// listen on port
	if err := r.Run(":8080"); err != nil {
		panic(err.Error())
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// ping handler
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.GET("/dr/config", func(ctx *gin.Context) {
		configs, err := DB.GetRecoveryScenarios(ctx)
		if err != nil {
			ctx.Error(err)
		}
		ctx.JSON(http.StatusOK, configs)
	})

	r.POST("/dr/config", func(ctx *gin.Context) {
		var rs configuration.RecoveryScenario

		// read body bytes
		jsonData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Error(err)
		}
		defer ctx.Request.Body.Close()

		// unmarshal into struct
		err = json.Unmarshal(jsonData, &rs)
		if err != nil {
			ctx.Error(err)
		}

		insertedObj, err := DB.InsertRecoveryScenario(ctx, rs)
		if err != nil {
			ctx.Error(err)
		}

		ctx.JSON(http.StatusCreated, insertedObj)
	})
	return r
}
