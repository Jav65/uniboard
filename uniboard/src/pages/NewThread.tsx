const universities = ["MIT", "Stanford", "UCLA", "Harvard", "UC Berkeley"]
const applicationProcesses = ["Undergraduate", "Graduate", "Transfer"]


import { useState, useEffect } from "react"
import { Paperclip } from "lucide-react"
import {useNavigate} from 'react-router-dom'

export default function NewThread() {
  const [isAnonymous, setIsAnonymous] = useState(false)
  const [author_id, setAuthorId] = useState(0);
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const [tags, setTags] = useState<any>([])
  const [selectedTags, setSelectedTags] = useState<string[]>([])
  const [selectedUniversity, setSelectedUniversity] = useState("")
  const [selectedProcess, setSelectedProcess] = useState("")
  const navigate = useNavigate()

  useEffect(() => {
    const fetchCurrentUser = async () => {
      try {
        // const res = await fetch('https://uniboard-1.onrender.com/api/user', {credentials: "include",})
        // if (!res.ok) {
        //   throw new Error("User not logged in");
        // }
        // const user = await res.json();
        // setAuthorId(user.id);
        const user = {username: "john", id: 2, email: "john@gmail.com"} //mock data, cookies setting error in vercel
        setAuthorId(user.id);
      } catch (err) {
        console.error("Error fetching user:", err);
        navigate("/login"); 
      }
    };

    fetchCurrentUser();
  }, [navigate]);
  
  const handleSubmit = async (e: React.FormEvent) => {
      e.preventDefault()
      
      const response = await fetch('https://uniboard-1.onrender.com/api/thread', {
        method: 'POST',
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({title, content, author_id, tags: selectedTags, })
      });
  
      const body = await response.json();
      console.log(body)
      alert('Thread created successfully!');
  
      navigate('/')
    };

  useEffect(() => {
      const fetchTags = async () => {
        const res = await fetch('https://uniboard-1.onrender.com/api/tags')
        const body = await res.json()
        setTags(body.data)
        console.log(body.data)
    }
      fetchTags()
    }, [])

  const handleTagToggle = (tag: string) => {
    setSelectedTags((prev) => (prev.includes(tag) ? prev.filter((t) => t !== tag) : [...prev, tag]))
  }

  

  return (
    <div className="max-w-2xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Create New Thread</h1>
      <form className="space-y-6" onSubmit={handleSubmit}>
        <div>
          <label htmlFor="title" className="block text-sm font-medium text-gray-300 mb-1">
            Title (required)
          </label>
          <input
            type="text"
            id="title"
            required
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label htmlFor="content" className="block text-sm font-medium text-gray-300 mb-1">
            Content (required)
          </label>
          <textarea
            id="content"
            required
            rows={6}
            value={content}
            onChange={(e) => setContent(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          ></textarea>
        </div>
        <div>
          <label htmlFor="university" className="block text-sm font-medium text-gray-300 mb-1">
            University (required)
          </label>
          <select
            id="university"
            required
            value={selectedUniversity}
            onChange={(e) => setSelectedUniversity(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">Select a university</option>
            {universities.map((uni) => (
              <option key={uni} value={uni}>
                {uni}
              </option>
            ))}
          </select>
        </div>
        <div>
          <label htmlFor="process" className="block text-sm font-medium text-gray-300 mb-1">
            Application Process (required)
          </label>
          <select
            id="process"
            required
            value={selectedProcess}
            onChange={(e) => setSelectedProcess(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">Select an application process</option>
            {applicationProcesses.map((process) => (
              <option key={process} value={process}>
                {process}
              </option>
            ))}
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-300 mb-1">Tags (required)</label>
          <div className="flex flex-wrap gap-2">
            {tags.map((tag: any) => (
              <button
                key={tag.id}
                type="button"
                onClick={() => handleTagToggle(tag.id)}
                className={`px-3 py-1 rounded-full text-sm ${
                  selectedTags.includes(tag.id) ? "bg-blue-600 text-white" : "bg-gray-700 text-gray-300 hover:bg-gray-600"
                }`}
              >
                {tag.name}
              </button>
            ))}
          </div>
        </div>
        <div>
          <label htmlFor="file" className="block text-sm font-medium text-gray-300 mb-1">
            Attach File (optional)
          </label>
          <div className="relative">
            <input type="file" id="file" className="hidden" />
            <label
              htmlFor="file"
              className="flex items-center justify-center w-full bg-gray-700 text-white rounded-lg py-2 px-4 cursor-pointer hover:bg-gray-600 transition-colors"
            >
              <Paperclip size={20} className="mr-2" />
              Choose file
            </label>
          </div>
        </div>
        <div className="flex items-center">
          <input
            type="checkbox"
            id="anonymous"
            checked={isAnonymous}
            onChange={(e) => setIsAnonymous(e.target.checked)}
            className="mr-2"
          />
          <label htmlFor="anonymous" className="text-sm text-gray-300">
            Post anonymously
          </label>
        </div>
        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg transition-colors"
        >
          Submit Thread
        </button>
      </form>
    </div>
  )
}



