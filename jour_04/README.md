# Jour 4 - GORM et Base de Données

## Objectifs
- Installation et configuration de GORM
- Connexion à une base de données (MySQL/PostgreSQL/SQLite)
- Définition de modèles (structs) avec tags GORM
- Auto-migration des tables
- Opérations CRUD avec GORM
- Intégration GORM + Gin
- Gestion des relations (One-to-Many)

## Fichiers disponibles

- **main.go** - Version SQLite (par défaut)
- **main-mysql.go** - Version MySQL
- **main-postgres.go** - Version PostgreSQL

## Installation

```bash
go mod download

# Lancer SQLite (par défaut)
go run main.go

# Lancer MySQL
go run main-mysql.go

# Lancer PostgreSQL
go run main-postgres.go
```

## Configuration Base de Données

### SQLite (main.go)
Aucune configuration nécessaire, crée `afaapay.db` automatiquement.

### MySQL (main-mysql.go)
Modifier la ligne 36 avec vos identifiants :
```go
dsn := "root:password@tcp(localhost:3306)/afaapay?charset=utf8mb4&parseTime=True&loc=Local"
```

Créer la base :
```sql
CREATE DATABASE afaapay CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### PostgreSQL (main-postgres.go)
Modifier la ligne 36 avec vos identifiants :
```go
dsn := "host=localhost user=postgres password=postgres dbname=afaapay port=5432 sslmode=disable"
```

Créer la base :
```sql
CREATE DATABASE afaapay;
```

## Endpoints

### Users
- `GET /v1/users` - Liste tous les utilisateurs
- `GET /v1/users/:id` - Récupère un utilisateur
- `POST /v1/users` - Crée un utilisateur
- `PUT /v1/users/:id` - Met à jour un utilisateur
- `DELETE /v1/users/:id` - Supprime un utilisateur

### Posts
- `GET /v1/posts` - Liste tous les posts
- `GET /v1/posts/:id` - Récupère un post
- `POST /v1/posts` - Crée un post
- `PUT /v1/posts/:id` - Met à jour un post
- `DELETE /v1/posts/:id` - Supprime un post

### Relations
- `GET /v1/users/:id/posts` - Posts d'un utilisateur

## Tests

```bash
# Créer un utilisateur
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Noah","email":"noah@example.com","age":25}'

# Créer un post (user_id = 1)
curl -X POST http://localhost:8080/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"Mon premier post","content":"Ceci est le contenu du post","user_id":1}'

# Récupérer les posts d'un utilisateur
curl http://localhost:8080/v1/users/1/posts
```
