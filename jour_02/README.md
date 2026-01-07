# Jour 02 - Routes REST et API CRUD

## Ce qui a été appris

### 1. Méthodes HTTP avec Gin
- **GET**: Récupérer des données
- **POST**: Créer de nouvelles données
- **PUT**: Mettre à jour des données existantes
- **DELETE**: Supprimer des données

### 2. Paramètres d'URL
- **Paramètres de route**: `/users/:id` - Récupérer avec `c.Param("id")`
- **Query parameters**: `/users?age=25` - Récupérer avec `c.Query("age")`

### 3. Manipulation de JSON
- **c.ShouldBindJSON()**: Lier et valider le JSON reçu
- **binding tags**: `binding:"required,email"` pour la validation
- **c.JSON()**: Envoyer une réponse JSON

### 4. Codes de statut HTTP
- `200 OK`: Succès
- `201 Created`: Ressource créée
- `400 Bad Request`: Erreur de validation
- `404 Not Found`: Ressource introuvable

## API Endpoints

### GET /users
Récupère tous les utilisateurs
```bash
curl http://localhost:8080/users
```

### GET /users/:id
Récupère un utilisateur spécifique
```bash
curl http://localhost:8080/users/1
```

### POST /users
Crée un nouvel utilisateur
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":28}'
```

### PUT /users/:id
Met à jour un utilisateur
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Noah Updated","email":"noah.new@example.com","age":26}'
```

### DELETE /users/:id
Supprime un utilisateur
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Comment exécuter

```bash
cd jour_02
go run main.go
```

Le serveur démarre sur http://localhost:8080

## Concepts clés
- **Structs**: Définir des structures de données
- **Tags JSON**: `json:"name"` pour la sérialisation
- **Binding tags**: Validation automatique des données
- **CRUD**: Create, Read, Update, Delete
- **Base de données en mémoire**: Utilisation d'un slice pour stocker les données
