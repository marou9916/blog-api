package main

import (
	"blog-api/pkg/db"
	"blog-api/pkg/routes"
	"log"
)

func main() {
	// Initialiser la base de données
	db.InitDatabase()

	// Configurer les routes
	router := routes.SetupRouter()

	log.Println("Le serveur fonctionne sur http://localhost:8080")
	//Démarrer le serveur
	router.Run(":8080")
}
