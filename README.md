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

<h3>🔧 Архитектура микросервисов</h3>
<table>
  <thead>
    <tr>
      <th>№</th>
      <th>Название сервиса</th>
      <th>Язык</th>
      <th>Описание функционала</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>✅ 1</td>
      <td><strong>user-service</strong></td>
      <td>Python</td>
      <td>Регистрация, аутентификация, управление пользователями</td>
    </tr>
    <tr>
      <td>✅ 2</td>
      <td><strong>product-service</strong></td>
      <td>Go</td>
      <td>CRUD для товаров, описание, наличие, цена</td>
    </tr>
    <tr>
      <td>🛠️ 3</td>
      <td><strong>inventory-service</strong></td>
      <td>Python</td>
      <td>Управление складом: остатки, резервы, обновление при заказе</td>
    </tr>
    <tr>
      <td>🛠️ 4</td>
      <td><strong>order-service</strong></td>
      <td>Go</td>
      <td>Создание заказов, статус, история заказов</td>
    </tr>
    <tr>
      <td>🛠️ 5</td>
      <td><strong>payment-service</strong></td>
      <td>Python</td>
      <td>Обработка платежей, интеграция с платёжными шлюзами</td>
    </tr>
    <tr>
      <td>🛠️ 6</td>
      <td><strong>notification-service</strong></td>
      <td>Go</td>
      <td>Отправка уведомлений по email/SMS при событиях (регистрация, заказ, оплата и т.п.)</td>
    </tr>
    <tr>
      <td>🛠️ 7</td>
      <td><strong>cart-service</strong></td>
      <td>Python</td>
      <td>Управление корзиной пользователя: добавление/удаление товаров</td>
    </tr>
    <tr>
      <td>🛠️ 8</td>
      <td><strong>review-service</strong></td>
      <td>Go</td>
      <td>Отзывы и рейтинги товаров</td>
    </tr>
    <tr>
      <td>🛠️ 9</td>
      <td><strong>analytics-service</strong></td>
      <td>Python</td>
      <td>Аналитика и отчёты: заказы, продажи, популярные товары</td>
    </tr>
    <tr>
      <td>🛠️ 10</td>
      <td><strong>gateway-service</strong></td>
      <td>Go</td>
      <td>API Gateway: маршрутизация запросов ко всем сервисам, авторизация</td>
    </tr>
  </tbody>
</table>

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
