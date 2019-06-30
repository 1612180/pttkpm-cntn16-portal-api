# Awesome Portal

API endpoints
```
GET     /api/students
GET     /api/students/:id
POST    /api/students
DELETE  /api/students/:mssv

POST    /api/auth/login

GET     /api/programs
POST    /api/programs

GET     /api/faculties
POST    /api/faculties

GET     /api/subjects
GET     /api/subjects/:id
POST    /api/subjects
DELETE  /api/subjects/:id

GET     /api/subject_types
POST    /api/subject_types
```

Sample `.env`
```
PORT=8080
GIN_MODE=debug
DATABASE_URL=postgres://postgres@localhost:54320/postgres?sslmode=disable
```