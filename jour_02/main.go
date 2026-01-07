package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Structure User pour stocker les utilisateurs
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age"`
}

// Base de données en mémoire
var users = []User{
	{ID: 1, Name: "Noah Mvondo", Email: "noah@example.com", Age: 25},
	{ID: 2, Name: "Alice Dupont", Email: "alice@example.com", Age: 30},
}

var nextID = 3

func main() {
	r := gin.Default()

	// Routes de base
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API CRUD d'utilisateurs",
			"version": "1.0",
		})
	})

	// GET - Récupérer tous les utilisateurs
	r.GET("/users", getUsers)

	// GET - Récupérer un utilisateur par ID
	r.GET("/users/:id", getUserByID)

	// POST - Créer un nouvel utilisateur
	r.POST("/users", createUser)

	// PUT - Mettre à jour un utilisateur
	r.PUT("/users/:id", updateUser)

	// DELETE - Supprimer un utilisateur
	r.DELETE("/users/:id", deleteUser)

	// Démarrer le serveur
	r.Run(":8080")
}

// GET /users - Récupérer tous les utilisateurs
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// GET /users/:id - Récupérer un utilisateur par ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Convertir le paramètre en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	for _, user := range users {
		if user.ID == idInt {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
}

// POST /users - Créer un nouvel utilisateur
func createUser(c *gin.Context) {
	var newUser User

	// Valider et lier le JSON
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assigner un nouvel ID
	newUser.ID = nextID
	nextID++

	// Ajouter à la liste
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

// PUT /users/:id - Mettre à jour un utilisateur
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User

	// Convertir le paramètre en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == idInt {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
}

// DELETE /users/:id - Supprimer un utilisateur
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// Convertir le paramètre en int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	for i, user := range users {
		if user.ID == idInt {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
}
