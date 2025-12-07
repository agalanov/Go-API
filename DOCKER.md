# Docker Guide

## Быстрый старт

1. Убедитесь, что у вас установлен Docker и Docker Compose

2. Скопируйте файл с переменными окружения:
```bash
cp .env.docker .env
```

3. Запустите сервисы:
```bash
docker-compose up -d
```

4. Проверьте статус:
```bash
docker-compose ps
```

5. Просмотрите логи:
```bash
docker-compose logs -f api
```

## Доступ к сервисам

- **API**: http://localhost:8080
- **Swagger**: http://localhost:8080/swagger/index.html
- **PostgreSQL**: localhost:5432

## Остановка

```bash
docker-compose down
```

Для удаления всех данных (включая базу данных):
```bash
docker-compose down -v
```

## Пересборка

Если вы изменили код, пересоберите образ:
```bash
docker-compose build
docker-compose up -d
```

## Разработка

Для разработки с hot-reload используйте:
```bash
docker-compose -f docker-compose.dev.yml up
```

Этот файл монтирует исходный код как volume, что позволяет видеть изменения без пересборки.

## Переменные окружения

Основные переменные окружения можно настроить в файле `.env`:

- `POSTGRES_USER` - пользователь PostgreSQL
- `POSTGRES_PASSWORD` - пароль PostgreSQL
- `POSTGRES_DATABASE_API` - имя базы данных
- `JWT_SECRET` - секретный ключ для JWT (обязательно измените в продакшене!)
- `PORT` - порт API сервера

## Troubleshooting

### База данных не запускается

Проверьте логи:
```bash
docker-compose logs postgres
```

Убедитесь, что порт 5432 не занят другим процессом.

### API не может подключиться к базе данных

Проверьте, что:
1. PostgreSQL контейнер запущен и здоров
2. Переменные окружения настроены правильно
3. Сеть между контейнерами работает

### Swagger не работает

Убедитесь, что документация сгенерирована:
```bash
make swagger
```

Или внутри контейнера:
```bash
docker-compose exec api swag init -g main.go
```

