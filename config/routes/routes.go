package routes

import (
	"net/http"
	_ "strings"
	"github.com/gin-gonic/gin"
	_ "github.com/qor/qor"
	"qor-test/app/controllers"
	_ "qor-test/config"
	_ "qor-test/db"
)

func Router() *http.ServeMux {
	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	router.GET("/", controllers.Home)

	var mux = http.NewServeMux()
	mux.Handle("/", router)

	return mux
}
