import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import DashboardLayout from './layouts/DashboardLayout';
import AllPosts from './pages/AllPosts';
import AddEdit from './pages/AddEdit';
import Preview from './pages/Preview';

export default function App(){
  return (
    <BrowserRouter>
      <DashboardLayout>
        <Routes>
          <Route path='/' element={<AllPosts />} />
          <Route path='/add' element={<AddEdit />} />
          <Route path='/edit/:id' element={<AddEdit />} />
          <Route path='/preview' element={<Preview />} />
        </Routes>
      </DashboardLayout>
    </BrowserRouter>
  );
}
