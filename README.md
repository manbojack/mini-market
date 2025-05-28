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
| №     | Название сервиса         | Язык   | Описание функционала                                                               |
| --    | ------------------------ | ------ | ---------------------------------------------------------------------------------- |
| ✅ 1  | **user-service**         | Python | Регистрация, аутентификация, управление пользователями                             |
| ✅ 2  | **product-service**      | Go     | CRUD для товаров, описание, наличие, цена                                          |
| ⏳ 3  | **inventory-service**    | Python | Управление складом: остатки, резервы, обновление при заказе                        |
| ⏳ 4  | **order-service**        | Go     | Создание заказов, статус, история заказов                                          |
| ⏳ 5  | **payment-service**      | Python | Обработка платежей, интеграция с платежными шлюзами                                |
| ⏳ 6  | **notification-service** | Go     | Отправка уведомлений по email/SMS при событиях (регистрация, заказ, оплата и т.п.) |
| ⏳ 7  | **cart-service**         | Python | Управление корзиной пользователя: добавление/удаление товаров                      |
| ⏳ 8  | **review-service**       | Go     | Отзывы и рейтинги товаров                                                          |
| ⏳ 9  | **analytics-service**    | Python | Аналитика и отчеты: заказы, продажи, популярные товары                             |
| ⏳ 10 | **gateway-service**      | Go     | API Gateway: маршрутизация запросов ко всем сервисам, авторизация                  |


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
