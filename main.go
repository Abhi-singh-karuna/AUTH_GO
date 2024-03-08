package main

import (
	"time"

	"github.com/Abhi-singh-karuna/AUTH_GO/api/route"
	"github.com/Abhi-singh-karuna/AUTH_GO/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
