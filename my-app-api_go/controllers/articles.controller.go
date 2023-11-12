package controllers

import (
	"Draft_API_GO/myapp/database"
	"Draft_API_GO/myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArticles ... Get all articles
// CRUD: Read
func GetArticles(connection *gin.Context) {
	var article []models.Article
	db := database.GetDataBase()
	
	if err := db.Order("id ASC").Find(&article).Error; err != nil  {
		connection.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	connection.JSON(http.StatusOK, article)
}

// GetArticleByID ... Get the article by id
// CRUD: Read
func GetArticleByID(connection *gin.Context) {
	var article []models.Article
	db := database.GetDataBase()

	id := connection.Param("id")

	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		connection.AbortWithStatus(http.StatusNotFound)
		return
	}
	connection.JSON(http.StatusOK, article)
}


// // CreateArticle ... Create article
// // CRUD: Create
func CreateArticle(connection *gin.Context) {
	var article []models.Article
	db := database.GetDataBase()

	if err := connection.BindJSON(&article); err != nil {
		connection.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Create(&article)
	connection.JSON(http.StatusOK, &article)
}

// // updateArticle ... Update article
// // CRUD: Update
func UpdateArticle(connection *gin.Context) {
	var article []models.Article
	db := database.GetDataBase()

	id := connection.Param("id")

	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		connection.AbortWithStatus(http.StatusNotFound)
		return
	}
	connection.BindJSON(&article)
	db.Save(&article)
	connection.JSON(http.StatusOK, &article)
}

// // DeleteArticle ... Delete article
// // CRUD: Delete
func DeleteArticle(connection *gin.Context) {
	var article []models.Article
	db := database.GetDataBase()

	id := connection.Param("id")

	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		connection.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&article)
	connection.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}