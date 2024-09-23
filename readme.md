# Crowdfunding App

This **Crowdfunding** application is a platform for creating and supporting fundraising campaigns. The app is built using **Golang (Gin framework)** and **PostgreSQL** as the database.

## Table of Contents

- [Features](#features)
- [Configuration](#configuration)
- [How to Run](#how-to-run)
- [Project Structure](#project-structure)
- [License](#license)

## Features

- **Authentication**: Register, login, forgot password using JWT.
- **Campaign Management**: Create, update, and view campaign details.
- **Donation**: Users can contribute to campaigns.

## Configuration

Create a `.env` file in the project root with the following format:

## Environment Variables

| Key                     | Example Value                                               | Description                                              |
|-------------------------|-------------------------------------------------------------|----------------------------------------------------------|
| `DATABASE_URL`           | `postgres://uname:pass@host:5432/go_cf_db?sslmode=disable`  | URL for PostgreSQL connection.                           |
| `PORT`                   | `8080`                                                     | Port to run the server.                                  |
| `JWT_AUTH_SECRET`        | `"auth"`                                                   | Secret key for JWT authentication token.                 |
| `JWT_REFRESH_SECRET`     | `"refresh"`                                                | Secret key for JWT refresh token.                        |
| `JWT_AUTH_EXP_IN_HOUR`   | `24`                                                       | Expiration time for authentication token (in hours).     |
| `JWT_REFRESH_EXP_IN_HOUR`| `48`                                                       | Expiration time for refresh token (in hours).            |

## How to Run

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Create the .env file based on the configuration above.

4. Run the database migrations:
    ```bash
    go run cmd/migrate.go
    ```
5. Run the application:
    ```bash
    go run cmd/main.go
    ```
    The application will be available at http://localhost:8080.

## Project Structure

```bash
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   └── v1/
│       └── auth/         # Authentication module
│       └── campaign/     # Campaign module
│   └── common/
│       └── models/       # Database models
│       └── services/     # Business logic and services
├── pkg/
│   └── response/         # API response standardization
│   └── middleware/       # Middleware like CORS
└── config/               # Application configuration loading
```

## License
This application is licensed under the MIT License.