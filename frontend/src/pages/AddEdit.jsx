import React, {useEffect, useState} from 'react';
import { createArticle, fetchArticle, updateArticle } from '../api';
import { useNavigate, useParams } from 'react-router-dom';

export default function AddEdit(){
  const { id } = useParams();
  const navigate = useNavigate();
  const [form, setForm] = useState({title:'', content:'', category:'', status:'draft'});
  const [loading,setLoading] = useState(false);

  useEffect(()=>{
    if (id) {
      setLoading(true);
      fetchArticle(id).then(p => setForm({
        title: p.title,
        content: p.content,
        category: p.category,
        status: p.status
      })).catch(e=>alert('failed')).finally(()=>setLoading(false));
    }
  },[id]);

  async function onSubmit(statusType) {
    const payload = {...form, status: statusType};
    try {
      if (id) {
        await updateArticle(id, payload);
        alert('updated');
      } else {
        await createArticle(payload);
        alert('created');
      }
      navigate('/');
    } catch(e){
      console.error(e);
      alert('failed: '+(e.message||e));
    }
  }

  return (
    <div className='card'>
      <h3>{id ? 'Edit Article' : 'Add New'}</h3>
      {loading ? <div>Loading...</div> : (
        <div style={{maxWidth:800}}>
          <div style={{marginBottom:8}}>
            <label>Title</label><br/>
            <input className='form-control' value={form.title} onChange={e=>setForm({...form,title:e.target.value})} />
          </div>
          <div style={{marginBottom:8}}>
            <label>Category</label><br/>
            <input className='form-control' value={form.category} onChange={e=>setForm({...form,category:e.target.value})} />
          </div>
          <div style={{marginBottom:8}}>
            <label>Content</label><br/>
            <textarea className='form-control' value={form.content} onChange={e=>setForm({...form,content:e.target.value})} rows={12} />
          </div>
          <div style={{marginTop:8}}>
            <button onClick={()=>onSubmit('publish')} className='btn primary' style={{marginRight:8}}>Publish</button>
            <button onClick={()=>onSubmit('draft')} className='btn'>Draft</button>
          </div>
        </div>
      )}
    </div>
  );
}
