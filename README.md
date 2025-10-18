# Post Article App

## Overview

- Backend: Go (modular: cmd/, config/, internal/)
- Frontend: React + Vite (Sleek dashboard dark)
- Database: MySQL with golang-migrate migrations
- Docker Compose included for easy local setup

## Local run (without Docker)

1. Start MySQL and create database `article`.
2. Backend:
   ```bash
   cd backend
   go mod tidy
   go run ./cmd
   ```
3. Frontend (dev):
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## Docker

```bash
docker compose up --build
```

Frontend served at http://localhost:3000 (nginx), backend at http://localhost:8080
