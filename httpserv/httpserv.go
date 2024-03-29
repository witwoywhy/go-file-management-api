package httpserv

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	app := gin.Default()
	app.Use(gin.Recovery())

	bindUploadRoute(app)
	bindListRoute(app)
	bindRemoveRoute(app)

	app.Run(fmt.Sprintf(":%v", viper.Get("app.port")))
}
