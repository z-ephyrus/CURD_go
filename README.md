Go CRUD Example (Echo + Bun + Postgres)

A minimal CRUD API written in Go using:

Echo – lightweight web framework

Bun – ORM for PostgreSQL

PostgreSQL as database

Context & goroutines for timeout/cancellation handling

This is a test project to learn backend structuring, context usage, and clean CRUD code.

📂 Project Structure

/test-crud
├── main.go
├── go.mod
├── .env
└── internal/
├── db/
│ └── db.go
└── user/
├── model.go
└── handler.go

⚡ Setup
1. Clone the repo

git clone https://github.com/
<your-username>/test-crud.git
cd test-crud

2. Install dependencies

go mod tidy

3. Configure .env

Create a .env file at the project root:

DB_DSN=postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
APP_PORT=8080

4. Run PostgreSQL (example using Docker)

docker run --name pg-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testdb -p 5432:5432 -d postgres

5. Run the server

go run main.go

Server will start at http://localhost:8080

🔧 API Endpoints (with curl examples)
Create User

curl -X POST http://localhost:8080/users
 \
-H "Content-Type: application/json" \
-d '{"name":"Alice","email":"alice@mail.com
"}'

Get Users

curl http://localhost:8080/users

Update User

curl -X PUT http://localhost:8080/users/1
 \
-H "Content-Type: application/json" \
-d '{"name":"Alice B","email":"aliceb@mail.com
"}'

Delete User

curl -X DELETE http://localhost:8080/users/1

🧠 Learnings

How to use Context with timeouts in database queries

How to structure Go projects with internal/ packages

How Bun ORM maps models to SQL queries

How Echo simplifies route handling

🚀 Future Improvements

Add JWT authentication

Docker Compose for full stack (Go + Postgres)

CI/CD pipeline

Unit tests with testcontainers-go
