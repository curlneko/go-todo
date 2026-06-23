# Gin Todo API

Simple Todo API built with Go and Gin.

## Run

```bash
go run .
```

Server starts on:

```text
http://localhost:8080
```

---

## API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/todos`     | Get all todos     |
| GET    | `/todos/:id` | Get todo by ID    |
| POST   | `/todos`     | Create a new todo |
| PUT    | `/todos/:id` | Update a todo     |
| DELETE | `/todos/:id` | Delete a todo     |

---

## Examples

### Get all todos

```bash
curl http://localhost:8080/todos
```

### Get todo by ID

```bash
curl http://localhost:8080/todos/1
```

### Create todo

```bash
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d "{\"title\":\"Learn Go\",\"completed\":false}"
```

### Create todo (Validation Error)

```bash
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d "{\"completed\":false}"
```

Expected response:

```json
{
  "error": "Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

### Update todo

```bash
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d "{\"title\":\"Learn Gin\",\"completed\":true}"
```

### Delete todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## Project Structure

```text
gin-todo/
├── controllers/
├── models/
├── routes/
├── main.go
├── go.mod
└── README.md
```
