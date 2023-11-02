package database

import (
	"Draft_API_GO/myapp/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

var database *gorm.DB
var errDB error

func Init() {

	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlconnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=10",
		dbhost,
		dbport,
		dbuser,
		dbpassword,
		dbname,
	)

	fmt.Println(psqlconnection)


	// open database
	database, errDB = gorm.Open(postgres.Open(psqlconnection), &gorm.Config{})

	if errDB != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", errDB)
	}
	
	fmt.Println("Connexion à la base de données PostgreSQL établie !")

	fmt.Println(database.Find(&models.Article{}))

	if !database.Migrator().HasTable(&models.Article{}) {
		if err := database.Create(&models.Article{}) ; err != nil {
			log.Println("Impossible de créer la table:", err)
		}
	} else {
		if err := database.Migrator().AutoMigrate(&models.Article{}) ; err != nil {
			log.Println("Impossible de mettre à jour la table:", err)
		}
	}
}	

func GetDataBase() *gorm.DB {
	return database
}

func CloseDataBase() {
    db, err := database.DB()
    if err != nil {
        log.Println("Erreur lors de la fermeture de la base de données :", err)
        return
    }
    db.Close()
}

