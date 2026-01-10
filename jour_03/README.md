# Jour 3 - Middlewares et Groupes de Routes

## Objectifs
- Création de middlewares personnalisés (Logger, Auth)
- Utilisation des middlewares Gin natifs (Recovery, Logger)
- Groupes de routes (v1, v2, admin)
- Application de middlewares à des groupes spécifiques
- Validation des données avec binding tags (required, min, max, email)
- Gestion des erreurs de validation et messages personnalisés

## Fonctionnalités implémentées

### 1. Middlewares personnalisés
- **Logger Middleware** : Log toutes les requêtes HTTP avec méthode, chemin et durée
- **Auth Middleware** : Authentification basique par token (Bearer token)

### 2. Groupes de routes
- **API v1** : Routes publiques avec CRUD utilisateurs
- **API v2** : Routes avec middleware Auth
- **Admin** : Routes admin protégées par authentification

### 3. Validation
- Validation des données avec tags `binding`
- Messages d'erreur personnalisés
- Validation de l'email, champs requis, min/max

## Installation

```bash
go mod download
```

## Exécution

```bash
go run main.go
```

Le serveur démarre sur `http://localhost:8080`

## Endpoints

### API v1 (Public)
- `GET /v1/users` - Liste tous les utilisateurs
- `GET /v1/users/:id` - Récupère un utilisateur
- `POST /v1/users` - Crée un utilisateur
- `PUT /v1/users/:id` - Met à jour un utilisateur
- `DELETE /v1/users/:id` - Supprime un utilisateur

### API v2 (Avec Auth)
- `GET /v2/users` - Liste tous les utilisateurs (nécessite token)
- `POST /v2/users` - Crée un utilisateur (nécessite token)

### Admin (Protégé)
- `GET /admin/stats` - Statistiques système (nécessite token)
- `GET /admin/users` - Liste admin des utilisateurs (nécessite token)

## Authentification

Pour accéder aux routes protégées, ajoutez le header :
```
Authorization: Bearer secret-token-123
```

## Tests

### Test sans authentification
```bash
curl http://localhost:8080/v1/users
```

### Test avec authentification
```bash
curl -H "Authorization: Bearer secret-token-123" http://localhost:8080/v2/users
```

### Test de création avec validation
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Noah","email":"noah@example.com","age":25}'
```

### Test de validation (email invalide)
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Noah","email":"invalid-email","age":25}'
```
