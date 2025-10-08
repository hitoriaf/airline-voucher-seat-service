# Airline Voucher Seat Service

## Summary
A service for generating airline seat vouchers. This project includes both frontend and backend components.
- **Frontend**: React with Next.js framework and Tailwind CSS
- **Backend**: Golang with Gin framework and GORM for database

## Tech Stack

### Backend
- **Language**: Go
- **Framework**: Gin
- **Database ORM**: GORM
- **Database**: SQLite

### Frontend
- **Framework**: NextJS with React 19
- **Language**: TypeScript
- **Styling**: Tailwind CSS


## Installation

### Prerequisites
- Docker & Docker Compose (Recommended)
- Go 1.25.1+ (for manual / development setup without docker)
- Node.js 18+ & npm (for manual / development setup without docker)

### Running with Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone <repository-url> <projectName>
cd <projectName>
```

2. Create Environment File
```bash
cd backend
cp .env.example .env
cd ..
```

3. Start all services:
Ensure you are at the root directory.
```bash
docker-compose up --build
```

3. Access the application:
- **Frontend**: http://localhost
- **Backend API**: http://localhost:8080

### Running Manually without Docker for Development

#### Backend Setup
1. Navigate to backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Create environment file:
```bash
cp .env.example .env
# Edit .env file with your configuration
```

4. Run the backend server:
```bash
go run src/main.go
```
Backend will be available at http://localhost:8080

#### Frontend Setup
1. Navigate to frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Run development server:
```bash
npm run dev
```
Frontend will be available at http://localhost:3000

## Environment Variables

### Backend (.env)
```bash
# Database configuration
DB_PATH=./data/sql.db

# Server configuration
PORT=8080
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/check` | Check voucher by flight number and date |
| POST | `/api/generate` | Generate new voucher for crew |

### API Examples

#### Check Voucher
```bash
POST http://localhost:8080/api/check
Content-Type: application/json

{
  "flightNumber": "SQ270",
  "date": "2025-10-10"
}
```

#### Generate Voucher
```bash
POST http://localhost:8080/api/generate
Content-Type: application/json

{
  "name": "Hitori",
  "id": "EMP100",
  "flightNumber": "SQ270",
  "date": "2025-10-10",
  "aircraft": "Boeing 737 Max"
}
```