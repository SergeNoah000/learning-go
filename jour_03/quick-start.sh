#!/bin/bash

# Script de démarrage rapide pour le jour 3
# Usage: ./quick-start.sh

echo "🚀 Démarrage rapide - Jour 3"
echo "============================"
echo ""

# Vérifier si Go est installé
if ! command -v go &> /dev/null
then
    echo "❌ Go n'est pas installé sur ce système"
    echo "📥 Téléchargez Go depuis: https://golang.org/dl/"
    echo ""
    exit 1
fi

echo "✅ Go est installé: $(go version)"
echo ""

# Installer les dépendances
echo "📦 Installation des dépendances..."
go mod download
echo ""

# Démarrer le serveur
echo "🌐 Démarrage du serveur sur http://localhost:8080"
echo ""
echo "📖 Endpoints disponibles:"
echo "   Public (v1):"
echo "   - GET    http://localhost:8080/v1/users"
echo "   - POST   http://localhost:8080/v1/users"
echo ""
echo "   Protégé (v2 - nécessite auth):"
echo "   - GET    http://localhost:8080/v2/users"
echo "   - GET    http://localhost:8080/v2/profile"
echo ""
echo "   Admin (nécessite auth):"
echo "   - GET    http://localhost:8080/admin/stats"
echo "   - GET    http://localhost:8080/admin/users"
echo ""
echo "🔐 Token pour routes protégées: Bearer secret-token-123"
echo ""
echo "---------------------------------------------------"
echo ""

# Lancer le serveur
go run main.go
