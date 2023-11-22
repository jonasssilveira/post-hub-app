package main

import (
	"PostHubApp/adapter/repository/database"
	"PostHubApp/cmd/api"
	handlerscomment "PostHubApp/posthubapi/handlers/comment"
	handlerspost "PostHubApp/posthubapi/handlers/post"
	"PostHubApp/posthubapi/handlers/web"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	connect := database.Connect()
	manager := database.NewEntityManager(connect)

	getPost := handlerspost.NewApiGetPost(manager)
	getAllPost := handlerspost.NewApiGetAllPost(manager)
	savePost := handlerspost.NewApiSavePost(manager)

	getPostComment := handlerscomment.NewApiGetComment(manager)
	getAllPostComment := handlerscomment.NewApiGetAllComment(manager)
	savePostComment := handlerscomment.NewApiSaveComment(manager)

	router.Group("/")
	{
		post := router.Group("posthub/api/v1/post")
		{
			post.GET("/:id", web.HandleFunc(getPost.Handler))
			post.GET("/", web.HandleFunc(getAllPost.Handler))
			post.POST("/", web.HandleFunc(savePost.Handler))
		}
		comment := router.Group("posthub/api/v1/comment")
		{
			comment.GET("/:id", web.HandleFunc(getPostComment.Handler))
			comment.GET("/", web.HandleFunc(getAllPostComment.Handler))
			comment.POST("/", web.HandleFunc(savePostComment.Handler))
		}
	}
	server := api.NewServer(router)
	server.Run()
}
