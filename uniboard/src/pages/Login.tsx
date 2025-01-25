import { useState } from 'react'
import { useNavigate } from "react-router-dom";

export default function Login() {
  const navigate = useNavigate()
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const res = await fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({ email, password }),
    });
    console.log(JSON.stringify({ email, password }));

    
    if (res.ok) {
      const data = await res.json()
      console.log(data)
      const _ = await fetch('http://localhost:8080/api/user', {credentials: "include",})
      navigate('/')
      window.location.reload()
    } else {
      alert("Invalid email or password");
    }
  
  };

  return (
    <div className="max-w-md mx-auto">
      <h1 className="text-3xl font-bold mb-6 text-center">Login to Uniboard</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="email" className="block text-sm font-medium text-gray-300 mb-1">
            Email
          </label>
          <input
            type="text"
            id="email"
            required
            name="email"
            onChange={e => setEmail(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label htmlFor="password" className="block text-sm font-medium text-gray-300 mb-1">
            Password
          </label>
          <input
            type="password"
            id="password"
            required
            name="password"
            onChange={e => setPassword(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg transition-colors"
        >
          Log In
        </button>
      </form>
      <p className="mt-4 text-center text-gray-400">
        Don't have an account?{' '}
        <a href="/register" className="text-blue-400 hover:underline">
          Register here
        </a>
      </p>
    </div>
  )
}

