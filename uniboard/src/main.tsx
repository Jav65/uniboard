import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css'
import Layout from './Layout.tsx'
import { BrowserRouter } from "react-router-dom";

ReactDOM.createRoot(document.getElementById('root')!).render(
  // <BrowserRouter>
  <React.StrictMode>
    <Layout/>
  </React.StrictMode>,
  // {/* </BrowserRouter>, */}
);



