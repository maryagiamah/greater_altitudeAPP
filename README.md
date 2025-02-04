# Greater Altitude - School Management System API Documentation

## Overview
Greater Altitude is a school management system specifically tailored for nursery schools. It is designed to enhance and streamline the Montessori education experience. This documentation details the various API endpoints available, their usage, and the associated models.

## Technologies Used
- **Go (Golang)**: Performance and scalability.
- **Gin Framework**: RESTful API development.
- **JWT Authentication**: Secure user login.
- **GORM**: Simplified database CRUD operations.
- **PostgreSQL**: Reliable data storage.
- **Redis**: Caching.

## Requirements
- Familiarity with Golang.
- A PostgreSQL server with a user and database setup. By default, use PostgreSQL with the user `postgres`, the database `postgres`, and no password. If this is your first use, check the PostgreSQL configuration docs before creating a user.
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
    - Ensure the Redis server is running.

## Run the Application
```bash
go run main.go
```

## Services

### Admin Services

  *Endpoints:*
  * `POST /logout`: Log out the current user.
  * `POST /signup`: Register a new user.
  * `PUT /users/:id/activate` - Activate a user by ID.
  * `PUT /users/:id/deactivate`: Deactivate a user by ID.

### Class Services

  *Endpoints:*
  * `POST /classes`: Create a new class.
  * `GET /classes`: Retrieve all classes.
  * `GET /classes/:id`: Retrieve a specific class by ID.
  * `PUT /classes/:id`: Update a class by ID.
  * `DELETE /classes/:id`: Delete a class by ID.
  * `POST /classes/:id/pupil`: Add a pupil to a class (Work in Progress).
  * `POST /classes/:id/teacher`: Assign a teacher to a class (Work in Progress).
  * `GET /classes/:id/pupils`: Get all pupils in a class.
  * `GET /classes/:id/teachers`: Get all teachers in a class.
  * `GET /classes/:id/activities`: Get all activities for a class (Work in Progress)

### Pupil Services

  *Endpoints:*
  * `POST /pupils`: Create a new pupil.
  * `GET /pupils`: Retrieve all pupils.
  * `GET /pupils/:id`: Retrieve a specific pupil by ID.
  * `PUT /pupils/:id`: Update a pupil by ID.
  * `DELETE /pupils/:id`: Delete a pupil by ID.
  * `GET /pupils/:id/classes`: Get all classes for a pupil (Work in Progress).

### Message Services
*Note: Work in progress*

  *Endpoints:*
  * `POST /messages`: Create a new message.
  * `GET /messages/inbox`: Retrieve inbox message.
  * `GET /messages/sent`: Retrieve sent messages.
  * `PUT /messages/:id/read`: Mark the message as read.
  * `DELETE /messages/:id`: Delete a message.

### Report Services

  *Endpoints:*
  * `POST /reports`: Create a new report.
  * `GET /reports`: Retrieve all reports.
  * `GET /reports/:id`: Retrieve a specific report by ID.
  * `PUT /reports/:id`: Update a report by ID.
  * `DELETE /reports/:id`: Delete a report by ID.
  * `GET /reports/pupil/:pupilId`: Get reports for a specific pupil.
  * `GET /reports/teacher/:teacherId`: Get reports for a specific teacher.

### Event Services
*Note: Work in progress*

   *Endpoints:*
   * `POST /events`: Create a new event.
   * `GET /events`: Retrieve all events.
   * `GET /events/:id`: Retrieve an event by ID.
   * `PUT /events/:id`: Update an event by ID.
   * `DELETE /events/:id`: Delete an event by ID.

### Program Services
#### Programs

   *Endpoints:*
   * `POST /programs`: Create a new program.
   * `GET /programs`: Retrieve all programs.
   * `GET /programs/:id`: Retrieve a specific program by ID.
   * `PUT /programs/:id`: Update a program by ID.
   * `DELETE /programs/:id`: Delete a program by ID.
   * `GET /programs/:id/classes`: Get classes associated with the program.
   * `GET /programs/:id/activities`: Get activities associated with the program (Work in Progress).
   * `POST /programs/:id/class`: Add class to program (Work in Progress).
   * `POST /programs/:id/activity`: Add activity to program (Work in Progress).
  
#### Activities

  *Endpoints:*
  * `POST /activity` - Create an activity.
  * `GET /activity` - Retrieve all activities.
  * `GET /activity/:id` - Retrieve an activity by ID.
  * `PUT/activity/:id` - Update an activity by ID.
  * `DELETE /activity/:id` - Delete an activity.

### Billing Services
#### Invoices
*Note: Work in progress*

  *Endpoints:*
  * `GET /invoices`: Retrieve all invoices (Work in Progress).
  * `GET /invoices/:id`: Retrieve an invoice by ID (Work in Progress).
  * `POST /invoices`: Create an invoice (Work in Progress).
  * `PUT /invoices/:id`: Update an invoice by ID (Work in Progress).
  * `DELETE /invoices/:id`: Delete an invoice by ID (Work in Progress).

#### Payments

  *Endpoints:*
  * `GET /payments`: Retrieve all payments (Work in Progress).
  * `GET /payments/:id`: Retrieve payment details by ID (Work in Progress).
  * `POST /payments`: Create a new payment (Work in Progress).
  * `PUT /payments/:id`: Update payment information (Work in Progress).
  * `DELETE /payments/:id`: Delete payment record (Work in Progress).

### User Services
#### Users
  
  *Endpoints:*
  * `GET /users`: Retrieve all users.
  * `GET /users/:id`: Retrieve user details by ID.
  * `PUT /users/:id`: Update user information.
  * `GET /users/me`: Get authenticated user's profile information.


### Parent Services

  *Endpoints:*
  * `GET /parents`: Retrieve all parents.
  * `GET /parents/:id`: Retrieve parent details by ID.
  * `POST /parents`: Create a new parent record.
  * `PUT /parents/:id`: Update parent information.
  * `DELETE /parents/:id`: Delete parent record.
  * `GET /parents/:id/wards`: Get pupils associated with the parent.
  * `POST /parents/:id/ward`: Add pupil to parent record (Work in Progress)

### Staff Services

  *Endpoints:*
  * `POST /staff`: Create a new staff member record.
  * `GET /staff`: Retrieve all staff members.
  * `GET /staff/:id`: Retrieve staff member details by ID.
  * `PUT /staff/:id`: Update staff member information.
  * `DELETE /staff/:id`: Delete staff member record.

## Important Notes
1. Areas marked as "Work in Progress" are not fully functional and should not be relied upon until completed.
2. For detailed information regarding request body types and examples, please refer to the following Postman link:
       [Postman API Documentation](https://web.postman.co/workspace/Nursery-School-communication-sy~29174a5e-9cfe-4766-bf7d-ae3bdc1948e4/overview).

If you have any questions or need further clarification, please feel free to reach out:
GitHub Username: maryagiamah
Email: under20techie@gmail.com
