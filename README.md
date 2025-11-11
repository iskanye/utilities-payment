# Сервис по симуляции оплаты услуг ЖКХ

## Технологический стек

### Веб-фрейморки

* [Gin](https://github.com/gin-gonic/gin) - высокопроизводительный фреймворк, предназначенный для создания веб-приложений
* [gRPC](https://grpc.io/) - фреймворк для быстрой и бесперебойной связи микросервисов между собой

### Авторизация

* [jwt-go](https://github.com/golang-jwt/jwt) - Библиотека для генерации JWT
* [crypto](https://pkg.go.dev/golang.org/x/crypto) - Стандартная библиотека для шифрования

### Базы данных

* [SQLite](https://github.com/glebarez/go-sqlite) - встраиваемая легковестная СУБД (используется для хранения данных пользователей)
* [PostgeSQL](https://github.com/lib/pq) - мощная и популярная СУБД (используется для хранения счётов)
* [Migrate](https://github.com/golang-migrate/migrate) - библиотека для написания миграций баз данных

### Тестирование

* [Testify](https://github.com/stretchr/testify) - библиотека для эффективного написания тестов на Go

## Сборка и запуск

### 1. Клонирование репозитория с подмодулями

```bash
git clone --recurse-submodules https://github.com/iskanye/utilities-payment.git
cd utilities-payment
```

Если репозиторий был склонирован без подмодулей, выполните:

```bash
git submodule init
git submodule update
```

### 2. Сборка

Для сборки потребуется установленный [Docker](https://www.docker.com/).

**`ВАЖНО`**: Заранее при сборке установите значение параметров окружения `AUTH_SECRET`, `POSTGRES_USER`, `POSTGRES_PASSWORD` и `POSTGRES_DB` для корректной работы сервиса, например:

```bash
export AUTH_SECRET="SUPER-SECRET"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="postgres"
export POSTGRES_DB="postgres"
```

Или в среде Powershell:

```powershell
$Env:AUTH_SECRET="SUPER-SECRET"
$Env:POSTGRES_USER="postgres"
$Env:POSTGRES_PASSWORD="postgres"
$Env:POSTGRES_DB="postgres"
```

После выполните следующую команду для сборки сервисов:

```bash
docker compose build
```

### 3. Запуск

Для запуска сервисов выполните:

```bash
docker compose up
```

После запуска с сервисом можно будет работать по адресу `localhost:8080`
