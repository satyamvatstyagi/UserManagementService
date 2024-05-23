# User Management Service

## Description

The User Management Service is a GoLang service that provides functionality for managing user accounts.

## Features

- Health Check Endpoint: `/user/health`
- User Registration Endpoint: `/user/register`
- User Login Endpoint: `/user/login`
- Get User by Username Endpoint: `/user/{username}`
- Get Fibonacci Number Endpoint: `/user/fibonacci/{number}`

## Installation

1. Clone the repository:

     ```bash
     https://github.com/satyamvatstyagi/UserManagementService.git
     ```

2. Install the dependencies:

     ```bash
     go mod download
     ```

3. Build the service:

     ```bash
     go build
     ```

## Usage

1. Start the service:

     ```bash
     ./user-management-service
     ```

2. Access the service API at `http://localhost:8080`.

## Configuration

The service can be configured using environment variables. The following variables are available:

- `DATABASE_USER`: The username for the database connection.
- `DATABASE_PASSWORD`: The password for the database connection.
- `DATABASE_NAME`: The name of the database.
- `DATABASE_HOST`: The host of the database.
- `DATABASE_PORT`: The port of the database.
- `SERVER_PORT`: The port for the server to listen on.
- `GIN_MODE`: The mode for the Gin framework.
- `BASIC_AUTH_USER`: The username for basic authentication.
- `BASIC_AUTH_PASSWORD`: The password for basic authentication.
- `JWT_SECRET`: The secret key used for JWT token generation.
- `JWT_EXPIRY`: The expiry time for JWT tokens in minutes.
- `LOG_FILE_NAME`: The name of the log file.

## Contributing

Contributions are welcome! Please read the [contribution guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [SATYAM License](LICENSE).