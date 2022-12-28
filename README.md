# Golang Base API

## Structure

The structure of this project is:

- **main.go:** main file
- **routes.go:** Creates routes to the corresponding handler
- **controllers.go:** Handles all the logic behind validating request parameters, query and finally sending **Responses** with correct codes.
- **services.go:** Handles extra logic before data interaction with db
- **repository.go:** Handles the interaction with the DB, independently of what may be (SQL, NoSQL, API, etc.)
