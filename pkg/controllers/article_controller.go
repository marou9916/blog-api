package controllers

import (
	"blog-api/pkg/db"
	"blog-api/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticles - Retourne la liste des articles
func GetArticles(c *gin.Context) {
	var articles models.Article

	if err := db.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les articles"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

// GetArticle - Retourne un article spécifique par ID
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	var article models.Article

	if err := db.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// CreateArticle - Crée un nouvel article
func CreateArticle(c *gin.Context) {
	var newArticle models.Article

	// Lire le JSON et vérifier les erreurs
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ajouter à la liste des articles dans la base de données
	if err := db.DB.Create(&newArticle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer l'article"})
		return
	}

	// Retourner la réponse
	c.JSON(http.StatusCreated, newArticle)
}

// UpdateArticle - Met à jour un article par ID
func UpdateArticle(c *gin.Context) {
	var updatedArticle models.Article
	id := c.Param("id")

	//Convertir l'id en entier
	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid article ID"})
		return
	}

	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var article models.Article

	//Récupérer l'article à mettre à jour
	if err := db.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	//Mettre à jour l'article
	article.Title = updatedArticle.Title
	article.Content = updatedArticle.Content
	article.CreatedAt = updatedArticle.CreatedAt
	article.UpdatedAt = updatedArticle.UpdatedAt

	//Sauvegarder les changements dans la base de données
	if err := db.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour l'article"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// DeleteArticle - Supprime un article par ID
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var article models.Article

	if err := db.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "impossible de récupérer l'article à supprimer"})
		return
	}

	if err := db.DB.Delete(&article, articleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de supprimer l'article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "article found and deleted"})
}
