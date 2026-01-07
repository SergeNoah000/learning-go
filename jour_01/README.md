# Jour 01 - Introduction à Go et Gin

## Ce qui a été appris

### 1. Les bases de Go
- Installation de Go
- Premier programme Hello World
- Syntaxe de base (package, import, func)

### 2. Introduction au framework Gin
- Installation: `go get -u github.com/gin-gonic/gin`
- Création d'un serveur HTTP basique
- Routes GET simples
- Réponses JSON avec `c.JSON()`

## Fichiers créés
- `main.go` - Programme Hello World en Go
- `server.go` - Serveur HTTP avec Gin

## Comment exécuter

### Hello World
```bash
go run main.go
```

### Serveur Gin
```bash
go run server.go
```

Puis ouvrez votre navigateur sur:
- http://localhost:8080/
- http://localhost:8080/hello

## Concepts clés
- **gin.Default()**: Crée une instance Gin avec Logger et Recovery middleware
- **r.GET()**: Définit une route qui répond aux requêtes GET
- **c.JSON()**: Renvoie une réponse JSON
- **gin.H**: Type map utilisé pour créer des objets JSON
