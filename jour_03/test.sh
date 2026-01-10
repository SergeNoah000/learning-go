#!/bin/bash

# Script de test pour l'API Jour 3
# Usage: ./test.sh

BASE_URL="http://localhost:8080"
TOKEN="Bearer secret-token-123"

echo "🧪 Tests de l'API Jour 3 - Middlewares et Groupes de Routes"
echo "============================================================"
echo ""

# Couleurs pour l'affichage
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Fonction pour afficher les résultats
test_endpoint() {
    local name=$1
    local method=$2
    local endpoint=$3
    local data=$4
    local auth=$5
    
    echo -e "${BLUE}Test: $name${NC}"
    
    if [ -z "$auth" ]; then
        if [ -z "$data" ]; then
            response=$(curl -s -w "\n%{http_code}" -X $method "$BASE_URL$endpoint")
        else
            response=$(curl -s -w "\n%{http_code}" -X $method "$BASE_URL$endpoint" \
                -H "Content-Type: application/json" \
                -d "$data")
        fi
    else
        if [ -z "$data" ]; then
            response=$(curl -s -w "\n%{http_code}" -X $method "$BASE_URL$endpoint" \
                -H "Authorization: $TOKEN")
        else
            response=$(curl -s -w "\n%{http_code}" -X $method "$BASE_URL$endpoint" \
                -H "Content-Type: application/json" \
                -H "Authorization: $TOKEN" \
                -d "$data")
        fi
    fi
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n-1)
    
    echo "Status: $http_code"
    echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
    echo ""
}

echo "📝 1. Tests des routes publiques (v1)"
echo "--------------------------------------"

test_endpoint "Lister tous les utilisateurs" "GET" "/v1/users"

test_endpoint "Récupérer utilisateur ID 1" "GET" "/v1/users/1"

test_endpoint "Créer un utilisateur valide" "POST" "/v1/users" \
    '{"name":"Test User","email":"test@example.com","age":25}'

test_endpoint "Créer un utilisateur (email invalide)" "POST" "/v1/users" \
    '{"name":"Test","email":"invalid-email","age":25}'

test_endpoint "Créer un utilisateur (nom trop court)" "POST" "/v1/users" \
    '{"name":"A","email":"test2@example.com","age":25}'

echo "🔐 2. Tests d'authentification (v2)"
echo "------------------------------------"

test_endpoint "Accès sans token (devrait échouer)" "GET" "/v2/users"

test_endpoint "Accès avec token valide" "GET" "/v2/users" "" "auth"

test_endpoint "Profil utilisateur" "GET" "/v2/profile" "" "auth"

echo "👑 3. Tests des routes admin"
echo "----------------------------"

test_endpoint "Stats système" "GET" "/admin/stats" "" "auth"

test_endpoint "Vue admin des utilisateurs" "GET" "/admin/users" "" "auth"

echo ""
echo -e "${GREEN}✅ Tests terminés !${NC}"
echo ""
echo "Note: Pour que ces tests fonctionnent, le serveur doit être en cours d'exécution."
echo "Démarrez le serveur avec: cd jour_03 && go run main.go"
