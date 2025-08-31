# REST API in Gin

Проект REST API на Go с использованием Gin framework и SQLite базы данных.

## Установка зависимостей

```bash
# Установка основных зависимостей
go mod tidy

# Установка инструмента миграций
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Миграции базы данных

### Создание новых миграций

Для создания новых миграций используйте команду:

```bash
~/go/bin/migrate create -ext sql -dir ./migrate/migrations -seq <название_миграции>
```

#### Примеры создания миграций:

```bash
# Создание таблицы пользователей
~/go/bin/migrate create -ext sql -dir ./migrate/migrations -seq create_users_table

# Создание таблицы событий
~/go/bin/migrate create -ext sql -dir ./migrate/migrations -seq create_events_table

# Создание таблицы участников
~/go/bin/migrate create -ext sql -dir ./migrate/migrations -seq create_attendees_table
```

### Применение миграций

Для применения миграций используйте Go-скрипт:

```bash
# Применить все миграции (создать таблицы)
go run migrate/main.go up

# Откатить все миграции (удалить таблицы)
go run migrate/main.go down
```

## Структура проекта

```
.
├── cmd/
│   └── api/
│       └── main.go          # Основной файл приложения
├── migrate/
│   ├── main.go              # Скрипт для выполнения миграций
│   └── migrations/          # Файлы миграций
│       ├── 000001_create_users_table.up.sql
│       ├── 000001_create_users_table.down.sql
│       ├── 000002_create_events_table.up.sql
│       ├── 000002_create_events_table.down.sql
│       ├── 000003_create_attendees_table.up.sql
│       └── 000003_create_attendees_table.down.sql
├── internal/
│   ├── database/            # Конфигурация базы данных
│   └── env/                 # Переменные окружения
├── go.mod
├── go.sum
└── README.md
```

## Структура базы данных

### Таблица `users`
- `id` - первичный ключ
- `username` - уникальное имя пользователя
- `email` - уникальный email
- `password` - хеш пароля
- `created_at` - время создания
- `updated_at` - время обновления

### Таблица `events`
- `id` - первичный ключ
- `owner_id` - внешний ключ на users.id
- `name` - название события
- `description` - описание события
- `date` - дата и время события
- `location` - место проведения

### Таблица `attendees`
- `id` - первичный ключ
- `user_id` - внешний ключ на users.id
- `event_id` - внешний ключ на events.id
- Уникальное ограничение на (user_id, event_id)

## Запуск приложения

```bash
go run cmd/api/main.go
```

## API Endpoints

### Аутентификация

#### Регистрация пользователя
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Логин пользователя
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

Ответ:
```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImVtYWlsIjoiam9obkBleGFtcGxlLmNvbSIsImV4cCI6MTcwNTQ0NzIwMH0.example-signature"
}
```

### Пользователи

#### Создание пользователя
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Получение всех пользователей
```http
GET /api/v1/users
```

#### Получение пользователя по ID
```http
GET /api/v1/users/:id
```

#### Обновление пользователя
```http
PUT /api/v1/users/:id
Content-Type: application/json

{
  "name": "John Updated",
  "email": "john.updated@example.com"
}
```

#### Удаление пользователя
```http
DELETE /api/v1/users/:id
```

### События

#### Создание события
```http
POST /api/v1/events
Content-Type: application/json

{
  "owner_id": 1,
  "name": "Team Meeting",
  "description": "Weekly team sync",
  "date": "2024-01-15T10:00:00Z",
  "location": "Conference Room A"
}
```

#### Получение всех событий
```http
GET /api/v1/events
```

#### Получение события по ID
```http
GET /api/v1/events/:id
```

#### Обновление события
```http
PUT /api/v1/events/:id
Content-Type: application/json

{
  "owner_id": 1,
  "name": "Team Meeting Updated",
  "description": "Updated weekly team sync",
  "date": "2024-01-15T11:00:00Z",
  "location": "Conference Room B"
}
```

#### Удаление события
```http
DELETE /api/v1/events/:id
```

#### Добавление участника к событию
```http
POST /api/v1/events/:id/attendees:user_id
```

#### Получение участников события
```http
GET /api/v1/events/:id/attendees
```

### Участники

#### Создание участника
```http
POST /api/v1/attendees
Content-Type: application/json

{
  "user_id": 1,
  "event_id": 1
}
```

#### Получение всех участников
```http
GET /api/v1/attendees
```

#### Получение участника по ID
```http
GET /api/v1/attendees/:id
```

#### Обновление участника
```http
PUT /api/v1/attendees/:id
Content-Type: application/json

{
  "user_id": 2,
  "event_id": 1
}
```

#### Удаление участника
```http
DELETE /api/v1/attendees/:id
```

## Примечания

- Пароли хешируются с помощью bcrypt для безопасности.
- Используются JWT токены для аутентификации с временем жизни 24 часа.
- Все запросы к базе данных имеют таймаут 3 секунды.
- API использует SQLite в качестве базы данных.
- JWT секрет настраивается через переменную окружения `JWT_SECRET` (по умолчанию "secret").


go run cmd/api/*.go &