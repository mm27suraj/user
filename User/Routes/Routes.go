package Routes

import (
	"User/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() {
	r := gin.Default()
	grp1 := r.Group("/v1/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}

	r.Run(":8000")

}
