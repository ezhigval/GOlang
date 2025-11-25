# DAILY CHECKLIST --- GO BACKEND INTENSIVE (Дни 15--28)

Фаза: УСКОРЕНИЕ → PRE-MID

## День 15 --- Архитектура

-   [ ] Создать шаблон проекта с handlers/service/repo
-   [ ] Настроить конфиги (Viper)
-   [ ] Добавить zap logger
-   [ ] Подготовить Swagger

## День 16 --- Бизнес-логика Order Service

-   [ ] Создать сущности товаров
-   [ ] Создать сущность заказа
-   [ ] Проработать схемы БД
-   [ ] Написать SQL миграции

## День 17 --- Handlers (REST)

-   [ ] GET /products
-   [ ] POST /products
-   [ ] GET /orders
-   [ ] POST /orders
-   [ ] Проверка авторизации

## День 18 --- Repository

-   [ ] product repo
-   [ ] order repo
-   [ ] transaction manager
-   [ ] логирование SQL

## День 19 --- Service Layer

-   [ ] CreateOrder()
-   [ ] ValidateOrder()
-   [ ] CalculateTotalPrice()
-   [ ] Error wrapping

## День 20 --- Транзакции

-   [ ] ACID
-   [ ] Begin/Commit/Rollback
-   [ ] Idempotency ключ
-   [ ] Обработка ошибок

## День 21 --- Интеграция Redis

-   [ ] Кеширование товаров
-   [ ] TTL
-   [ ] Инвалидация
-   [ ] Rate limit

## День 22 --- RabbitMQ 1

-   [ ] Установить RabbitMQ
-   [ ] Поднять очередь order_created
-   [ ] Создать продьюсер

## День 23 --- RabbitMQ 2

-   [ ] Создать consumer
-   [ ] Обработка заказа
-   [ ] Retry + DLQ

## День 24 --- Docker-compose 5 сервисов

-   [ ] api
-   [ ] postgres
-   [ ] redis
-   [ ] rabbit
-   [ ] prometheus scaffolding

## День 25 --- Тесты

-   [ ] Unit-тесты service слоя
-   [ ] Интеграционные тесты order
-   [ ] mock repo

## День 26 --- Логи + Метрики

-   [ ] Метрики: latency, rps
-   [ ] Встроить Prometheus
-   [ ] Настроить запросы

## День 27 --- Ревью и рефакторинг

-   [ ] Пройтись по слоям
-   [ ] Убрать дублирование
-   [ ] Привести структуру к production-стилю

## День 28 --- Финал ORDER SERVICE

-   [ ] Полный тест функционала
-   [ ] Документация swagger
-   [ ] README
