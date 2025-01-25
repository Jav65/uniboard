import { MessageSquare, Eye, ThumbsUp } from 'lucide-react'

interface ThreadCardProps {
  id: number
  title: string
  author_name : string
  comments: number
  views: number
  likes: number
  // university: string
  // applicationProcess: string
  tags_name: string[]
  created_at : string | Date
}

 

export default function ThreadCard({ id, title, author_name, comments, views, likes, tags_name }: ThreadCardProps) {
  console.log(comments, views, likes)
  return (
    <a href={`/thread/${id}`} className="block">
      <div className="bg-gray-800 rounded-lg p-4 hover:bg-gray-700 transition-colors">
        <div className="flex items-start space-x-4">
          {/* Image */}
          <div className="flex-grow">
            <h3 className="text-lg font-semibold text-white mb-1">{title}</h3>
            <p className="text-gray-400 text-sm mb-2">{author_name}</p>
            <div className="flex flex-wrap gap-2 mb-2">
              {/* <span className="bg-blue-600 text-white text-xs px-2 py-1 rounded-full">{university}</span>
              <span className="bg-green-600 text-white text-xs px-2 py-1 rounded-full">{applicationProcess}</span> */}
              {tags_name.map((tag) => (
                <span key={tag} className="bg-gray-600 text-white text-xs px-2 py-1 rounded-full">
                  {tag}
                </span>
              ))}
            </div>
          </div>
          <div className="flex items-center space-x-4 text-gray-400">
            <div className="flex items-center">
              <MessageSquare size={16} className="mr-1" />
              <span className="text-sm">{comments}</span>
            </div>
            <div className="flex items-center">
              <Eye size={16} className="mr-1" />
              <span className="text-sm">{views}</span>
            </div>
            <div className="flex items-center">
              <ThumbsUp size={16} className="mr-1" />
              <span className="text-sm">{likes}</span>
            </div>
          </div>
        </div>
      </div>
    </a>
  )
}

