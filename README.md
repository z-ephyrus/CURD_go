Go CRUD Example (Echo + Bun + Postgres)A minimal, robust CRUD (Create, Read, Update, Delete) API written in Go, focusing on modern backend best practices, including proper context handling and clean project structure.🌟 Key TechnologiesThis project utilizes a powerful and lightweight stack:Echo: A high-performance, minimalist Go web framework.Bun: A simple, fast, and feature-rich ORM (Object-Relational Mapper) for PostgreSQL.PostgreSQL: The robust relational database used for persistence.context & goroutines: Used extensively for managing operation timeouts and cancellation signals within the API and database layer.📂 Project StructureThis structure follows common Go practices, utilizing the internal/ directory to encapsulate private code:/test-crud
├── main.go               # Entry point, framework setup, and routing
├── go.mod                # Go module file
├── .env                  # Environment variables for configuration
└── internal/
    ├── db/
    │    └── db.go         # Database initialization and connection logic (Bun setup)
    └── user/
        ├── model.go      # Defines the User struct and Bun database table
        └── handler.go    # HTTP handlers (CRUD logic)
⚡ Setup1. Clone the repogit clone [https://github.com/](https://github.com/)<your-username>/test-crud.git
cd test-crud
2. Install dependenciesgo mod tidy
3. Configure .envCreate a file named .env in the project root with your database connection string and application port:DB_DSN=postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
APP_PORT=8080
4. Run PostgreSQL (Example using Docker)Use the following command to spin up a local PostgreSQL instance that matches the .env credentials:docker run --name pg-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testdb -p 5432:5432 -d postgres
5. Run the servergo run main.go
The server will start at http://localhost:8080.🔧 API Endpoints (with curl examples)All API endpoints are prefixed with /users.MethodEndpointDescriptionPOST/usersCreates a new user.GET/usersRetrieves all users.PUT/users/:idUpdates an existing user by ID.DELETE/users/:idDeletes a user by ID.Create Usercurl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"Alice","email":"alice@mail.com"}'
Get Userscurl http://localhost:8080/users
Update Usercurl -X PUT http://localhost:8080/users/1 \
-H "Content-Type: application/json" \
-d '{"name":"Alice B","email":"aliceb@mail.com"}'
Delete Usercurl -X DELETE http://localhost:8080/users/1
🧠 LearningsThis project served as a practical exploration of several key concepts in Go and modern backend development:Context Management: How to use the context package effectively with timeouts (context.WithTimeout) in database queries (db.QueryRowContext) to prevent resource leaks and handle slow operations gracefully.Go Project Structure: Implementing the recommended directory structure using internal/ packages (db/, user/) to enforce separation of concerns and maintain a clean codebase.Bun ORM: Understanding how the Bun ORM simplifies PostgreSQL interactions, automatically mapping Go structs to SQL table definitions and queries.Echo Framework: Utilizing Echo for simple, fast routing, parameter binding, and response handling.🚀 Future ImprovementsAdd JWT authentication for secure access to user endpoints.Implement a full-stack development environment using Docker Compose (Go API + Postgres).Set up a CI/CD pipeline for automated testing and deployment.Write comprehensive Unit tests for the handlers and data layer, potentially utilizing testcontainers-go.
