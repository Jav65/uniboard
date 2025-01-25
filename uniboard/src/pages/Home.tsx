import { useState, useEffect } from "react"
import { Plus, X } from "lucide-react"
import ThreadCard from "../components/ThreadCard"

interface Thread {
  id: string;
  title: string;
  author_name: string;
  author_id: number;
  content: string;
  tags_name: string[];
  likes: number;
  views: number;
  comments: number;
  created_at: string;
}


const tabs = ["Trending", "Your Posts"]
const sortOptions = ["Most Recent","Most Comments", "Most Views", "Most Likes"]


export default function Home({ searchQuery }: { searchQuery: string }) {
  const [activeTab, setActiveTab] = useState(tabs[0])
  const [sortBy, setSortBy] = useState(sortOptions[0])
  const [currentId, setCurrentId] = useState(null);
  const [threads, setThreads] = useState<Thread[]>([])
  const [copyThreads, setCopyThreads] = useState<Thread[]>([])
  const [tags, setTags] = useState<string[]>([])
  const [isFilterOpen, setIsFilterOpen] = useState(false)
  const [selectedUniversities, setSelectedUniversities] = useState<string[]>([])
  const [selectedProcesses, setSelectedProcesses] = useState<string[]>([])
  const [selectedTags, setSelectedTags] = useState<string[]>([])

  useEffect(() => {
    const fetchCurrentUser = async () => {
      try {
        const res = await fetch('https://uniboard-1.onrender.com//api/user', {credentials: "include",})
        if (!res.ok) {
          throw new Error("User not logged in");
        }
        const user = await res.json();
        setCurrentId(user.id);
      } catch (err) {
        console.error("Error fetching user:", err);
      }
    };

    fetchCurrentUser();
  }, []);


  useEffect(() => {
    const fetchThreads = async () => {
      const res = await fetch(`https://uniboard-1.onrender.com//api/threads?sortBy=${sortBy}&search=${searchQuery}`)
      const body = await res.json()
      setThreads(body.data)
      setCopyThreads(body.data)
      console.log(body.data)
  }
    fetchThreads()
  }, [sortBy, searchQuery])

  useEffect(() => {
    const fetchTags = async () => {
      const res = await fetch('https://uniboard-1.onrender.com//api/tags/name')
      const body = await res.json()
      setTags(body.data)
      console.log(body.data)
  }
    fetchTags()
  }, [])
  

  useEffect(() => {
    let filteredThreads = copyThreads

    if (activeTab === "Your Posts") {
      filteredThreads = filteredThreads.filter((thread) => thread.author_id === currentId)
    }

    // if (selectedUniversities.length > 0) {
    //   filteredThreads = filteredThreads.filter((thread) => selectedUniversities.includes(thread.university))
    // }

    // if (selectedProcesses.length > 0) {
    //   filteredThreads = filteredThreads.filter((thread) => selectedProcesses.includes(thread.applicationProcess))
    // }

    if (selectedTags.length > 0) {
      filteredThreads = filteredThreads.filter((thread) => selectedTags.every((tag) => thread.tags_name.includes(tag)));
    }

    setThreads(filteredThreads)
  }, [activeTab, sortBy, selectedUniversities, selectedProcesses, selectedTags])

  const toggleFilter = (type: "university" | "process" | "tag", value: string) => {
    switch (type) {
      case "university":
        setSelectedUniversities((prev) => (prev.includes(value) ? prev.filter((u) => u !== value) : [...prev, value]))
        break
      case "process":
        setSelectedProcesses((prev) => (prev.includes(value) ? prev.filter((p) => p !== value) : [...prev, value]))
        break
      case "tag":
        setSelectedTags((prev) => (prev.includes(value) ? prev.filter((t) => t !== value) : [...prev, value]))
        break
    }
  }

  const resetFilters = () => {
    setSelectedUniversities([])
    setSelectedProcesses([])
    setSelectedTags([])
  }

  return (
    <div className="relative">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">{activeTab}</h1>
        <a
          href="/new-thread"
          className="bg-blue-600 hover:bg-blue-700 text-white rounded-full p-2 transition-colors"
        >
          <Plus size={24} />
        </a>
      </div>

      <div className="flex space-x-4 mb-6">
        {tabs.map((tab) => (
          <button
            key={tab}
            className={`px-4 py-2 rounded-full transition-colors ${
              activeTab === tab ? "bg-blue-600 text-white" : "bg-gray-700 text-gray-300 hover:bg-gray-600"
            }`}
            onClick={() => setActiveTab(tab)}
          >
            {tab}
          </button>
        ))}
      </div>

      <div className="flex justify-between items-center mb-4">
        <div className="relative">
          <select
            value={sortBy}
            onChange={(e) => setSortBy(e.target.value)}
            className="appearance-none bg-gray-700 text-white rounded-lg py-2 px-4 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            {sortOptions.map((option) => (
              <option key={option} value={option}>
                {option}
              </option>
            ))}
          </select>
        </div>
        <button
          onClick={() => setIsFilterOpen(true)}
          className="bg-blue-600 text-white rounded-lg py-2 px-4 hover:bg-blue-700 transition-colors"
        >
          Filter
        </button>
      </div>

      {isFilterOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-gray-800 p-6 rounded-lg w-full max-w-md">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">Filters</h2>
              <button onClick={() => setIsFilterOpen(false)} className="text-gray-400 hover:text-white">
                <X size={24} />
              </button>
            </div>
            {/* <div className="mb-4">
              <h3 className="text-lg font-semibold mb-2">Universities</h3>
              <div className="flex flex-wrap gap-2">
                {universities.map((uni) => (
                  <button
                    key={uni}
                    onClick={() => toggleFilter("university", uni)}
                    className={`px-3 py-1 rounded-full text-sm ${
                      selectedUniversities.includes(uni)
                        ? "bg-blue-600 text-white"
                        : "bg-gray-700 text-gray-300 hover:bg-gray-600"
                    }`}
                  >
                    {uni}
                  </button>
                ))}
              </div>
            </div> */}
            {/* <div className="mb-4">
              <h3 className="text-lg font-semibold mb-2">Application Processes</h3>
              <div className="flex flex-wrap gap-2">
                {applicationProcesses.map((process) => (
                  <button
                    key={process}
                    onClick={() => toggleFilter("process", process)}
                    className={`px-3 py-1 rounded-full text-sm ${
                      selectedProcesses.includes(process)
                        ? "bg-blue-600 text-white"
                        : "bg-gray-700 text-gray-300 hover:bg-gray-600"
                    }`}
                  >
                    {process}
                  </button>
                ))}
              </div>
            </div> */}
            <div className="mb-4">
              <h3 className="text-lg font-semibold mb-2">Tags</h3>
              <div className="flex flex-wrap gap-2">
                {tags.map((tag) => (
                  <button
                    key={tag}
                    onClick={() => toggleFilter("tag", tag)}
                    className={`px-3 py-1 rounded-full text-sm ${
                      selectedTags.includes(tag)
                        ? "bg-blue-600 text-white"
                        : "bg-gray-700 text-gray-300 hover:bg-gray-600"
                    }`}
                  >
                    {tag}
                  </button>
                ))}
              </div>
            </div>
            <div className="flex justify-between">
              <button
                onClick={resetFilters}
                className="bg-gray-700 text-white rounded-lg py-2 px-4 hover:bg-gray-600 transition-colors"
              >
                Reset
              </button>
              <button
                onClick={() => setIsFilterOpen(false)}
                className="bg-blue-600 text-white rounded-lg py-2 px-4 hover:bg-blue-700 transition-colors"
              >
                Apply Filters
              </button>
            </div>
          </div>
        </div>
      )}

      <div className="space-y-4">
        {threads.length > 0 ? (
          threads.map((thread) => <ThreadCard key={thread.id} {...thread} />)
        ) : (
          <p>No threads available.</p>
        )}
      </div>
    </div>
  )
}

