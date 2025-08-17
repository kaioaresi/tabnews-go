# TabNews-Go Documentation

## 1. Introduction

TabNews-Go is a project that aims to recreate the functionalities of [TabNews](https://www.tabnews.com.br/) using the Go programming language. This documentation provides a comprehensive overview of the project, its structure, and how to use it.

## 2. Installation

To get started with TabNews-Go, you need to have Go installed on your system. You can download it from the official website: [https://golang.org/](https://golang.org/)

Once you have Go installed, you can clone the repository and install the dependencies:

```bash
git clone https://github.com/aresi/tabnews-go.git
cd tabnews-go
go mod tidy
```

## 3. API Endpoints

The following are the available API endpoints:

*   `GET /`: Returns a welcome message.
*   `GET /api/v1/status`: Returns the status of the application, including database information.
*   `GET /api/v1/migrations`: Returns a message for the migrations page.

## 4. Packages

The project is organized into the following packages:

### 4.1. `db`

This package is responsible for handling the database connection. It provides a `DBAccess` interface and a `DBConfig` struct to manage the database connection. It also includes functions to get database information, such as the version, max connections, and current connections.

### 4.2. `logger`

This package provides a logging interface and an implementation using the `zap` library. It allows for logging at different levels, such as `Info`, `Error`, and `Warning`.

### 4.3. `web`

This package is responsible for handling the web server and routing. It uses the `http` package to create the server and the `gorilla/mux` router to handle the routes. It also includes handlers for the API endpoints.

## 5. Configuration

The database connection can be configured using the `.env.development` file. This file should contain the following variables:

```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

## 6. Running the Application

To run the application, you can use the following command:

```bash
go run cmd/tabnews-go/main.go
```

This will start the web server on port `8080`.

## 7. Dependencies

The main dependencies of the project are:

*   [gorilla/mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Go.
*   [lib/pq](https://github.com/lib/pq): A pure Go Postgres driver for the `database/sql` package.
*   [go.uber.org/zap](https://github.com/uber-go/zap): A blazing fast, structured, leveled logger in Go.
