package controllers

import (
	"blog-api/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var articles = []models.Article{
	{ID: 1, Title: "First Article", Content: "Content of the first article", CreatedAt: "2024-11-01", UpdatedAt: "2024-11-01"},
}

// GetArticles - Retourne la liste des articles
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

// GetArticle - Retourne un article spécifique par ID
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	for _, article := range articles {
		if string(article.ID) == id {
			c.JSON(http.StatusOK, article)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

// CreateArticle - Crée un nouvel article
func CreateArticle(c *gin.Context) {
	var newArticle models.Article

	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles = append(articles, newArticle)
	c.JSON(http.StatusCreated, newArticle)
}

// UpdateArticle - Met à jour un article par ID
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var updatedArticle models.Article

	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, article := range articles {
		if string(article.ID) == id {
			articles[i] = updatedArticle
			c.JSON(http.StatusOK, updatedArticle)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

// DeleteArticle - Supprime un article par ID
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	for i, article := range articles {
		if string(article.ID) == id {
			articles = append(articles[:i], articles[:i+1]...)
			c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}
