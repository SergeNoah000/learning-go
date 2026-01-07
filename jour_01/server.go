package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Créer une instance de Gin
	r := gin.Default()

	// Route GET simple
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenue sur mon premier serveur Gin!",
		})
	})

	// Route GET avec un message
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World avec Gin Framework!",
		})
	})

	// Démarrer le serveur sur le port 8080
	r.Run(":8080")
}
