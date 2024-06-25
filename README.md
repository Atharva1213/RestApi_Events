# Event REST API using Go with Mux and MySQL

## Project Overview
This project is a RESTful API built with Go language using the Mux router and MySQL database. It provides endpoints to manage events, including CRUD operations (Create, Read, Update, Delete).

## Technologies Used
- Go programming language
- Mux HTTP router
- MySQL database

## Prerequisites
- Go (version >= X.X)
- MySQL (version >= X.X)
- MySQL driver for Go (e.g., github.com/go-sql-driver/mysql)

## Installation
1. Clone the repository:
  git clone https://github.com/your-username/event-api.git
2. Navigate into the project directory:
  cd event-api
3. Install dependencies (if any):
  go mod tidy
4. Set up the MySQL database:
- Create a MySQL database named `eventdb`.
- Import the SQL schema from `database/schema.sql` to create the necessary tables.

5. Configure the database connection:
- Open `config/config.go` and update the MySQL connection details (host, username, password).

6. Build and run the application:
  go build
  ./event-api

## API Endpoints
- **GET /events**: Get all events
- **GET /events/{id}**: Get a specific event by ID
- **POST /events**: Create a new event
- **PUT /events/{id}**: Update an existing event
- **DELETE /events/{id}**: Delete an event by ID

## Usage
1. Open your API testing tool (e.g., Postman).
2. Use the above endpoints to interact with the API:
- Send GET requests to retrieve events.
- Send POST requests with JSON payload to create new events.
- Send PUT requests with JSON payload to update existing events.
- Send DELETE requests to delete events by ID.

## Error Handling
- The API returns appropriate HTTP status codes and error messages for different scenarios (e.g., invalid requests, database errors).

## Future Improvements
- Implement authentication and authorization for secure endpoints.
- Add pagination support for fetching large datasets.
- Implement validation for input data to ensure data integrity.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
