# Go Project

## Overview
This project is a scalable Go application designed to manage questions and projects through a RESTful API. It utilizes PostgreSQL as the database and includes multiple HTTP endpoints for managing questions and projects.

## Directory Structure
```
go-project
├── cmd
│   └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── postgres.go
│   ├── handlers
│   │   ├── question_handler.go
│   │   └── project_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── models
│   │   ├── question.go
│   │   └── project.go
│   ├── repositories
│   │   ├── question_repository.go
│   │   └── project_repository.go
│   ├── routes
│   │   ├── question_routes.go
│   │   └── project_routes.go
│   └── services
│       ├── question_service.go
│       └── project_service.go
├── pkg
│   └── utils
│       └── utils.go
├── go.mod
└── README.md
```

## Features
- **Question Service**: Manage questions with CRUD operations.
- **Project Service**: Manage projects with CRUD operations.
- **API Key Authentication**: All endpoints are secured with API key validation.
- **PostgreSQL Database**: Utilizes PostgreSQL for data storage.

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-project
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up the PostgreSQL database and configure the connection settings in `internal/config/config.go`.

4. Run the application:
   ```
   go run cmd/main.go
   ```

## API Endpoints
### Question Service
- **GET /questions**: Retrieve all questions.
- **POST /questions**: Create a new question.
- **PUT /questions/{id}**: Update an existing question.
- **DELETE /questions/{id}**: Delete a question.

### Project Service
- **GET /projects**: Retrieve all projects.
- **POST /projects**: Create a new project.
- **PUT /projects/{id}**: Update an existing project.
- **DELETE /projects/{id}**: Delete a project.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License
This project is licensed under the MIT License.