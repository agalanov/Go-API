# Diflow API - Go Implementation

Это Go-версия API, основанная на PHP Yii2 API проекте.

## Структура проекта

```
go/
├── main.go                 # Точка входа приложения
├── config/                 # Конфигурация
│   └── config.go
├── database/              # Работа с базой данных
│   └── database.go
├── models/                 # Модели данных
│   ├── user.go
│   ├── token.go
│   ├── building.go
│   ├── equipment.go
│   └── ...
├── controllers/            # Контроллеры
│   ├── base_controller.go
│   ├── user_controller.go
│   ├── building_controller.go
│   └── ...
├── middleware/             # Middleware
│   ├── auth.go
│   ├── cors.go
│   └── error_handler.go
├── routes/                 # Маршруты
│   └── routes.go
└── utils/                  # Утилиты
    ├── jwt.go
    └── response.go
```

## Установка и запуск

1. Установите зависимости:
```bash
cd go
go mod download
```

2. Создайте файл `.env` на основе `.env.example`:
```bash
cp .env.example .env
```

3. Настройте переменные окружения в `.env` файле.

4. Запустите приложение:
```bash
go run main.go
```

Или скомпилируйте и запустите:
```bash
go build -o api main.go
./api
```

## API Endpoints

### Аутентификация

- `POST /api/v1/sso/login` - Вход через SSO
- `POST /api/v1/oauth/login` - Вход через OAuth
- `POST /api/v1/oauth/register` - Регистрация
- `POST /api/v1/oauth/refresh` - Обновление токена

### Пользователи

- `GET /api/v1/user/:id` - Получить пользователя

### Здания

- `GET /api/v1/building` - Список зданий
- `POST /api/v1/building` - Создать здание
- `GET /api/v1/building/:id` - Получить здание
- `PUT /api/v1/building/:id` - Обновить здание
- `DELETE /api/v1/building/:id` - Удалить здание

### Оборудование

- `GET /api/v1/equipment` - Список оборудования
- `POST /api/v1/equipment` - Создать оборудование
- `GET /api/v1/equipment/:id` - Получить оборудование
- `PUT /api/v1/equipment/:id` - Обновить оборудование
- `DELETE /api/v1/equipment/:id` - Удалить оборудование

### Этажи

- `GET /api/v1/floor` - Список этажей
- `POST /api/v1/floor` - Создать этаж
- `GET /api/v1/floor/:id` - Получить этаж
- `PUT /api/v1/floor/:id` - Обновить этаж
- `DELETE /api/v1/floor/:id` - Удалить этаж

И другие endpoints согласно PHP API.

## Аутентификация

API использует JWT токены для аутентификации. Токен можно передать:
- В заголовке `Authorization: Bearer <token>`
- В query параметре `?access_token=<token>`

## База данных

API использует PostgreSQL. Схема базы данных создается автоматически при первом запуске через GORM AutoMigrate.

## Swagger документация

После запуска сервера Swagger документация доступна по адресу:
- http://localhost:8080/swagger/index.html

Для генерации Swagger документации:
```bash
make swagger
# или
swag init -g main.go
```

## Docker

### Запуск с Docker Compose

1. Создайте файл `.env` или используйте `.env.docker`:
```bash
cp .env.docker .env
```

2. Запустите сервисы:
```bash
docker-compose up -d
```

3. Остановите сервисы:
```bash
docker-compose down
```

4. Просмотр логов:
```bash
docker-compose logs -f api
```

### Разработка с Docker

Для разработки используйте `docker-compose.dev.yml`:
```bash
docker-compose -f docker-compose.dev.yml up
```

### Сборка образа

```bash
docker-compose build
# или
make docker-build
```

## Технологии

- **Gin** - HTTP веб-фреймворк
- **GORM** - ORM для работы с базой данных
- **JWT** - Аутентификация
- **PostgreSQL** - База данных
- **Swagger** - API документация
- **Docker** - Контейнеризация

