# 📝 Post Article App

A full-stack article management application built with Golang backend and React frontend.

## 🎯 Overview

This application allows users to create, read, update, and delete articles with a beautiful dashboard interface. It includes features like article status management (Publish, Draft, Thrash), form validation, and pagination.

## 🚀 Quick Start

### Prerequisites

- Go 1.19+
- Node.js 16+
- MySQL 5.7+

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/willywin99/Post-Article-App.git
   cd Post-Article-App
   ```

2. #### Setup Backend

```bash
cd backend
go mod tidy
# Configure .env file
go run ./cmd
```

3. #### Setup Frontend

```bash
cd frontend
npm install
npm run dev
```

4. #### Access the application

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

## 📚 Documentation

- Backend Documentation
- Frontend Documentation

## 🏗️ Architecture

### Backend (Golang)

- Framework: Gin
- ORM: GORM
- Database: MySQL
- Validation: Custom validation middleware

### Frontend (React)

- Build Tool: Vite
- UI Library: SleekDashboard
- HTTP Client: Axios
- Routing: React Router

## 📡 API Specification

| Method | Endpoint                   | Description    |
| ------ | -------------------------- | -------------- |
| POST   | /article/                  | Create article |
| GET    | /article?limit=10&offset=0 | List articles  |
| GET    | /article/:id               | Get article    |
| PUT    | /article/:id               | Update article |
| DELETE | /article/:id               | Delete article |

### Validation Rules

- Title: Required, min 20 characters
- Content: Required, min 200 characters
- Category: Required, min 3 characters
- Status: Must be "publish", "draft", or "thrash"

## 🗃️ Database Schema

```sql
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    status VARCHAR(100) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 🎨 Features

### Backend Features

- ✅ RESTful API design
- ✅ Input validation
- ✅ Pagination support
- ✅ Error handling
- ✅ CORS configuration

### Frontend Features

- ✅ Dashboard with tabs
- ✅ Form validation
- ✅ Article preview
- ✅ Responsive design
- ✅ Status management

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📝 License

MIT License

## 👥 Author

Willywin99  
GitHub: @willywin99
