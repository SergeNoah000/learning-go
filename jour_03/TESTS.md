# Tests pour le Jour 3

## Prérequis
- Go 1.21 ou supérieur installé
- Package Gin installé

## Installation de Go
Si Go n'est pas installé, téléchargez-le depuis https://golang.org/dl/

## Tests manuels

### 1. Démarrer le serveur
```bash
cd jour_03
go mod tidy
go run main.go
```

### 2. Tester les routes publiques (v1)

#### Lister tous les utilisateurs
```bash
curl http://localhost:8080/v1/users
```

#### Récupérer un utilisateur spécifique
```bash
curl http://localhost:8080/v1/users/1
```

#### Créer un utilisateur (avec validation)
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Serge Noah","email":"serge@example.com","age":27}'
```

#### Test validation - email invalide
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"invalid-email","age":25}'
```

#### Test validation - nom trop court
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"A","email":"test@example.com","age":25}'
```

#### Test validation - age hors limites
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","age":200}'
```

### 3. Tester les routes protégées (v2)

#### Sans authentification (devrait échouer)
```bash
curl http://localhost:8080/v2/users
```

#### Avec authentification (devrait réussir)
```bash
curl -H "Authorization: Bearer secret-token-123" \
  http://localhost:8080/v2/users
```

#### Token invalide (devrait échouer)
```bash
curl -H "Authorization: Bearer wrong-token" \
  http://localhost:8080/v2/users
```

#### Profil utilisateur authentifié
```bash
curl -H "Authorization: Bearer secret-token-123" \
  http://localhost:8080/v2/profile
```

### 4. Tester les routes admin

#### Stats système (nécessite auth)
```bash
curl -H "Authorization: Bearer secret-token-123" \
  http://localhost:8080/admin/stats
```

#### Vue admin des utilisateurs (nécessite auth)
```bash
curl -H "Authorization: Bearer secret-token-123" \
  http://localhost:8080/admin/users
```

### 5. Tester les middlewares

Le middleware Logger affiche dans la console :
- Méthode HTTP
- Chemin de la requête
- Durée du traitement
- Code de statut

Exemple de sortie attendue :
```
[LOGGER] GET /v1/users - Durée: 234.5µs - Status: 200
[LOGGER] POST /v1/users - Durée: 1.2ms - Status: 201
[LOGGER] GET /v2/users - Durée: 156.7µs - Status: 401
```

## Résultats attendus

### Routes publiques (v1)
- ✅ Accès libre sans authentification
- ✅ Validation des données d'entrée
- ✅ Messages d'erreur détaillés

### Routes protégées (v2 et admin)
- ✅ Refus d'accès sans token
- ✅ Refus d'accès avec token invalide
- ✅ Accès autorisé avec token valide

### Middlewares
- ✅ Logger enregistre toutes les requêtes
- ✅ Auth bloque les accès non autorisés
- ✅ Messages d'erreur clairs

## Points clés de l'implémentation

1. **Middleware Logger** : Log chaque requête avec timing
2. **Middleware Auth** : Vérifie le token Bearer
3. **Groupes de routes** : v1 (public), v2 (auth), admin (auth)
4. **Validation** : Tags binding sur les structs
5. **Gestion d'erreurs** : Messages personnalisés et informatifs
