# Jour 5 - Blog avec PocketBase

Application de blog avec articles, commentaires, likes et upload d'images.

## Installation

```bash
# Télécharger les dépendances
go mod tidy

# Lancer le serveur
go run main.go serve
```

## Accès

- **Frontend** : http://127.0.0.1:8090/
- **Admin** : http://127.0.0.1:8090/_/

## Configuration initiale (Admin)

1. Accéder à http://127.0.0.1:8090/_/
2. Créer un compte admin
3. Créer les collections suivantes :

### Collection `articles`
| Champ | Type | Options |
|-------|------|---------|
| title | Text | Required |
| content | Text | Required |
| category | Select | tech, lifestyle, travel, food, other |
| images | File | Multiple |
| author | Relation | → users |
| likes_count | Number | Default: 0 |

**API Rules :**
- List/View: (public)
- Create/Update/Delete: @request.auth.id != ""

### Collection `comments`
| Champ | Type | Options |
|-------|------|---------|
| content | Text | Required |
| article | Relation | → articles |
| user | Relation | → users |

**API Rules :**
- List/View: (public)
- Create: @request.auth.id != ""
- Update/Delete: @request.auth.id = user

### Collection `likes`
| Champ | Type | Options |
|-------|------|---------|
| article | Relation | → articles |
| user | Relation | → users |

**API Rules :**
- List/View: (public)
- Create: @request.auth.id != ""
- Delete: @request.auth.id = user

**Index unique :** article + user (empêche double like)

## Fonctionnalités

- ✅ Inscription / Connexion
- ✅ Créer des articles avec images multiples
- ✅ Catégories d'articles
- ✅ Commenter les articles
- ✅ Liker / Unliker
- ✅ Articles publics, actions authentifiées
