# CRUD API with Go (Fiber) and PostgreSQL

## About the Project
This project is a simple implementation of a CRUD (Create, Read, Update, Delete) API built using Go with the Fiber framework and a PostgreSQL database. It provides a basic API for managing "TODO" items, which can be used in a frontend application such as a ToDo app or extended for other use cases.

## Why You Can Use It
- **For a Frontend Todo App:** You can build a simple frontend Todo app that interacts with this API for adding, retrieving, updating, and deleting tasks.
- **Customization:** You can further develop this API to meet specific needs, adding more complex functionalities or integrating it into a bigger project.

## How You Can Use It

### Dependencies:
You can install the required dependencies by using `go.mod`:
1. Install Go: [Go installation guide](https://go.dev/doc/install).
2. Run `go mod tidy` to install the dependencies.

### Setup:
1. **Install PostgreSQL:**
   - Install PostgreSQL following the official guide: [PostgreSQL installation](https://www.postgresql.org/download/).
   - Create a user and a database:
     ```sql
     CREATE USER go_user WITH PASSWORD 'password';
     CREATE DATABASE go_db;
     ALTER ROLE go_user SET client_encoding TO 'utf8';
     ALTER ROLE go_user SET default_transaction_isolation TO 'read committed';
     ALTER ROLE go_user SET timezone TO 'UTC';
     GRANT ALL PRIVILEGES ON DATABASE go_db TO go_user;
     ```

2. **Create a `.env` File:**
   In the root directory, create a `.env` file (see `.env.exemple` file) containing the following:
   ```env
   PORT=5000
   DBHOST=localhost
   DBUSER=go_user
   DBUSERPW=password
   DBNAME=go_db
   DBPORT=5432
3. **Run the API**
- In the root of the project, run:
  ```bash
  go run main.go
## How to Test It
You can use `Postman` (There is a provided Postman collection `TODO_COLLECTION.postman_collection`)
Or use the `curl` tool to test the API endpoints.

### Example Requests

1. **GET (Retrieve Data):**
   ```bash
   curl -X GET http://localhost:5000/api/todos
2. **POST (Create Data):**
   ```bash
   curl -X POST http://localhost:5000/api/todos -H "Content-Type: application/json" -d '{"body": "New Task"}'
3. **PATCH (Update Data):**
   ```bash
   curl -X PATCH http://localhost:5000/api/todos/1 -H "Content-Type: application/json" -d '{"completed": true}'
5. **DELETE (Delete Data):**
   ```bash
   curl -X DELETE http://localhost:5000/api/todos/1

You can also use query parameters by replacing `c.Params()` with `c.Query()`, then test with: `curl -X DELETE http://localhost:5000/api/todos?id=1` see [Fiber Docs](https://docs.gofiber.io/api/ctx/#query).

## Technologies & Tools Used
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" height="40"/> [Go Documentation](https://go.dev/doc/)
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/fiber/fiber-original.svg" width="40" height="40"/> [Fiber Framework](https://docs.gofiber.io/)
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/golang/golang-original.svg" width="40" height="40"/> [GORM Documentation](https://gorm.io/docs/)
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" width="40" height="40"/> [PostgreSQL Documentation](https://www.postgresql.org/docs/)
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postman/postman-original.svg" width="40" height="40"/> [Postman Documentation](https://learning.postman.com/)
