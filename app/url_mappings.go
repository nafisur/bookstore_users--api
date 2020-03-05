package app

import (
	"github.com/nafisur/bookstore_users-api/controllers"
)
func MapURLs() {

	router.GET("/users/:user_id", controllers.GetUser)
	//router.GET("/users/:search", controllers.FindUser)

	router.POST("/createUser", controllers.CreateUser)

}
