import './index.css'
import Header from "./layouts/Header";
import Footer from "./layouts/Footer";
import { useState } from "react";
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Home from './pages/Home';
import NewThread from './pages/NewThread';
import ThreadView from './pages/ThreadView';
import Login from './pages/Login';
import Register from './pages/Register';


export default function Layout() {
  const [searchQuery, setSearchQuery] = useState("");
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Home searchQuery={searchQuery}/>,
    },
    {
      path: '/login',
      element: <Login />,
    },
    {
      path: '/register',
      element: <Register />,
    },
    {
      path: '/new-thread',
      element: <NewThread />,
    },
    {
      path: '/thread/:id',
      element: <ThreadView />,
    },
  ]);

  return (
    <html lang="en" className="dark">
      <body className={`bg-gray-900 text-gray-100 min-h-screen flex flex-col`}>
        <Header searchQuery={searchQuery} setSearchQuery={setSearchQuery}/>
        <main className="flex-grow container mx-auto px-4 py-8">
          <RouterProvider router={router} />
        </main>
        <Footer />
      </body>
    </html>
  )
}




