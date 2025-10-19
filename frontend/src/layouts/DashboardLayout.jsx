import React from 'react'
import { Link, useLocation } from 'react-router-dom'

export default function DashboardLayout({ children }) {
  const loc = useLocation();
  const active = (path) => loc.pathname === path ? 'active' : '';

  return (
    <div className='wrapper'>
      <nav className='sidebar'>
        <div className='sidebar-header'>
          <h3 className='text-white text-center py-3'>Sharing Vision</h3>
        </div>
        <ul className='components'>
          <li><Link className={active('/')} to='/'>All Posts</Link></li>
          <li><Link className={active('/add')} to='/add'>Add New</Link></li>
          <li><Link className={active('/preview')} to='/preview'>Preview</Link></li>
        </ul>
      </nav>

      <div className='page-content'>
        <header className='header'>
          {/* no branding per request */}
        </header>
        <main className='main'>
          {children}
        </main>
        <footer className='footer'>Sharing Vision - Demo</footer>
      </div>
    </div>
  )
}
