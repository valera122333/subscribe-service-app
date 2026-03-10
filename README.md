 Subscriptions Service

REST-сервис для агрегации данных об онлайн-подписках пользователей.  

Реализованы все CRUDL-операции и подсчёт суммарной стоимости подписок за выбранный период.

---

##  Функционал

- **Создание подписки**  
- **Получение подписки по ID**  
- **Список всех подписок**  
- **Удаление подписки по ID**  
- **Подсчёт суммарной стоимости подписок** с фильтром по `user_id` и `service_name`  

---

##  Технологии

- Go 1.25  
- PostgreSQL  
- Docker + Docker Compose  
- Goose для миграций  
- PowerShell для тестов API  

---

##  Конфигурация

Конфигурационные данные хранятся в `.env` файле:

```env
APP_PORT=8080
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=subscriptions
