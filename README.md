# to-do-list

Простой REST API для управления пользователями и задачами с использованием Gin, GORM и MySQL.

## Возможности

Регистрация и аутентификация пользователей 

Создание задач


     `POST /register` - Создать новый аккаунт  
     
     `POST /login` - Аутентификация пользователя
     
     `POST /logout` - Выход из сессии 
     
     `GET /:username/profile` - Информация о профиле 
     
     `POST /:username/addTask` - Создать новую задачу пользователя

### Установка

1. Клонируйте репозиторий:
```
git clone https://github.com/gegxkss/to-do-list.git

cd to-do-list
```

2. Установите зависимости:
```
go mod download
```

3. Run database migrations:
```
go run migrate/migrate.go
```

4. Запустите сервер:
```
go run main.go

http://localhost:8080 
```
