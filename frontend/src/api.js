const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080';

export async function fetchArticles(limit=10, offset=0, status='') {
  const url = new URL(`${API_BASE}/articles`);
  url.searchParams.set('limit', limit);
  url.searchParams.set('offset', offset);
  if (status) url.searchParams.set('status', status);
  const res = await fetch(url);
  if (!res.ok) throw new Error('failed to fetch');
  return res.json();
}

export async function fetchArticle(id) {
  const res = await fetch(`${API_BASE}/articles/${id}`);
  if (!res.ok) throw new Error('failed to fetch article');
  return res.json();
}

export async function createArticle(body) {
  const res = await fetch(`${API_BASE}/articles`, {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(body)
  });
  return res.json();
}

export async function updateArticle(id, body) {
  const res = await fetch(`${API_BASE}/articles/${id}`, {
    method: 'PUT',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(body)
  });
  return res.json();
}

export async function deleteArticle(id) {
  const res = await fetch(`${API_BASE}/articles/${id}`, { method:'DELETE' });
  return res.json();
}
