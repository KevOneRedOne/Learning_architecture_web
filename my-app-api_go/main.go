package main

import (
	"Draft_API_GO/myapp/controllers"
	"Draft_API_GO/myapp/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Démarrage du serveur...")

	log.Println("Chargement des variables d'environnements...")
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", errEnv)
	}

	// log.Println("Connexion à Redis...")
	// redis.Init()

	log.Println("Connexion à la base de données...")
	database.Init()

	router := gin.Default()
	apiPath := router.Group("/api/v1")

	{
		articles := apiPath.Group("/articles")
		{
			articles.GET("/", controllers.GetArticles)
			articles.GET("/:id", controllers.GetArticleByID)
			articles.POST("/create", controllers.CreateArticle)
			articles.PUT("/:id", controllers.UpdateArticle)
			articles.DELETE("delete/:id", controllers.DeleteArticle)
		}
	}
	router.Use(cors.Default())
	router.Run()

	defer database.CloseDataBase()
}
