# Go_Auth 🔐

A production-grade authentication system built with Go, featuring JWT-based stateless authentication, secure password hashing, PostgreSQL database integration, and middleware-based authorization with clean separation of concerns.

## 🌟 Features

- **User Registration & Login**:  Secure signup and login endpoints with email/password authentication
- **JWT Authentication**: Stateless authentication using access tokens (15 min) and refresh tokens (7 days)
- **Password Security**: Bcrypt hashing with cost factor 12 for secure password storage
- **PostgreSQL Integration**: pgx/v5 connection pooling for efficient database operations
- **Protected Routes**: Middleware-based authorization for secure endpoints
- **Comprehensive User Profiles**: Support for personal information and address details
- **Clean Architecture**: Well-organized codebase with separation of concerns

## 📁 Project Structure

```
Go_Auth/
├── db/                  # Database connection and configuration
│   └── db.go           # PostgreSQL connection pool setup
├── handlers/           # HTTP request handlers
│   ├── auth.go        # Signup and login handlers
│   └── profile.go     # User profile handlers
├── middleware/        # HTTP middleware
│   └── auth. go       # JWT authentication middleware
├── models/           # Data models
│   └── user.go      # User struct definition
├── routes/          # Route registration
│   └── routes.go   # API route definitions
├── services/       # Business logic
│   ├── jwt.go     # JWT token generation and validation
│   └── password.go # Password hashing and verification
├── tmp/           # Temporary files (for Air hot-reload)
├── .air.toml      # Air configuration for development
├── .gitignore     # Git ignore rules
├── go.mod         # Go module dependencies
├── go.sum         # Dependency checksums
└── main.go        # Application entry point
```

## 🚀 Getting Started

### Prerequisites

- Go 1.23. 2 or higher
- PostgreSQL database
- Environment variables configured

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/M-ayank2005/Go_Auth.git
   cd Go_Auth
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/dbname? sslmode=disable
   JWT_SECRET=your-super-secret-jwt-key
   ```

4. **Create the database schema**
   ```sql
   CREATE TABLE users (
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       email VARCHAR(255) UNIQUE NOT NULL,
       password_hash VARCHAR(255) NOT NULL,
       first_name VARCHAR(100),
       last_name VARCHAR(100),
       date_of_birth DATE,
       phone VARCHAR(20),
       address_line_1 VARCHAR(255),
       address_line_2 VARCHAR(255),
       city VARCHAR(100),
       state VARCHAR(100),
       country VARCHAR(100),
       postal_code VARCHAR(20),
       is_email_verified BOOLEAN DEFAULT FALSE,
       is_active BOOLEAN DEFAULT TRUE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

### Development Mode with Hot Reload

Install Air for hot-reloading:
```bash
go install github.com/air-verse/air@latest
air
```

## 📡 API Endpoints

### Public Endpoints

#### **POST** `/auth/signup`
Register a new user account. 

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe",
  "date_of_birth": "1990-01-01T00:00:00Z",
  "phone": "+1234567890",
  "address_line_1": "123 Main St",
  "address_line_2": "Apt 4B",
  "city": "New York",
  "state": "NY",
  "country": "USA",
  "postal_code":  "10001"
}
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

#### **POST** `/auth/login`
Authenticate an existing user.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### Protected Endpoints

#### **GET** `/profile`
Retrieve the authenticated user's profile information.

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "date_of_birth": "1990-01-01T00:00:00Z",
  "phone": "+1234567890",
  "address_line_1": "123 Main St",
  "address_line_2":  "Apt 4B",
  "city": "New York",
  "state": "NY",
  "country": "USA",
  "postal_code": "10001",
  "is_email_verified": false,
  "created_at": "2025-12-20T06:00:00Z"
}
```

## 🔧 Technologies Used

- **[Go](https://go.dev/)** - Programming language
- **[PostgreSQL](https://www.postgresql.org/)** - Database
- **[pgx/v5](https://github.com/jackc/pgx)** - PostgreSQL driver and toolkit
- **[JWT (golang-jwt/jwt)](https://github.com/golang-jwt/jwt)** - JSON Web Tokens
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable management
- **[Air](https://github.com/air-verse/air)** - Live reload for Go apps

## 🔒 Security Features

1. **Password Hashing**:  Bcrypt with cost factor 12
2. **Token-Based Authentication**: JWT with HS256 signing algorithm
3. **Token Expiration**: Short-lived access tokens (15 minutes) and long-lived refresh tokens (7 days)
4. **SQL Injection Protection**: Parameterized queries via pgx
5. **Password Requirements**: Minimum 8 characters
6. **Email Normalization**: Lowercase and trimmed email addresses
7. **Secure Password Serialization**: Password hash excluded from JSON responses

## 🏗️ Architecture

### Clean Separation of Concerns

- **`db/`**:  Database connection management and configuration
- **`handlers/`**: HTTP request/response handling
- **`middleware/`**: Request processing pipeline (authentication, logging)
- **`models/`**: Data structures and domain models
- **`routes/`**: API endpoint registration
- **`services/`**: Business logic (JWT, password hashing)

### Database Connection Pooling

The application uses pgx connection pooling for efficient database operations:
- Automatic connection reuse
- Health checks via ping
- Query execution mode configuration

### Middleware Architecture

Protected routes are wrapped with authentication middleware that:
1. Extracts JWT from Authorization header
2. Validates token signature and expiration
3. Injects user ID into request context
4. Allows/denies access based on token validity

## 📝 License

This project is open source and available under the MIT License. 

## 👤 Author

**M-ayank2005**
- GitHub: [@M-ayank2005](https://github.com/M-ayank2005)

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! 

## 📧 Support

For support, please open an issue in the GitHub repository.

---

**Built with ❤️ using Go**
