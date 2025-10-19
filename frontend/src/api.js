import axios from 'axios';

const API_BASE = "http://localhost:8080/article";

// Get all posts with pagination and status filter
export const fetchArticles = async (limit = 10, offset = 0, status = '') => {
  const params = { limit, offset };
  if (status) params.status = status;
  
  const response = await axios.get(`${API_BASE}`, { params });
  return response.data;
};

// Get single post by ID
export const fetchArticle = async (id) => {
  const response = await axios.get(`${API_BASE}/${id}`);
  return response.data;
};

// Create new post
export const createArticle = async (postData) => {
  const response = await axios.post(`${API_BASE}/`, postData);
  return response.data;
};

// Update post
export const updateArticle = async (id, postData) => {
  const response = await axios.put(`${API_BASE}/${id}`, postData);
  return response.data;
};

// Delete post
export const deleteArticle = async (id) => {
  const response = await axios.delete(`${API_BASE}/${id}`);
  return response.data;
};

// Export as default object too (for compatibility)
export default {
  fetchArticles,
  fetchArticle,
  createArticle,
  updateArticle,
  deleteArticle
};