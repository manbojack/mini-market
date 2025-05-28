### 🎯 Цель проекта
Создание простого интернет-магазина,
реализованного на основе микросервисной архитектуры.

### 🧰 Технологический стек
- <b>Базы данных и кэширование:</b>
    - PostgreSQL
    - Redis

- <b>Системы обмена сообщениями:</b>
    - Kafka

- <b>Инфраструктура и контейнеризация:</b>
    - Docker
    - Minikube/Kubernetes
    - Ansible

- <b>CI/CD:</b>
    - GitHub Actions

### 🔧 Архитектура микросервисов
| Статус | Название сервиса         | Язык   | Описание функционала                                                               |
|------- |--------------------------|--------|-------------------------------------------------------------------------------------|
| ✅    | **user-service**         | Python | Регистрация, аутентификация, управление пользователями                             |
| ✅    | **product-service**      | Go     | CRUD для товаров, описание, наличие, цена                                          |
| 🛠️    | **inventory-service**    | Python | Управление складом: остатки, резервы, обновление при заказе                        |
| 🛠️    | **order-service**        | Go     | Создание заказов, статус, история заказов                                          |
| 🛠️    | **payment-service**      | Python | Обработка платежей, интеграция с платёжными шлюзами                                |
| 🛠️    | **notification-service** | Go     | Отправка уведомлений по email/SMS при событиях (регистрация, заказ, оплата и т.п.) |
| 🛠️    | **cart-service**         | Python | Управление корзиной пользователя: добавление/удаление товаров                      |
| 🛠️    | **review-service**       | Go     | Отзывы и рейтинги товаров                                                          |
| 🛠️    | **analytics-service**    | Python | Аналитика и отчёты: заказы, продажи, популярные товары                             |
| 🛠️    | **gateway-service**      | Go     | API Gateway: маршрутизация запросов ко всем сервисам, авторизация                  |


### 🔁 Связи между сервисами
- user-service ↔ order-service, cart-service
- cart-service ↔ product-service, inventory-service
- order-service ↔ payment-service, inventory-service
- order-service → notification-service
- review-service ↔ user-service, product-service
- analytics-service ← Kafka events от order-service, payment-service, product-service

### 🚧 В планах:  
- <b>Мониторинг</b>
- <b>Сбор логов</b>
