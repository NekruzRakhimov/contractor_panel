# Contractor panel

Микросервис функционала "Панель конторагента":

## Запуск сервиса

* Склонировать репозиторий;
* Задать переменные окружения;
* Установить goose: `make migrations-install-tool`
* Установить все миграции: `make migrations-up`
* Запустить сервис используя команду _serve_: `go run main.go serve`

### Переменные окружения

Все конфигурационные параметры, используемые сервисом, должны быть заданы через переменные окружения.

Name | Type  | Default | Description
---|---|---|---
APP_NAME | string | - | Название сервиса
APP_INSTANCE | string | - | Название экземпляра сервиса
PORT | int | - | Порт, на котором запускается сервис
LOG_LEVEL | string | info | Минимальный уровень логирования (https://github.com/sirupsen/logrus#level-logging)
LOG_PRETTY_PRINT | bool | false | Если true, то производит красивое форматирование JSON записи лога
CORS_ALLOWED_ORIGINS | []string | - | Список разрешенных origins (разделенные пробелом)
CORS_ALLOWED_METHODS | []string | - | Список разрешенны методов (разделенные пробелом)
CORS_ALLOWED_HEADERS | []string | - | Список разрешенных заголовков (разделенные пробелом)
HEALTHCHECK_TIMEOUT | duration | - | Таймаут запроса проверки состояния сервиса
HTTP_REQUEST_TIMEOUT | duration | 60s | Таймаут исходящих http запросов
ATASOURCES_POSTGRES_HOST | string | - | Адрес Postgres
DATASOURCES_POSTGRES_PORT | string | - | Порт Postgres
DATASOURCES_POSTGRES_USER | string | - | Пользователь Postgres
DATASOURCES_POSTGRES_PASSWORD | string | - | Пароль пользователя Postgres
DATASOURCES_POSTGRES_DATABASE | string | - | БД Postgres
DATASOURCES_POSTGRES_SCHEMA | string | - | Схема Postgres
ACCESS_SECRET | string | - | Access token key
REFRESH_SECRET | string | - | Refresh token key
REDIS_DSN | string | - | Redis dsn

## Работа с сервисом

### Логгирование

Логгирование реализовано используя [logrus](https://github.com/sirupsen/logrus).  
Для всего сервиса настроен базовый формат JSON логирования с минимальным количеством параметров. Логирование
бизнес-данных или же любые иные логи, пишушиеся вручную, рекоммендуются делать через форматтер, который находится в
контексте. В таком случае лог будет содержать информацию о пользователе, id запроса и т.д.

### База данных

Для работы с БД используются следующие библиотеки:
PostgreSQL - [pgx](https://github.com/jackc/pgx). Redis - [redis](https://github.com/go-redis/redis) 

Для упрощения маппинга результатов SQL-запросов в модели, необходимо:

+ реализовать интерфейс провайдера модели `model.DbModelProvider`
+ использовать методы `postgres.Query`
+ кастить результат в тип модели (согласно контракту провайдера)

## Тестовая среда

_TODO..._