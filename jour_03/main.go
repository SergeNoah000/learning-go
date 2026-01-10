package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Structure User avec validation
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required,min=2,max=50"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=1,max=150"`
}

// Base de données en mémoire
var users = []User{
	{ID: 1, Name: "Noah Mvondo", Email: "noah@example.com", Age: 25},
	{ID: 2, Name: "Alice Dupont", Email: "alice@example.com", Age: 30},
	{ID: 3, Name: "Bob Martin", Email: "bob@example.com", Age: 28},
}

var nextID = 4

// Middleware personnalisé pour logger les requêtes
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Heure de début
		startTime := time.Now()

		// Traiter la requête
		c.Next()

		// Calculer la durée
		duration := time.Since(startTime)

		// Logger
		fmt.Printf("[LOGGER] %s %s - Durée: %v - Status: %d\n",
			c.Request.Method,
			c.Request.URL.Path,
			duration,
			c.Writer.Status(),
		)
	}
}

// Middleware d'authentification basique
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer le token du header Authorization
		token := c.GetHeader("Authorization")

		// Vérifier le format "Bearer <token>"
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token d'authentification requis",
			})
			c.Abort()
			return
		}

		// Extraire le token
		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format de token invalide. Utilisez: Bearer <token>",
			})
			c.Abort()
			return
		}

		// Vérifier le token (ici token simple pour démo)
		if parts[1] != "secret-token-123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalide",
			})
			c.Abort()
			return
		}

		// Token valide, continuer
		c.Next()
	}
}

func main() {
	// Créer le routeur avec les middlewares par défaut (Logger et Recovery)
	r := gin.Default()

	// Ajouter notre middleware personnalisé à toutes les routes
	r.Use(LoggerMiddleware())

	// Route d'accueil
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API avec Middlewares et Groupes de Routes",
			"version": "3.0",
			"endpoints": gin.H{
				"v1":    "/v1/* (public)",
				"v2":    "/v2/* (nécessite authentification)",
				"admin": "/admin/* (nécessite authentification)",
			},
		})
	})

	// === GROUPE V1 - Routes publiques ===
	v1 := r.Group("/v1")
	{
		// Routes CRUD pour les utilisateurs
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUserByID)
		v1.POST("/users", createUser)
		v1.PUT("/users/:id", updateUser)
		v1.DELETE("/users/:id", deleteUser)
	}

	// === GROUPE V2 - Routes avec authentification ===
	v2 := r.Group("/v2")
	v2.Use(AuthMiddleware()) // Appliquer le middleware d'auth à tout le groupe
	{
		v2.GET("/users", getUsers)
		v2.POST("/users", createUser)
		v2.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Profil utilisateur authentifié",
				"user":    "Noah Mvondo",
			})
		})
	}

	// === GROUPE ADMIN - Routes administrateur ===
	admin := r.Group("/admin")
	admin.Use(AuthMiddleware()) // Routes protégées
	{
		admin.GET("/stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"total_users":      len(users),
				"server_uptime":    "2h30m",
				"requests_handled": 1523,
			})
		})

		admin.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"users":       users,
				"total":       len(users),
				"admin_view":  true,
			})
		})
	}

	// Démarrer le serveur
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080")
	fmt.Println("📖 Routes disponibles:")
	fmt.Println("   - GET    /v1/users           (public)")
	fmt.Println("   - POST   /v1/users           (public)")
	fmt.Println("   - GET    /v2/users           (auth requise)")
	fmt.Println("   - GET    /admin/stats        (auth requise)")
	fmt.Println("\n🔐 Token pour test: Bearer secret-token-123")
	
	r.Run(":8080")
}

// GET /users - Récupérer tous les utilisateurs
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": len(users),
	})
}

// GET /users/:id - Récupérer un utilisateur par ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalide - doit être un nombre entier",
		})
		return
	}

	for _, user := range users {
		if user.ID == idInt {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("Utilisateur avec ID %d non trouvé", idInt),
	})
}

// POST /users - Créer un nouvel utilisateur avec validation
func createUser(c *gin.Context) {
	var newUser User

	// Valider et lier le JSON avec les tags binding
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erreur de validation",
			"details": err.Error(),
			"help": gin.H{
				"name":  "requis, min 2 caractères, max 50",
				"email": "requis, format email valide",
				"age":   "requis, entre 1 et 150",
			},
		})
		return
	}

	// Vérifier si l'email existe déjà
	for _, user := range users {
		if user.Email == newUser.Email {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Un utilisateur avec cet email existe déjà",
			})
			return
		}
	}

	// Assigner un nouvel ID
	newUser.ID = nextID
	nextID++

	// Ajouter à la liste
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Utilisateur créé avec succès",
		"user":    newUser,
	})
}

// PUT /users/:id - Mettre à jour un utilisateur
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalide",
		})
		return
	}

	// Valider les données
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erreur de validation",
			"details": err.Error(),
		})
		return
	}

	// Mettre à jour l'utilisateur
	for i, user := range users {
		if user.ID == idInt {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"message": "Utilisateur mis à jour avec succès",
				"user":    updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Utilisateur non trouvé",
	})
}

// DELETE /users/:id - Supprimer un utilisateur
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalide",
		})
		return
	}

	for i, user := range users {
		if user.ID == idInt {
			// Supprimer l'utilisateur
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Utilisateur %s supprimé avec succès", user.Name),
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Utilisateur non trouvé",
	})
}
