package routes

import (
	routes "nivasBackendMain/App/routes/Authentation"

	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.Engine) {
	route := router.Group("api/v1/app")
	routes.SignUpRoutes(route)

}
