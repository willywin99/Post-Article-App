import React, {useEffect, useState} from 'react';
import { fetchArticles } from '../api';

export default function Preview(){
  const [posts,setPosts] = useState([]);
  const [limit] = useState(5);
  const [offset, setOffset] = useState(0);

  async function load(){
    try {
      const data = await fetchArticles(limit, offset, 'publish');
      setPosts(data);
    } catch(e){
      alert('failed');
    }
  }
  useEffect(()=>{ load(); }, [offset]);

  return (
    <div>
      <h3>Preview (Published)</h3>
      {posts.length === 0 && <div>No published posts</div>}
      {posts.map(p => (
        <div key={p.id} className='card' style={{marginBottom:10}}>
          <h4>{p.title}</h4>
          <p><em>{p.category}</em></p>
          <div>{p.content.slice(0,1000)}</div>
        </div>
      ))}
      <div>
        <button onClick={()=>setOffset(Math.max(0, offset - limit))} disabled={offset===0} className='btn'>Prev</button>
        <button onClick={()=>setOffset(offset + limit)} className='btn' style={{marginLeft:8}}>Next</button>
      </div>
    </div>
  );
}
