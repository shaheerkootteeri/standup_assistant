# Standup Updates Application

This is a simple web application for managing standup updates of team members. It provides endpoints for submitting updates and fetching updates via HTTP requests.

## Getting Started

1. **Prerequisites:**
   - Go (Golang) installed on your machine.
   - PostgreSQL database set up with appropriate permissions. You can also use other databases supported by Go, but the application is configured to work with PostgreSQL.

2. **Setting Up the Database:**
   - Create a PostgreSQL database and update the connection string in `standup/database.go` with your database credentials.
   - Run the SQL script provided in `db_setup.sql` to create the necessary table.

3. **Running the Application:**
   - Navigate to the project directory in your terminal.
   - Run the command: `go run main.go`.
   - The server will start at `http://localhost:8080`.

4. **Using a Different Database:**
   - If you're using a different database, such as MySQL or SQLite, you'll need to adjust the connection string in `standup/database.go`.
   - Replace the `connStr` variable in the `init()` function with the appropriate connection URL for your database.

## Project Structure

- `main.go`: Entry point of the application. Initializes the HTTP server.
- `standup/`: Package containing the application logic.
  - `router.go`: Defines the HTTP routes using Gorilla Mux.
  - `handlers.go`: Contains HTTP request handlers for various endpoints.
  - `database.go`: Handles database operations.
  - `model.go`: Defines the data model for standup updates.
- `index.html`: HTML template for the web interface.
- `script.js`: JavaScript file for client-side interactions.

## Contributing
.

## License

This project is licensed under the [MIT License](LICENSE).
