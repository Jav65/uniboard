import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom'; 
import {useNavigate} from 'react-router-dom'

interface Comment {
  id: string;
  thread_id: string;
  content: string;
  author_id: string;
  author_name: string;
  created_at: string;
}

interface Thread {
  id: string;
  title: string;
  author_name: string;
  content: string;
  tags_name: string[];
  likes: number;
  views: number;
  comments: number;
  created_at: string;
}


export default function ThreadView() {
  const { id } = useParams<{ id : string}>();
  const [thread, setThread] = useState<Thread | null>(); 
  const [comments, setComments] = useState<Comment[]>([]);
  const [commentContent, setCommentContent] = useState("");
  const [author_id, setAuthorId] = useState(0);
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

  useEffect(() => {
    const fetchThread = async () => {
      const res = await fetch(`https://uniboard-1.onrender.com/api/thread/${id}`)
      const body = await res.json()
      setThread(body.data)
  } 
  fetchThread()
  }, [])

  useEffect(() => {
    const fetchComments = async () => {
      const res = await fetch(`https://uniboard-1.onrender.com/api/comments/${id}`)
      const body = await res.json()
      setComments(body.data)
      console.log(body.data)
  }
    fetchComments()
  }, [])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    await fetch(`https://uniboard-1.onrender.com/api/comment/${id}`, {
      method: 'POST',
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({thread_id: Number(id), content: commentContent, author_id,})
    });
    console.log(JSON.stringify({thread_id: id, content: commentContent, author_id,}))

    alert('Comment created successfully!');
    window.location.reload()
  };

  if (!thread) {
    return <p>Thread not found.</p>;
  }

  return (
    <div className="max-w-3xl mx-auto">
      <h1 className="text-3xl font-bold mb-4">{thread.title}</h1>
      <div className="flex items-center space-x-4 mb-6">
        {/* Image */}
        <div>
          <p className="text-gray-300 py-3">{thread.author_name}</p>
          <div className="flex space-x-2">
            {thread.tags_name.map((tag: string) => (
              <span key={tag} className="bg-blue-600 text-white text-xs px-2 py-1 rounded-full">
                {tag}
              </span>
            ))}
          </div>
        </div>
      </div>
      <div className="bg-gray-800 rounded-lg p-6 mb-8">
        <p className="text-gray-100">{thread.content}</p>
      </div>
      <h2 className="text-2xl font-bold mb-4">Comments</h2>
      <div className="space-y-4 mb-8">
        {comments.map((comment: any) => (
          <div key={comment.id} className="bg-gray-800 rounded-lg p-4">
            <div className="flex items-start space-x-4">
              {/* //Image */}
              <div className="flex-grow">
                <p className="text-gray-300 font-semibold">{comment.author_name}</p>
                <p className="text-gray-100 mt-1">{comment.content}</p>
              </div>
              {/* <button className="flex items-center text-gray-400 hover:text-blue-500 transition-colors">
                <ThumbsUp size={16} className="mr-1" />
                <span>{comment.upvotes}</span>
              </button> */}
            </div>  
          </div>
        ))}
      </div>
      <form className="space-y-4" onSubmit={handleSubmit}>
        <textarea
          placeholder="Add your comment..."
          rows={4}
          value={commentContent}
          onChange={(e) => setCommentContent(e.target.value)}
          className="w-full bg-gray-700 text-white rounded-lg py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
        ></textarea>
        <button
          type="submit"
          className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg transition-colors"
        >
          Post Comment
        </button>
      </form>
    </div>
  )
}

