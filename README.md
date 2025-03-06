#  Go Product API 🚀 

A RESTful API for product management built with Go. This project features user authentication, CRUD operations for products and categories, and comprehensive API documentation.

## Table of Contents 📑 
- [Go Product API 🚀](#go-product-api-)
  - [Table of Contents 📑](#table-of-contents-)
  - [Features ✨](#features-)
  - [Tech Stack 🔧](#tech-stack-)
  - [Project Structure 📁](#project-structure-)
- [Getting Started 🎯](#getting-started-)
  - [Prerequisites 📋](#prerequisites-)
  - [Installation 🔧](#installation-)
  - [Usage 💻](#usage-)
    - [Running the Application 🚀](#running-the-application-)
  - [API Documentation 📚](#api-documentation-)

## Features ✨
- 🔐 **Authentication:** JWT-based authentication system
- 👥 **User Management:** CRUD operations for users
- 📦 **Product Management:** CRUD operations with filtering and pagination
- 🗂️ **Category Management:** CRUD operations for product categories
- 🔄 **Middleware:**
  - 🛡️ Authentication
  - ⚠️ Error handling
  - 📊 Performance logging
  - 💨 Response caching

## Tech Stack 🔧
- **Go**
- **Gin** – Web framework
- **GORM** – ORM library
- **SQL Server** – Database
- **Swagger** – API documentation

## Project Structure 📁
```
src/
├── controllers/       # API endpoint handlers
├── database/          # Database connection
├── docs/              # Swagger documentation
├── middleware/        # Custom middleware components
├── models/            # Data models
├── routes/            # API route definitions
├── services/          # Business logic
├── utils/             # Helper functions
└── main.go            # Application entry point
```

# Getting Started 🎯

## Prerequisites 📋
- Go 1.24.0
- SQL Server instance

## Installation 🔧
1. **Install dependencies:**
   ```bash
   go mod download
   ```
2. **Configure environment variables:**
   - Create a `.env` file with the required configurations.

## Usage 💻
### Running the Application 🚀
- **Locally:**
  ```bash
  go run main.go
  ```
- **Using Visual Studio Code:**
  Open the project in VS Code, then press F5 or use the "Run and Debug" panel.

## API Documentation 📚
After starting the application, access the documentation at:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)