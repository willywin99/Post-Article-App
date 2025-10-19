import React, {useEffect, useState} from 'react';
import { fetchArticles, updateArticle } from '../api';
import { useNavigate } from 'react-router-dom';

const TABS = ['publish','draft','thrash'];

export default function AllPosts(){
  const [posts, setPosts] = useState([]);
  const [tab, setTab] = useState('publish');
  const [limit] = useState(20);
  const [offset,setOffset] = useState(0);
  const [loading,setLoading] = useState(false);
  const navigate = useNavigate();

  async function load(){
    setLoading(true);
    try {
      const data = await fetchArticles(limit, offset, tab);
      setPosts(data);
    } catch(e){
      console.error(e);
      alert('Failed loading posts');
    } finally { setLoading(false); }
  }

  useEffect(()=>{ load(); }, [offset, tab]);

  function onEdit(id){
    navigate(`/edit/${id}`);
  }

  async function onTrash(id) {
    try {
      const existing = posts.find(p => p.id === id);
      if (!existing) { alert('not found'); return; }
      const payload = {
        title: existing.title,
        content: existing.content,
        category: existing.category,
        status: 'thrash'
      };
      await updateArticle(id, payload);
      await load();
    } catch(e){
      console.error(e);
      alert('failed to trash');
    }
  }

  return (
    <div className='card'>
      <h3>All Posts</h3>
      <div style={{marginBottom:10}}>
        {TABS.map(t => (
          <button key={t} onClick={()=>setTab(t)} className={'btn '+ (t===tab?'primary':'')} style={{marginRight:8}}>
            {t[0].toUpperCase()+t.slice(1)}
          </button>
        ))}
      </div>

      {loading ? <div>Loading...</div> :
        <table className='table'>
          <thead>
            <tr><th>Title</th><th>Category</th><th>Action</th></tr>
          </thead>
          <tbody>
            {posts.length === 0 && <tr><td colSpan='3'>No posts</td></tr>}
            {posts.map(p => (
              <tr key={p.id}>
                <td>{p.title}</td>
                <td>{p.category}</td>
                <td>
                  <button onClick={()=>onEdit(p.id)} className='btn' style={{marginRight:8}}>✏️</button>
                  <button onClick={()=>onTrash(p.id)} className='btn'>🗑️</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      }

      <div style={{marginTop:12}}>
        <button onClick={()=>setOffset(Math.max(0, offset - limit))} disabled={offset===0} className='btn'>Prev</button>
        <span style={{margin: '0 12px'}}>Page offset: {offset}</span>
        <button onClick={()=>setOffset(offset + limit)} className='btn'>Next</button>
      </div>
    </div>
  );
}
