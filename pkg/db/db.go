package db

import (
	"blog-api/pkg/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB est une instance globale de la connexion à la base
var DB *gorm.DB

// InitDatabase initialise la connexion à SQLite
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Échec de la connexion à la base de données:", err)
	}
	log.Println("Base de données connectée avec succès !")

	// Migration automatique des modèles
	err = DB.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Échec de la migration de la base de données:", err)
	}
	log.Println("Migration de la base de données effectuée avec succès !")

}
