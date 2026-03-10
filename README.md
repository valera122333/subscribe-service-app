 Subscriptions Service

REST-сервис для агрегации данных об онлайн-подписках пользователей.  

Реализованы все CRUDL-операции и подсчёт суммарной стоимости подписок за выбранный период.

---

##  Функционал

- **Создание подписки**  
- **Получение подписки по ID**  
- **Список всех подписок**  
- **Удаление подписки по ID**  
- **Подсчёт суммарной стоимости подписок** с фильтром по user_id и service_name 

---

##  Технологии

- Go 1.25  
- PostgreSQL  
- Docker + Docker Compose  
- Goose для миграций  
- PowerShell для тестов API  

---

##  Конфигурация

Конфигурационные данные хранятся в .env файле:

env
APP_PORT=8080
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=subscriptions


Тестовые операции в файле test-subscriptions.ps1.
выполняются командой .\test-subscriptions.ps1
Запуск контейнера командой docker compose up --build


Вывод подписок http://localhost:8080/subscriptions/list
<img width="2519" height="387" alt="image" src="https://github.com/user-attachments/assets/79143660-6e94-4b43-8b80-04490e2b4d9e" />

 
