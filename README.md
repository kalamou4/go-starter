
# Go Authentication Starter

A lightweight, secure authentication starter kit for Go applications using standard packages. This project provides a solid foundation for implementing authentication in your Go web applications without external authentication libraries.

## Features
- JWT-based authentication
- Secure password hashing
- Middleware for protected routes
- Environment-based configuration
- RESTful authentication endpoints
- Built with standard Go packages

## Project Structure
```
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── auth/
│   │   ├── auth.go         # Authentication logic
│   │   └── handler.go      # HTTP handlers for auth routes
│   ├── config/
│   │   └── config.go       # Configuration management
│   └── middleware/
│       └── middleware.go   # Authentication middleware
└── .env                    # Environment variables
```

## Prerequisites
- Go 1.21 or higher
- A text editor or IDE
- Basic understanding of Go and authentication concepts

## Installation

Clone the repository:
```bash
git clone https://github.com/yourusername/go-auth-starter.git
cd go-auth-starter
```

Install dependencies:
```bash
go mod tidy
```

Create a `.env` file in the root directory:
```env
PORT=8080
JWT_SECRET=your_jwt_secret_key
# Add other environment variables as needed
```

## Quick Start

1. Configure your environment variables in `.env`.
2. Run the application:
   ```bash
   go run cmd/api/main.go
   ```
   The server will start on [http://localhost:8080](http://localhost:8080) (or your configured port).

## API Endpoints

### Authentication Routes
- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `GET /auth/profile` (Protected route)

### Example Requests

#### Register a new user
```bash
curl -X POST http://localhost:8080/auth/register   -H "Content-Type: application/json"   -d '{"email": "user@example.com", "password": "securepassword"}'
```

#### Login
```bash
curl -X POST http://localhost:8080/auth/login   -H "Content-Type: application/json"   -d '{"email": "user@example.com", "password": "securepassword"}'
```

## Configuration

Configure the application through environment variables in the `.env` file:
```env
PORT=8080                    # Server port
JWT_SECRET=your_secret_key    # JWT signing key
JWT_EXPIRY=24h               # JWT expiration time
```

## Security Considerations
- Uses secure password hashing
- Implements JWT token-based authentication
- Includes middleware for protecting routes
- Environment-based configuration for sensitive data

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Go standard library
- Contributors and maintainers

## Support

For support, please open an issue in the GitHub repository or contact the maintainers.

