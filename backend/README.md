# Post Article Backend API

A Golang backend service for managing articles with RESTful API endpoints.

## 🚀 Tech Stack

- Golang, Gin, GORM, MySQL, CORS

## 📋 Features

- RESTful API endpoints
- CRUD operations for articles
- Input validation
- Pagination support
- Database migrations
- CORS configured

## 🏗️ Project Structure

```
backend/
├── cmd/
├── config/
├── db/
├── internal/
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 🛠️ Installation

```bash
git clone https://github.com/willywin99/Post-Article-App.git
cd Post-Article-App/backend
go mod tidy
```

### Setup environment variables

Create `.env` file:

```
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=article
DB_PORT=3306
```

Create database:

```sql
CREATE DATABASE article;
```

### Running the Application

```bash
go run ./cmd
```

Server runs on http://localhost:8080

## 📡 API Endpoints

- POST /article/ → Create article
- GET /article → Get paginated articles
- GET /article/:id → Get specific article
- PUT /article/:id → Update article
- DELETE /article/:id → Delete article

### Request Format

```json
{
	"title": "Article title with at least 20 characters",
	"content": "Article content with at least 200 characters...",
	"category": "Technology",
	"status": "publish"
}
```

### Response Format

```json
{
	"id": 1,
	"title": "Article title",
	"content": "Article content",
	"category": "Technology",
	"status": "publish",
	"created_date": "2024-01-01T00:00:00Z",
	"updated_date": "2024-01-01T00:00:00Z"
}
```

### Validation Rules

- Title: min 20 characters
- Content: min 200 characters
- Category: min 3 characters
- Status: must be publish, draft, or thrash

## 🧪 Testing

Use Postman or curl:

```bash
curl "http://localhost:8080/article?limit=5&offset=0"
```

Postman Collection:
[Post-Article-App REST API.postman_collection.json](/backend/Post-Article-App%20REST%20API.postman_collection.json)

## 📝 License

MIT License
