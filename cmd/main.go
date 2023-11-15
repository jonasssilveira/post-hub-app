package main

import (
	"PostHubApp/adapter/repository/database"
	"PostHubApp/cmd/api"
	handlers "PostHubApp/posthubapi/handlers/post_handler"
	"PostHubApp/posthubapi/handlers/web"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	connect := database.Connect()
	manager := database.NewEntityManager(connect)
	getPost := handlers.NewApiGet(manager)
	savePost := handlers.NewApiSave(manager)

	router.Group("/")
	{
		post := router.Group("posthub/api/v1/post")
		{
			post.GET("/:id", web.HandleFunc(getPost.Handler))
			post.POST("/", web.HandleFunc(savePost.Handler))
			//books.GET("/", controllers.GetAll())
			//books.("/", controllers.UpdateBook())
			//books.DELETE("/:id", controllers.DeleteBook())
		}
	}
	server := api.NewServer(router)
	server.Run()
}
