# Go CRUD Example (Echo + Bun + Postgres)

A minimal CRUD API written in **Go** using:

- [Echo](https://echo.labstack.com/) – lightweight web framework
- [Bun](https://bun.uptrace.dev/) – ORM for PostgreSQL
- PostgreSQL as database
- Context & goroutines for timeout/cancellation handling

This is a **test project** to learn backend structuring, context usage, and clean CRUD code.

---

## 📂 Project Structure
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

yaml
Copy code

---

## ⚡ Setup

### 1. Clone the repo
```bash
git clone https://github.com/<your-username>/test-crud.git
cd test-crud
2. Install dependencies
bash
Copy code
go mod tidy
3. Configure .env
Create a .env file at the project root:

env
Copy code
DB_DSN=postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
APP_PORT=8080
4. Run PostgreSQL (example using Docker)
bash
Copy code
docker run --name pg-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testdb -p 5432:5432 -d postgres
5. Run the server
bash
Copy code
go run main.go
Server will start at http://localhost:8080

🔧 API Endpoints (with curl examples)
Create User
bash
Copy code
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@mail.com"}'
Get Users
bash
Copy code
curl http://localhost:8080/users
Update User
bash
Copy code
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice B","email":"aliceb@mail.com"}'
Delete User
bash
Copy code
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
