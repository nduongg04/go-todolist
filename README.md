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

### GET /todos

Retrieves all todos

-   Response: Array of todo objects
-   Status: 200 OK

### POST /todos

Creates a new todo

-   Request Body: Todo object
-   Response: Created todo object
-   Status: 201 Created

### GET /todos/{id}

Retrieves a specific todo by ID

-   Response: Todo object
-   Status: 200 OK
-   Error: 404 Not Found if todo doesn't exist

### PUT /todos/{id}

Updates an existing todo

-   Request Body: Todo object
-   Response: Updated todo object
-   Status: 200 OK
-   Error: 404 Not Found if todo doesn't exist

### DELETE /todos/{id}

Deletes a todo

-   Response: No content
-   Status: 204 No Content
-   Error: 404 Not Found if todo doesn't exist

## Request/Response Format

### Todo Object Structure

```json
{
	"id": "1",
	"title": "Complete project",
	"description": "Finish the todo list API project",
	"completed": false,
	"created_at": "2024-03-20T15:00:00Z",
	"updated_at": "2024-03-20T15:00:00Z"
}
```

### Create/Update Todo Request Body

```json
{
	"title": "Complete project",
	"description": "Finish the todo list API project",
	"completed": false
}
```

## Error Responses

```json
{
	"error": "Error message here"
}
```

## Development

### Prerequisites

-   Go 1.21 or later
-   Git

### Local Development

1. Clone the repository:

```bash
git clone https://github.com/yourusername/todo-api
cd todo-api
```

2. Install dependencies:

```bash
go mod download
```

3. Run the server:

```bash
go run main.go
```

## License

MIT License - feel free to use this project as you wish.
