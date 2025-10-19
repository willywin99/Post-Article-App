# Post Article Frontend

React frontend application for managing articles with SleekDashboard UI.

## 🚀 Tech Stack

- React, Vite, SleekDashboard, Axios, React Router

## 📋 Features

- Dashboard with tabs (Published, Drafts, Thrash)
- Add/Edit articles with validation
- Preview articles
- Responsive design
- Status management

## 🏗️ Project Structure

```
frontend/
├── src/
│   ├── components/UI/
│   ├── pages/
│   │   ├── AllPosts.jsx
│   │   ├── AddEdit.jsx
│   │   └── Preview.jsx
│   ├── api.js
│   └── main.jsx
├── public/
├── package.json
└── vite.config.mjs
```

## 🛠️ Installation

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on http://localhost:5173

## 🔌 API Integration

- fetchArticles(limit, offset, status)
- fetchArticle(id)
- createArticle(articleData)
- updateArticle(id, articleData)
- deleteArticle(id)

## 🎨 UI Components & Validation

- Title: min 20 chars
- Content: min 200 chars
- Category: min 3 chars
- Status: required (publish/draft/thrash)

## 📝 License

MIT License

## 👥 Author

Willywin99  
GitHub: @willywin99
