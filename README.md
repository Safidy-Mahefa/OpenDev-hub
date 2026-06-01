# OpenDev — Backend API

REST API backend for the OpenDev platform. Built with Go, Fiber, and PostgreSQL.

---

## Table of contents

- [Overview](#overview)
- [Requirements](#requirements)
- [Project structure](#project-structure)
- [Getting started](#getting-started)
- [Environment variables](#environment-variables)
- [Database](#database)
- [API reference](#api-reference)
- [Dependencies](#dependencies)
- [Known limitations](#known-limitations)

---

## Overview

OpenDev is a platform for developers. This repository contains the backend API only. It exposes HTTP endpoints over Fiber, connects to a PostgreSQL database via sqlx, and is structured to remain extensible as the platform grows.

The current state of the project covers:

- HTTP server setup with request logging and panic recovery
- PostgreSQL connection management
- User model with read and insert operations
- Basic routing with a health check endpoint

Authentication (JWT), input validation, and CORS are not yet implemented and are in progress.

---

## Requirements

- Go 1.21 or later
- PostgreSQL 14 or later
- pgAdmin 4 (optional, for visual database management)

---

## Project structure

```
.
├── cmd/
│   └── api/
│       └── main.go           # Entry point. Loads env, connects to database, starts server.
├── internal/
│   ├── database/
│   │   └── database.go       # Opens and holds the PostgreSQL connection.
│   ├── server/
│   │   ├── server.go         # Fiber instance setup, middleware registration.
│   │   └── routes.go         # Route declarations.
│   └── users/
│       ├── user.go           # User struct with db and json tags.
│       └── repository.go     # Database queries for the users table.
├── migrations/
│   ├── 001_create_user.sql
│   ├── 002_create_project.sql
│   └── 003_create_challenge.sql
├── .env.example
├── go.mod
└── go.sum
```

The `internal/` directory is Go-enforced private. Nothing inside it can be imported by an external module.

---

## Getting started

**1. Clone the repository**

```bash
git clone https://github.com/your-username/opendev.git
cd opendev
```

**2. Install dependencies**

```bash
go mod download
```

**3. Configure environment variables**

Copy the example file and fill in your values:

```bash
cp .env.example .env
```

See [Environment variables](#environment-variables) for the full list.

**4. Set up the database**

Create a database named `opendev` in PostgreSQL, then run the migrations in order. Using psql:

```bash
psql -U postgres -d opendev -f migrations/001_create_user.sql
psql -U postgres -d opendev -f migrations/002_create_project.sql
psql -U postgres -d opendev -f migrations/003_create_challenge.sql
```

Or paste each file into the pgAdmin Query Tool and execute.

**5. Start the server**

```bash
go run cmd/api/main.go
```

The server starts on the port defined in your `.env` file. Default is `3000`.

---

## Environment variables

| Variable | Required | Description |
|---|---|---|
| `PORT` | No | Port the HTTP server listens on. Defaults to `3000`. |
| `ENV` | No | Runtime environment. Example: `development`, `production`. |
| `DATABASE_URL` | Yes | Full PostgreSQL connection string. |

Example `DATABASE_URL`:

```
postgres://postgres:yourpassword@localhost:5432/opendev?sslmode=disable
```

The `.env` file is listed in `.gitignore` and must never be committed.

---

## Database

The database schema is managed through sequential SQL migration files located in `migrations/`. Each file represents one state change. They must be applied in order and are never modified after the fact. New changes always go into a new file.

**Current schema**

`001_create_user.sql` defines the `users` table:

| Column | Type | Constraints |
|---|---|---|
| id | VARCHAR(10) | PRIMARY KEY |
| username | VARCHAR(255) | UNIQUE, NOT NULL |
| email | VARCHAR(255) | NOT NULL |
| name | VARCHAR(255) | nullable |
| bio | VARCHAR(500) | nullable |
| avatarurl | VARCHAR(255) | nullable |
| role | VARCHAR(100) | NOT NULL |
| githubusername | VARCHAR(20) | nullable |
| linkedinurl | VARCHAR(25) | nullable |
| portfoliourl | VARCHAR(25) | nullable |
| city | VARCHAR(15) | nullable |
| totalpoints | INT | DEFAULT 0 |
| seasonpoints | INT | DEFAULT 0 |
| createdat | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |
| updatedat | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |

Note: PostgreSQL lowercases all unquoted identifiers. Column names in Go struct tags use the lowercased form.

`002_create_project.sql` and `003_create_challenge.sql` are reserved and not yet defined.

---

## API reference

Base URL: `http://localhost:{PORT}`

### Health check

```
GET /health
```

Returns the server status and current environment.

Response:

```json
{
  "Statue": "Ok",
  "env": "development"
}
```

---

### Ping

```
GET /ping
```

Basic liveness check.

Response:

```json
{
  "message": "pong"
}
```

---

### List users

```
GET /users
```

Returns all users from the database.

Response:

```json
{
  "Utilisateurs": [
    {
      "id": "Safidy06",
      "username": "Safidy06",
      "email": "example@email.com",
      "name": null,
      "bio": null,
      "avatarurl": null,
      "role": "admin",
      "githubusername": null,
      "linkedinurl": null,
      "portfoliourl": null,
      "city": null,
      "totalpoints": 0,
      "seasonpoints": 0,
      "createdat": "2025-01-01T00:00:00Z",
      "updatedat": "2025-01-01T00:00:00Z"
    }
  ]
}
```

Note: This route currently calls `users.Create` on every request as a development fixture. This is temporary and will be removed.

---

## Dependencies

| Package | Version | Purpose |
|---|---|---|
| `github.com/gofiber/fiber/v2` | v2.52.13 | HTTP framework |
| `github.com/jmoiron/sqlx` | v1.4.0 | SQL with struct mapping |
| `github.com/lib/pq` | v1.12.3 | PostgreSQL driver |
| `github.com/joho/godotenv` | v1.5.1 | Load .env file |
| `github.com/google/uuid` | v1.6.0 | UUID generation |

---

## Known limitations

The following are known gaps in the current implementation. They are tracked here for the next development phases.

- No authentication. All routes are publicly accessible.
- No input validation. Malformed request bodies are not rejected cleanly.
- No CORS configuration. Cross-origin requests from a frontend will be blocked.
- The `GET /users` route inserts a hardcoded user on every call. This is a development artifact and must be removed before any staging or production deployment.
- The migration SQL file uses `CREATE TABLE USER` which conflicts with a reserved keyword in PostgreSQL. It should be `CREATE TABLE users`.
- The `.env.example` file is empty. It should document all required variables.
- `updatedAt` is set once at insert time and is never updated on row modification.
- No structured logging. `fmt.Println` is used in several places instead of a logger.`
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv         <!-- Pour utiliser les fichiers .env -->
`
