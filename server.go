package main

import (
	"golang-mux-api/cache"
	"golang-mux-api/controller"
	router "golang-mux-api/http"
 
)

var (
	userCache      cache.UserCache           = cache.NewRedisCache("localhost:6379", 0, 180)
	userController controller.UserController = controller.NewUserController(userCache)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/users", userController.GetUsers)
	httpRouter.GET("/users/{id}", userController.GetUserByID)
	httpRouter.POST("/users", userController.AddUser)
	httpRouter.DELETE("/users/{id}", userController.DeleteUser)

	httpRouter.SERVE("localhost:8000")
}
