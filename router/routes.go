package router

import (
	"github.com/bbcoder-gh/api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(routeMap map[string]*gin.HandlerFunc, apiRouter *gin.Engine, version int) {
	if apiRouter == nil {
		apiRouter = gin.Default()
	}

	userGroup := apiRouter.Group("/u")
	userGroup.GET("/", controllers.GetUsers)
	userGroup.GET("/:id", controllers.GetUser)
	userGroup.POST("/:id", controllers.UpdateUser)
	userGroup.PUT("/:id", controllers.AddUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)

}
