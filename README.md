# uniboard
Uniboard is a university application web forum with a modern and user-friendly dark mode design

**Setup Instructions**
1. cd backend (navigate into backend dir)
2. go run main.go

3. cd uniboard (navigate into frontend dir)
4. npm install
5. npm run dev (make sure it run on localhost:5173)

**To run on local**
1. Create a database: database='uniboard', host='localhost', port='5432', user='postgres', password='postgres'
2. Uncomment local dbURL in 'main.go'
4. change all api fetching in uniboard dir from 'https://uniboard-1.onrender.com/' to 'https://localhost:8080/'
