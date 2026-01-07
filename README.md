# CarZone API

A comprehensive REST API for managing cars and engines built with Go, PostgreSQL, and modern monitoring tools.

## Features

### 🚗 Car Management
- Create, read, update, and delete cars
- Search cars by brand
- Engine relationship via foreign key
- Complete data validation

### 🔧 Engine Management
- Full CRUD operations for engines
- Technical specifications including displacement, cylinder count, and range

### 🔐 Authentication
- JWT token-based login system
- Middleware for protecting routes
- Credential validation

### 📊 Monitoring & Observability
- **Prometheus**: Metrics collection
- **Grafana**: Visualization and dashboards
- **OpenTelemetry + Jaeger**: Distributed tracing

### 🐳 Containerization
- Docker and docker-compose for development and production
- Containerized PostgreSQL
- Database health checks

## Architecture

The project uses a layered architecture (layered architecture):

```
┌─────────────┐
│  Handlers   │  ← HTTP endpoints
├─────────────┤
│  Services   │  ← Business logic
├─────────────┤
│   Stores    │  ← Data access
├─────────────┤
│  Database   │  ← PostgreSQL
└─────────────┘
```

### Technologies

- **Go 1.25**: Programming language
- **Gorilla Mux**: HTTP router
- **PostgreSQL**: Database
- **JWT**: Authentication
- **OpenTelemetry**: Tracing
- **Prometheus**: Metrics
- **Grafana**: Dashboards
- **Jaeger**: Tracing UI
- **Docker**: Containerization

## API Endpoints

### Authentication
- `POST /login` - Login and receive JWT token

### Cars (Protected)
- `GET /cars/{id}` - Get car by ID
- `GET /cars?brand={brand}` - Get cars by brand
- `POST /cars` - Create new car
- `PUT /cars/{id}` - Update car
- `DELETE /cars/{id}` - Delete car

### Engines (Protected)
- `GET /engine/{id}` - Get engine by ID
- `POST /engine` - Create new engine
- `PUT /engine/{id}` - Update engine
- `DELETE /engine/{id}` - Delete engine

### Monitoring
- `GET /metrics` - Prometheus metrics endpoint

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.25+ (optional for local development)

### Setup with Docker

```bash
# Clone the repository
git clone <repository-url>
cd CarZone

# Run all services
docker-compose up --build

# Or run in background
docker-compose up -d --build
```

### Service Access

- **API**: http://localhost:8080
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (username: admin, password: admin)
- **Jaeger UI**: http://localhost:16686

### Environment Variables

For local development, create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgress
DB_PASSWORD=12345
DB_NAME=postgress
PORT=8080
JAEGER_AGENT_HOST=localhost
JAEGER_AGENT_PORT=4318
```

## Data Models

### Car
```json
{
  "id": "uuid",
  "name": "string",
  "year": "string",
  "brand": "string",
  "fuelType": "Gasoline|Diesel|Electric|Hybrid",
  "engine": {...},
  "price": "float64",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Engine
```json
{
  "engine_id": "uuid",
  "displacement": "int64",
  "no_of_cylinder": "int64",
  "car_range": "int64"
}
```

## Usage Examples

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

### Get Cars
```bash
curl -X GET http://localhost:8080/cars \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create Car
```bash
curl -X POST http://localhost:8080/cars \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Tesla Model 3",
    "year": "2024",
    "brand": "Tesla",
    "fuelType": "Electric",
    "engine": {
      "engine_id": "uuid-here",
      "displacement": 0,
      "noOfCylinders": 0,
      "carRange": 350
    },
    "price": 45000.00
  }'
```

## Development

### Project Structure

```
CarZone/
├── db/                    # Database Dockerfile
├── driver/               # Database connection
├── handler/              # HTTP handlers
│   ├── car/
│   ├── engine/
│   └── login/
├── middleware/           # Auth & metrics middleware
├── models/               # Data models & validation
├── service/              # Business logic
├── store/                # Data access layer
│   ├── car/
│   ├── engine/
│   └── schema.sql       # Database schema
├── docker-compose.yml    # Multi-service setup
├── Dockerfile           # Application container
├── prometheus.yml       # Metrics config
├── main.go              # Application entry point
└── go.mod
```

### Running Tests

```bash
go test ./...
```

### Build Binary

```bash
go build -o carzone .
```

## Monitoring

### Metrics
- Request count
- Response time
- Error rate
- Resource usage

### Tracing
- Request tracing across services
- Bottleneck identification
- Performance issue debugging

## License

This project is licensed under the MIT License.
