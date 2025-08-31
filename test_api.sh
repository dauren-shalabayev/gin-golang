#!/bin/bash

# Тестовый скрипт для API логина
# Убедитесь, что сервер запущен на порту 8080

BASE_URL="http://localhost:8080/api/v1"

echo "🧪 Тестирование API логина"
echo "=========================="

# Тест 1: Регистрация пользователя
echo -e "\n1️⃣ Регистрация пользователя:"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }')

echo "Ответ: $REGISTER_RESPONSE"

# Тест 2: Логин пользователя
echo -e "\n2️⃣ Логин пользователя:"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }')

echo "Ответ: $LOGIN_RESPONSE"

# Тест 3: Логин с неправильным паролем
echo -e "\n3️⃣ Логин с неправильным паролем:"
WRONG_PASSWORD_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "wrongpassword"
  }')

echo "Ответ: $WRONG_PASSWORD_RESPONSE"

# Тест 4: Логин с несуществующим email
echo -e "\n4️⃣ Логин с несуществующим email:"
WRONG_EMAIL_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "nonexistent@example.com",
    "password": "password123"
  }')

echo "Ответ: $WRONG_EMAIL_RESPONSE"

echo -e "\n✅ Тестирование завершено!"
echo -e "\n💡 Для запуска тестов убедитесь, что сервер запущен:"
echo "   go run cmd/api/main.go"
