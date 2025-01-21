# Todo List API

A simple RESTful API for managing todos built with Go.

## Setup

1. Install Go (1.21 or later)
2. Clone this repository
3. Install dependencies:

```bash
go mod download
```

4. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

-   `GET /todos` - Get all todos
-   `POST /todos` - Create a new todo
-   `GET /todos/{id}` - Get a specific todo
-   `PUT /todos/{id}` - Update a todo
-   `DELETE /todos/{id}` - Delete a todo

## Request Body Example (POST/PUT)

```json
{
	"title": "Complete project",
	"description": "Finish the todo list API project",
	"completed": false
}
```
