# HL_project_management
# Сервис для управления проектами

## Цель

Создать REST API для управления задачами, включающими три сущности: Пользователь, Задача и Проект. Сервис  поддерживает операции CRUD и запущен с использованием Docker Compose и Makefile.

## Описание

Сервис управления задачами позволяет вести учет пользователей, задач и проектов, поддерживая создание, чтение, обновление и удаление данных (CRUD). Также реализованы пути для получения списка задач конкретного пользователя и списка задач в проекте.  https://hl-project-management.onrender.com/swagger/index.html

## Сущности

### Пользователь

- **ID**: уникальный идентификатор пользователя
- **Имя**: полное имя пользователя
- **Email**: электронная почта пользователя
- **Дата регистрации**: дата регистрации пользователя
- **Роль**: роль пользователя в системе (например, администратор, менеджер, разработчик)

### Задача

- **ID**: уникальный идентификатор задачи
- **Название**: название задачи
- **Описание**: краткое описание задачи
- **Приоритет**: уровень приоритета задачи (низкий, средний, высокий)
- **Состояние**: состояние задачи (новая, в процессе, завершена)
- **Ответственный**: идентификатор пользователя, ответственного за задачу
- **Проект**: идентификатор проекта, к которому относится задача
- **Дата создания**: дата создания задачи
- **Дата завершения**: дата завершения задачи

### Проект

- **ID**: уникальный идентификатор проекта
- **Название**: название проекта
- **Описание**: краткое описание проекта
- **Дата начала**: дата начала проекта
- **Дата завершения**: дата завершения проекта
- **Менеджер**: идентификатор пользователя-менеджера проекта

## API Пути

### /users

- GET /users: получить список всех пользователей
- POST /users: создать нового пользователя
- GET /users/{id}: получить данные конкретного пользователя
- PUT /users/{id}: обновить данные конкретного пользователя
- DELETE /users/{id}: удалить конкретного пользователя
- GET /users/{id}/tasks: получить список задач конкретного пользователя
- GET /users/search?name={name}: найти пользователей по имени
- GET /users/search?email={email}: найти пользователей по электронной почте

### /tasks

- GET /tasks: получить список всех задач
- POST /tasks: создать новую задачу
- GET /tasks/{id}: получить данные конкретной задачи
- PUT /tasks/{id}: обновить данные конкретной задачи
- DELETE /tasks/{id}: удалить конкретную задачу
- GET /tasks/search?title={title}: найти задачи по названию
- GET /tasks/search?status={status}: найти задачи по состоянию
- GET /tasks/search?priority={priority}: найти задачи по приоритету
- GET /tasks/search?assignee={userId}: найти задачи по идентификатору ответственного
- GET /tasks/search?project={projectId}: найти задачи по идентификатору проекта

### /projects

- GET /projects: получить список всех проектов
- POST /projects: создать новый проект
- GET /projects/{id}: получить данные конкретного проекта
- PUT /projects/{id}: обновить данные конкретного проекта
- DELETE /projects/{id}: удалить конкретный проект
- GET /projects/{id}/tasks: получить список задач в проекте
- GET /projects/search?title={title}: найти проекты по названию
- GET /projects/search?manager={userId}: найти проекты по идентификатору менеджера

## Ответы HTTP

- GET, PUT, DELETE: 200 при успешном выполнении
- POST: 201 при успешном создании
- 404: ресурс не найден
- 400: некорректный запрос
- 405: метод не поддерживается

## Технические требования

- Создание веб-приложения с CRUD функционалом
- Работа с базой данных PostgreSQL
- Контейнеризация и контроль версий
- Тестирование и оптимизация
- Интеграция и развертывание
- Документация

## Установка и запуск

1. Клонировать репозиторий
   ```sh
   git clone https://github.com/username/project-management.git
   cd project-management
   ```
   Не забудьте создать .env файл!
2. Собрать проект
   ```sh
   make build
   ```
3. Запустить проект
   ```sh
   make up
   ```
4. Остановить проект
   ```sh
   make down
   ```
# Документация API
Документация API доступна по пути /swagger/ после запуска сервера.  https://hl-project-management.onrender.com/swagger/index.html
