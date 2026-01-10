# Exemples de requêtes pour Postman/Insomnia

## Configuration de base
- Base URL: http://localhost:8080
- Token d'authentification: Bearer secret-token-123

---

## 1. Routes publiques (v1)

### GET - Liste tous les utilisateurs
```
GET http://localhost:8080/v1/users
```

**Réponse attendue:**
```json
{
  "users": [
    {
      "id": 1,
      "name": "Noah Mvondo",
      "email": "noah@example.com",
      "age": 25
    },
    {
      "id": 2,
      "name": "Alice Dupont",
      "email": "alice@example.com",
      "age": 30
    },
    {
      "id": 3,
      "name": "Bob Martin",
      "email": "bob@example.com",
      "age": 28
    }
  ],
  "total": 3
}
```

---

### GET - Récupérer un utilisateur par ID
```
GET http://localhost:8080/v1/users/1
```

**Réponse attendue:**
```json
{
  "id": 1,
  "name": "Noah Mvondo",
  "email": "noah@example.com",
  "age": 25
}
```

---

### POST - Créer un utilisateur
```
POST http://localhost:8080/v1/users
Content-Type: application/json

{
  "name": "Serge Mvondo",
  "email": "serge@example.com",
  "age": 27
}
```

**Réponse attendue:**
```json
{
  "message": "Utilisateur créé avec succès",
  "user": {
    "id": 4,
    "name": "Serge Mvondo",
    "email": "serge@example.com",
    "age": 27
  }
}
```

---

### POST - Test validation (email invalide)
```
POST http://localhost:8080/v1/users
Content-Type: application/json

{
  "name": "Test",
  "email": "invalid-email",
  "age": 25
}
```

**Réponse attendue (Erreur 400):**
```json
{
  "error": "Erreur de validation",
  "details": "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag",
  "help": {
    "name": "requis, min 2 caractères, max 50",
    "email": "requis, format email valide",
    "age": "requis, entre 1 et 150"
  }
}
```

---

### PUT - Mettre à jour un utilisateur
```
PUT http://localhost:8080/v1/users/1
Content-Type: application/json

{
  "name": "Noah Mvondo Updated",
  "email": "noah.updated@example.com",
  "age": 26
}
```

---

### DELETE - Supprimer un utilisateur
```
DELETE http://localhost:8080/v1/users/3
```

**Réponse attendue:**
```json
{
  "message": "Utilisateur Bob Martin supprimé avec succès"
}
```

---

## 2. Routes protégées (v2) - Nécessite authentification

### GET - Liste utilisateurs (avec auth)
```
GET http://localhost:8080/v2/users
Authorization: Bearer secret-token-123
```

---

### GET - Sans authentification (Erreur attendue)
```
GET http://localhost:8080/v2/users
```

**Réponse attendue (Erreur 401):**
```json
{
  "error": "Token d'authentification requis"
}
```

---

### GET - Profil utilisateur
```
GET http://localhost:8080/v2/profile
Authorization: Bearer secret-token-123
```

**Réponse attendue:**
```json
{
  "message": "Profil utilisateur authentifié",
  "user": "Noah Mvondo"
}
```

---

### POST - Créer utilisateur (avec auth)
```
POST http://localhost:8080/v2/users
Authorization: Bearer secret-token-123
Content-Type: application/json

{
  "name": "New User",
  "email": "new@example.com",
  "age": 24
}
```

---

## 3. Routes Admin - Nécessite authentification

### GET - Statistiques système
```
GET http://localhost:8080/admin/stats
Authorization: Bearer secret-token-123
```

**Réponse attendue:**
```json
{
  "total_users": 3,
  "server_uptime": "2h30m",
  "requests_handled": 1523
}
```

---

### GET - Vue admin des utilisateurs
```
GET http://localhost:8080/admin/users
Authorization: Bearer secret-token-123
```

**Réponse attendue:**
```json
{
  "users": [ ... ],
  "total": 3,
  "admin_view": true
}
```

---

## 4. Tests de validation

### Nom trop court (< 2 caractères)
```
POST http://localhost:8080/v1/users
Content-Type: application/json

{
  "name": "A",
  "email": "test@example.com",
  "age": 25
}
```
❌ Erreur 400 attendue

---

### Age hors limites (> 150)
```
POST http://localhost:8080/v1/users
Content-Type: application/json

{
  "name": "Test User",
  "email": "test@example.com",
  "age": 200
}
```
❌ Erreur 400 attendue

---

### Champ manquant
```
POST http://localhost:8080/v1/users
Content-Type: application/json

{
  "name": "Test User",
  "age": 25
}
```
❌ Erreur 400 attendue (email requis)

---

## 5. Tests d'authentification

### Token invalide
```
GET http://localhost:8080/v2/users
Authorization: Bearer wrong-token
```
❌ Erreur 401 attendue

---

### Format de token invalide
```
GET http://localhost:8080/v2/users
Authorization: wrong-format
```
❌ Erreur 401 attendue

---

## Notes
- Tous les endpoints retournent du JSON
- Les codes de statut HTTP sont appropriés (200, 201, 400, 401, 404, 409)
- Les messages d'erreur sont en français et informatifs
- Le middleware Logger affiche les logs dans la console du serveur
