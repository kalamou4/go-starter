# Go Authentication Starter

A production-ready starter template for building authenticated web applications with Go.

## Features

- User authentication (register/login)
- JWT token-based authentication
- Database migrations
- Docker support
- Environmental configuration
- Clean architecture

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and update the values
3. Run `docker-compose up` to start the database
4. Run `make migrate-up` to apply migrations
5. Run `make run` to start the server

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   ├── config/
│   ├── database/
│   └── user/
├── pkg/
│   ├── hash/
│   └── jwt/
```

## License

MIT
