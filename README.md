# JSON-CRUD-API

A simple RESTful API built with Go, Gin, and GORM for creating, reading, updating, and deleting JSON data.

## Features

* **Create:** Allows creating new posts.  `POST /posts`
* **Read:** Retrieves a list of posts or a single post by ID. `GET /posts`, `GET /posts/:id`
* **Update:** Updates an existing post by ID. `PUT /posts/:id`
* **Delete:** Deletes a post by ID. `DELETE /posts/:id`

## Technologies Used

* **Go:** The programming language.
* **Gin:** A web framework for Go.
* **GORM:** An Object-Relational Mapper (ORM) for Go.
* **PostgreSQL:** (Assumed) The database used (Adapt if using a different database).


## Setup

1. **Install Go:** Ensure you have Go installed on your system.  [https://go.dev/doc/install](https://go.dev/doc/install)
2. **Clone the repository:**
   ```bash
   git clone https://github.com/Francismensah/JSON-CRUD-API.git
   ```
3. **Install dependencies:**
   ```bash
   go mod tidy
   ```
4. **Set up environment variables:** Create a `.env` file in the root directory and set the `DB_URL` environment variable with your PostgreSQL connection string.  Example:
   ```
   DB_URL="postgres://user:password@host:port/database?sslmode=disable"
   ```
5. **Run migrations:** Run the `main.go` file to create the database tables and start the server.  This also runs the migrations.
   ```bash
   go run main.go
   ```

## API Endpoints

* **POST `/posts`:** Creates a new post. Request body should be a JSON object with `title` and `body` fields.
* **GET `/posts`:** Retrieves a list of all posts.
* **GET `/posts/:id`:** Retrieves a single post by ID.
* **PUT `/posts/:id`:** Updates an existing post.  Request body should be a JSON object with `title` and `body` fields.
* **DELETE `/posts/:id`:** Deletes a post by ID.


## Future Enhancements

* Add error handling and input validation.
* Improve API documentation.
* Implement authentication and authorization.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.