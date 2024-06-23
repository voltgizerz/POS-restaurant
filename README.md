# REST API Go for POS
This project provides simple REST APIs for POS (Point of Sale) software.
# Build With
- Go 1.20 or higher
- Fiber
- MySQL
- JWT Auth

# How to Run Project

Follow these steps to run the project:

1. Setup Your Database Environment

2. Run Database Migrations `goose -dir=./database/migrations mysql "root@tcp(localhost:3306)/db_pos?parseTime=true" up`

3. Run the Application `go run ./cmd/app.go`
