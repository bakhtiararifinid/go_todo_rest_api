# Go Todo Rest API

This is a todo app rest api project build with go

## Prequisite

To run this project, you need [go](https://go.dev/) and [mysql](https://mysql.com/) already installed in your computer.

## API Endpoints

- /todos
  - GET - Get list of all todos
  - POST - Add new todo

- /todos/:id
  - GET - Get todo by id
  - PUT - Update todo title
  - PATCH - Toggle complete todo
  - DELETE - Delete todo

## To run this project

1. Setup mysql database for this project.
```
$ mysql -u root -p
Enter password:

mysql> create database todo_app;
```

2. Create a file called .env to hold your environment variables.
```
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=todo_app
DB_USER=root
DB_PASSWORD=your_database_user_password
```

3. Run your project.
```
$ go run .
```