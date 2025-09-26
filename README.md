# PetProject

PetProject — учебный проект на Go, реализующий CRUD для пользователей и задач.  
Проект построен по слоям (service, repository, handlers), покрыт тестами и использует GORM для работы с БД.

---

## Стек

- **GORM**
- **PostgreSQL**
- **Testify**
- **Echo**
- **Docker**

### Структура проекта

PetProject/
├── cmd/app/               # точка входа (main.go)
├── internal/
│   ├── db/                # работа с БД
│   ├── handlers/          # HTTP-обработчики
│   ├── taskService/       # бизнес-логика задач
│   ├── userService/       # бизнес-логика пользователей
│   └── web/               # веб-слой (маршруты / API)

#### Запуск
Пишем в терминале:
1. make create-container      //Docker
2. make migrate
3. make run
4. Готовая коллекция для Postman: https://github.com/AbdulkaevTimur/PetProject/blob/master/PetProject.postman_collection.json




