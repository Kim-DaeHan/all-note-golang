# ALL Note App API

## Introduction

This project is a backend API for a note application. It is built using the Gin framework and provides CRUD functionality for users, departments, notes, todos, projects, project tasks, meetings, and recruitments.

## Features

User Management: Create, read, update, delete users
Department Management: Create, read, update, delete departments
Note Management: Create, read, update, delete notes
Todo Management: Create, read, update, delete todos
Project Management: Create, read, update, delete projects
Project Task Management: Create, read, update, delete project tasks
Meeting Management: Create, read, update, delete meetings
Job Application Management: Create, read, update, delete recruitment postings

## Getting Started

### Prerequisites

- Go (v1.22.2+)
- Gin framework
- Mongo database
- Other required libraries are managed via go mod.

### Installation

1. Clone the repository:

```
git clone https://github.com/Kim-DaeHan/all-note-golang.git
cd all-note-golang
```

2. Install the necessary libraries:

```
go mod tidy
```

3. Add your database configuration to the .env file:

```
cp .env.example .env
```

4. Start the server:

```
go run main.go
```

The server will run on localhost:8080.

## API Doc

http://localhost:8080/docs/index.html

## Contributing

Contributions are welcome! To contribute to this project, follow these steps:

Fork the repository.
Create a feature branch (git checkout -b feature/AmazingFeature).
Commit your changes (git commit -m 'Add some AmazingFeature').
Push to the branch (git push origin feature/AmazingFeature).
Open a pull request.
