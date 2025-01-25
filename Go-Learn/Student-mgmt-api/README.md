
Overview

    Student Management REST API is a backend service built with Go (Golang) to manage student records. It provides endpoints to perform CRUD (Create, Read, Update, Delete) operations on student data and is designed to be lightweight, scalable, and easy to use.

Features

    Add Students: Create new student records.
    View Students: Fetch details of all students or a specific student by ID.
    Update Students: Modify existing student records.
    Delete Students: Remove student records from the system.

Tech Stack

    Programming Language: Go (Golang)
    Database: PostgreSQL (default) or SQLite for lightweight setups
    Frameworks/Libraries:
        Gin for HTTP request handling
        GORM for ORM-based database interaction
        Viper for configuration management

API Endpoints
Base URL: http://localhost:8080/api/v1

Endpoints:

    Method 	Endpoint 	Description 	Request Body 	Response (Success)
    GET 	/students 	Fetch all students 	None 	[{id, name, age, grade, created_at}]
    GET 	/students/{id} 	Fetch student by ID 	None 	{id, name, age, grade, created_at}
    POST 	/students 	Add a new student 	{ "name": "John Doe", "age": 20, "grade": "A" } 	{ "message": "Student added" }
    PUT 	/students/{id} 	Update an existing student 	{ "name": "John Smith", "age": 21, "grade": "A+" } 	{ "message": "Student updated" }
    DELETE 	/students/{id} 	Delete a student by ID 	None 	{ "message": "Student deleted" }
    Setup and Installation

Prerequisites

    Go (version 1.19 or later)
    Docker
    Git

Steps

    1.Clone the repository:

        git clone https://github.com/your-username/student-management-api.git
        cd student-management-api

    2.Install dependencies:

        go mod tidy

    3.Set environment variables: Create a .env file at the project root:

        DB_HOST=localhost
        DB_PORT=5432
        DB_USER=your_db_user
        DB_PASSWORD=your_db_password
        DB_NAME=student_db

    4.Run database migrations:

        go run migrations/main.go

    5.Start the server:

        go run main.go

    
Folder Structure

student-mgt-api/

    ├── controllers/ - Contains HTTP handler functions

    ├── models/ - Database models and schema definitions

    ├── routes/ - API route definitions

    ├── config/ - Configuration setup (e.g., database connection)

    ├── migrations/ - Database migration scripts

    ├── main.go - Entry point of the application

    └── README.md - Documentation