import { Menu, Search, LogIn, UserPlus, LogOut } from 'lucide-react'
import { useState, useEffect } from 'react'

export default function Header({ searchQuery, setSearchQuery }: { searchQuery: string; setSearchQuery: (query: string) => void }) {
  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      setSearchQuery(event.currentTarget.value);
    }
  };
  const [user, setUser] = useState<any>(null)

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('https://uniboard-1.onrender.com//api/user', {credentials: "include",})
        if(res.ok){
          const body = await res.json()
          setUser(body)
          console.log(body)
        } 
      } catch (error) {
        console.error('Error fetching user:', error)
      }
    }
    fetchUser()
  }, [])

  const handleLogout = async () => {
    const res = await fetch("https://uniboard-1.onrender.com//api/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    
    if (res.ok) {
      const data = await res.json()
      console.log(data)
      setUser(null);
    } else {
      alert("Logout failed. Please try again.");
    }
  }

  return (
    <header className="bg-gray-800 py-4">
      <div className="container mx-auto px-4 flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <button className="text-gray-300 hover:text-white transition-colors">
            <Menu size={24} />
          </button>
          <a href="/" className="text-2xl font-bold text-white">Uniboard</a>
        </div>
        <div className="flex-grow max-w-xl mx-4">
          <div className="relative">
            <input
              type="text"
              placeholder="Search Uniboard"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              onKeyDown={handleKeyDown}
              className="w-full bg-gray-700 text-white rounded-full py-2 px-4 pl-10 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <Search className="absolute left-3 top-2.5 text-gray-400" size={20} />
          </div>
        </div>
        <div className="flex items-center space-x-4">
          {user == null ?  (
            <>
              <a href="/login" className="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded-full transition-colors flex items-center">
                <LogIn size={18} className="mr-2" />
                Log in
              </a>
              <a href="/register" className="bg-green-600 hover:bg-green-700 text-white py-2 px-4 rounded-full transition-colors flex items-center">
                <UserPlus size={18} className="mr-2" />
                Register
              </a>
            </>
          ) : (
            <>
              <span className="text-white">Welcome, {user.username}</span>
              <button onClick={handleLogout} className="bg-red-600 hover:bg-red-700 text-white py-2 px-4 rounded-full transition-colors flex items-center">
                <LogOut size={18} className="mr-2" />
                Log out
              </button>
            </ >
          ) 
        }
        </div>
      </div>
    </header>
  )
}

