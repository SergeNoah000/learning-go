package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Modèle User avec tags GORM
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name" binding:"required,min=2,max=50"`
	Email string `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=1,max=150"`
	Posts []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}

// Modèle Post (One-to-Many avec User)
type Post struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"not null" json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=10"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

var db *gorm.DB

func main() {
	// Connexion MySQL
	// Format: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:password@tcp(localhost:3306)/afaapay?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erreur de connexion MySQL: " + err.Error())
	}

	// Auto-migration des modèles
	db.AutoMigrate(&User{}, &Post{})
	fmt.Println("✅ MySQL connecté et tables créées")

	// Routeur Gin
	r := gin.Default()
	r.Use(LoggerMiddleware())

	// Routes publiques v1
	v1 := r.Group("/v1")
	{
		// Users
		v1.GET("/users", getAllUsers)
		v1.GET("/users/:id", getUserByID)
		v1.POST("/users", createUser)
		v1.PUT("/users/:id", updateUser)
		v1.DELETE("/users/:id", deleteUser)

		// Posts
		v1.GET("/posts", getAllPosts)
		v1.GET("/posts/:id", getPostByID)
		v1.POST("/posts", createPost)
		v1.PUT("/posts/:id", updatePost)
		v1.DELETE("/posts/:id", deletePost)

		// Relations
		v1.GET("/users/:id/posts", getUserPosts)
	}

	// Info API
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API avec GORM et MySQL",
			"version": "4.0",
			"db":      "MySQL",
		})
	})

	fmt.Println("🚀 Serveur MySQL démarré sur http://localhost:8080")
	r.Run(":8080")
}

// === MIDDLEWARES ===

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[API] %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

// === USERS HANDLERS ===

func getAllUsers(c *gin.Context) {
	var users []User
	if err := db.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur BD"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users, "total": len(users)})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.Preload("Posts").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur BD"})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier si l'email existe
	var count int64
	db.Model(&User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email déjà utilisé"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur création"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utilisateur créé", "user": user})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur mise à jour"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur mis à jour", "user": user})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// Supprimer les posts d'abord (FK)
	db.Where("user_id = ?", id).Delete(&Post{})

	if err := db.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé"})
}

// === POSTS HANDLERS ===

func getAllPosts(c *gin.Context) {
	var posts []Post
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur BD"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts, "total": len(posts)})
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	var post Post

	if err := db.Preload("User").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur BD"})
		}
		return
	}
	c.JSON(http.StatusOK, post)
}

func createPost(c *gin.Context) {
	var post Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier si l'utilisateur existe
	var user User
	if err := db.First(&user, post.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur création"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post créé", "post": post})
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var post Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&Post{}).Where("id = ?", id).Updates(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur mise à jour"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post mis à jour", "post": post})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post supprimé"})
}

// === RELATIONS ===

func getUserPosts(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var user User
	if err := db.Preload("Posts").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur BD"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":        user.Name,
		"posts_count": len(user.Posts),
		"posts":       user.Posts,
	})
}
