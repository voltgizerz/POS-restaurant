# REST API for POS
This project provides simple REST APIs for Point of Sale (POS) software.

## Built With
- Go 1.21.3
- MySQL
- JWT Authentication

## Frameworks and Library
- Fiber V3: HTTP Framework
- Goose: Database Migrations
- SQLX: Database ORM
- Validator V10: Struct Validation
- Logrus: Application Logging
- Jaeger: Opentracing
- Cleanenv: Environment Configuration
- Mockgen: Mock Generation

## How to Run the Project

Follow these steps to run the project:

1. **Setup Your Database Environment**
   - Ensure MySQL is installed and running.
   - Create a database named `db_pos`.

2. **Run Database Migrations**
   - Use the following command to run the migrations:
     ```sh
     goose -dir=./database/migrations mysql "root:password@tcp(localhost:3306)/db_pos?parseTime=true" up
     ```
   - Replace `password` with your MySQL root password.

3. **Run the Application**
   - Start the application using the command:
     ```sh
     go run ./cmd/app.go
     ```
