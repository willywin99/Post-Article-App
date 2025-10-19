# Sharing Vision - Sleek Frontend (local assets)

This package contains a React + Vite frontend pre-integrated with a minimal local Sleek-like CSS (dark sidebar).

## Install & Run
1. Install dependencies:
   ```bash
   npm install
   ```
2. Run dev server:
   ```bash
   npm run dev
   ```
3. Open http://localhost:5173

The frontend expects the backend API at: http://localhost:8080/articles
You can set a different API base by setting `VITE_API_BASE` env variable.

Note: This package includes a minimal local CSS that mimics Sleek Dashboard. Replace `public/assets/css/sleek.css` with the real Sleek CSS if you have it.
