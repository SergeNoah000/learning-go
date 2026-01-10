# 🚀 Jour 3 - Middlewares, Groupes de Routes et Validation

## 📋 Vue d'ensemble

Le **Jour 3** implémente une API REST avec middlewares personnalisés, authentification, groupes de routes et validation avancée des données. Le projet représente une progression significative vers une architecture API professionnelle.

---

## 📁 Structure du projet

```
jour_03/
├── main.go                  # Code principal de l'API
├── go.mod                   # Dépendances Go
├── test.sh                  # Script de test automatisé
├── README.md                # Documentation générale
├── TESTS.md                 # Guide de tests
├── EXEMPLES_REQUETES.md     # Exemples de requêtes HTTP
├── IMPLEMENTATION.md        # Détails d'implémentation
└── EVOLUTION.md             # Évolution du projet (jours 1-3)
```

---

## ✨ Fonctionnalités principales

### 1. 🔌 Middlewares personnalisés

#### Logger Middleware
- Enregistre toutes les requêtes HTTP
- Calcule la durée d'exécution
- Affiche méthode, chemin, durée et status HTTP

#### Auth Middleware
- Authentification par Bearer token
- Validation du format du token
- Blocage des accès non autorisés
- Messages d'erreur clairs

### 2. 🗂️ Groupes de routes

#### API v1 - Public
- Routes CRUD accessibles sans authentification
- Validation des données d'entrée
- 5 endpoints disponibles

#### API v2 - Protégé
- Routes nécessitant authentification
- Middleware Auth appliqué au groupe
- 3 endpoints disponibles

#### Admin - Administrateur
- Routes pour fonctions administratives
- Protection par authentification
- 2 endpoints disponibles

### 3. ✅ Validation avancée

```go
type User struct {
    Name  string `binding:"required,min=2,max=50"`
    Email string `binding:"required,email"`
    Age   int    `binding:"required,min=1,max=150"`
}
```

- Validation automatique avec tags binding
- Messages d'erreur détaillés
- Aide contextuelle sur les règles

### 4. 🔐 Sécurité

- Authentification Bearer token
- Validation stricte des entrées
- Gestion d'erreurs robuste
- Séparation des routes publiques/privées

---

## 🎯 Endpoints disponibles

### Routes publiques (v1)
```
GET    /v1/users        # Liste tous les utilisateurs
GET    /v1/users/:id    # Récupère un utilisateur
POST   /v1/users        # Crée un utilisateur
PUT    /v1/users/:id    # Met à jour un utilisateur
DELETE /v1/users/:id    # Supprime un utilisateur
```

### Routes protégées (v2)
```
GET  /v2/users    # Liste utilisateurs (auth requise)
POST /v2/users    # Crée utilisateur (auth requise)
GET  /v2/profile  # Profil utilisateur (auth requise)
```

### Routes admin
```
GET /admin/stats  # Statistiques système (auth requise)
GET /admin/users  # Vue admin utilisateurs (auth requise)
```

---

## 🛠️ Installation et exécution

### Prérequis
- Go 1.21 ou supérieur
- Package Gin installé

### Installation
```bash
cd jour_03
go mod download
```

### Démarrage
```bash
go run main.go
```

Le serveur démarre sur http://localhost:8080

---

## 🧪 Tests

### Test manuel avec curl
```bash
# Route publique
curl http://localhost:8080/v1/users

# Route protégée
curl -H "Authorization: Bearer secret-token-123" \
  http://localhost:8080/v2/users

# Création avec validation
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com","age":25}'
```

### Script de test automatisé
```bash
chmod +x test.sh
./test.sh
```

---

## 📊 Statistiques du projet

- **Lignes de code:** ~300
- **Middlewares:** 2 (Logger, Auth)
- **Groupes de routes:** 3 (v1, v2, admin)
- **Endpoints totaux:** 11
- **Règles de validation:** 7
- **Codes HTTP gérés:** 6 (200, 201, 400, 401, 404, 409)

---

## 🎓 Concepts Go maîtrisés

### Niveau intermédiaire
- ✅ Middlewares avec `gin.HandlerFunc`
- ✅ Groupes de routes avec `r.Group()`
- ✅ Application de middlewares à des groupes
- ✅ Tags de validation complexes
- ✅ Gestion avancée des erreurs
- ✅ Manipulation de headers HTTP
- ✅ Time et calcul de durées
- ✅ Manipulation de strings avancée

### Architecture
- ✅ Séparation des concerns
- ✅ Versioning d'API (v1, v2)
- ✅ Middleware chain
- ✅ Authentification et autorisation
- ✅ Messages d'erreur informatifs

---

## 📝 Token pour les tests

Pour accéder aux routes protégées, utilisez:
```
Authorization: Bearer secret-token-123
```

---

## 🔄 Comparaison avec les jours précédents

| Métrique | Jour 1 | Jour 2 | Jour 3 |
|----------|--------|--------|--------|
| Endpoints | 2 | 6 | 11 |
| Middlewares | 0 | Natifs | 2 custom |
| Validation | ❌ | Basique | Avancée |
| Auth | ❌ | ❌ | ✅ |
| Groupes | ❌ | ❌ | ✅ |

---

## 🚀 Prochaines étapes (Jour 4)

Selon le compte rendu développeur :
- 🗄️ Connexion à une base de données (MySQL/PostgreSQL)
- 🔧 Utilisation d'un ORM (GORM)
- 💾 Opérations CRUD avec persistance
- 🔄 Migrations de base de données

---

## 📚 Documentation complète

- **[README.md](README.md)** - Documentation générale
- **[TESTS.md](TESTS.md)** - Guide de tests détaillé
- **[EXEMPLES_REQUETES.md](EXEMPLES_REQUETES.md)** - Exemples de requêtes
- **[IMPLEMENTATION.md](IMPLEMENTATION.md)** - Détails techniques
- **[EVOLUTION.md](EVOLUTION.md)** - Historique du projet

---

## 💡 Points forts du projet

1. ✅ **Architecture propre** : Séparation claire entre routes publiques et privées
2. ✅ **Sécurité** : Authentification et validation robustes
3. ✅ **Maintenabilité** : Code organisé avec middlewares réutilisables
4. ✅ **Extensibilité** : Facile d'ajouter de nouveaux groupes ou middlewares
5. ✅ **Documentation** : Documentation complète et exemples
6. ✅ **Tests** : Script de test automatisé inclus

---

## 👨‍💻 Auteur

Noah Mvondo Serge - Projet Afaapay (OSSECA)

---

## 📅 Date

08 janvier 2026 - Jour 3 de l'apprentissage Go

---

**✅ Status : IMPLÉMENTÉ ET DOCUMENTÉ**
