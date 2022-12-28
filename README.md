# Golang Base API

## Structure

The structure of this project is:

- **main.go:** main Go file
- **routes/:** Creates routes to the corresponding handler
- **controllers/** Handles all the logic behind validating request parameters, query and finally sending **Responses** with correct codes.
- **services/:** Handles extra logic before data interaction with db
- **repository/:** Handles the interaction with the DB, independently of what may be (SQL, NoSQL, API, etc.)
- **db/:** Handles the db configuration
- **models/:** Handles the schema definition of the Models
- **factories/:** Handles the objects creation
