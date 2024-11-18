package main

import "blog-api/pkg/routes"

func main() {
	//Configuration des routes
	router := routes.SetupRouter()

	//Démarrer le serveur
	router.Run(":8080") //sur le port 8080

}
