# Greater Altitude - School Management System API Documentation

## Overview
Greater Altitude is a school management system designed to enhance and streamline the Montessori education experience, specifically tailored for nursery schools. This documentation provides details on the various API endpoints available, their usage, and the associated models.

## Technologies Used
- **Go (Golang)**: Performance and scalability.
- **Gin Framework**: RESTful API development.
- **JWT Authentication**: Secure user login.
- **GORM**: Simplified database CRUD operations.
- **PostgreSQL**: Reliable data storage.
- **Redis**: Caching.

## Requirements
- Familiarity with Golang.
- PostgreSQL server with user and database setup. By default, use PostgreSQL with user `postgres`, database `postgres`, and no password. Check PostgreSQL configuration docs before creating a user if this is your first time using it.
- Redis server must be downloaded and running.

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/greater-altitude.git
    ```

2. **Navigate to the project directory:**
    ```bash
    cd greater-altitude
    ```

3. **Install dependencies:**
    ```bash
    go mod tidy
    ```

4. **Setup PostgreSQL and Redis:**
    - Ensure PostgreSQL is running and properly configured.
    - Ensure Redis server is running.

## Run the Application
```bash
go run main.go

# API Documentation

## Authentication Endpoints

### POST /login
**Description**: Logs in a user by providing the correct credentials.

**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "userpassword"
}
