import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css'
import Layout from './Layout.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  // <BrowserRouter>
  <React.StrictMode>
    <Layout/>
  </React.StrictMode>,
  // {/* </BrowserRouter>, */}
);



