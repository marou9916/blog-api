package routes

import (
	"blog-api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Routes pour les articles
	router.GET("/articles", controllers.GetArticles)
	router.GET("/articles/:id", controllers.GetArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.PUT("/articles/:id", controllers.UpdateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)

	return router
}
