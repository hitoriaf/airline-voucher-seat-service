# Airline Voucher Seat Service Backend

## Summary

This is a backend service with the primary function of creating vouchers to be assigned to 3 unique seat numbers. Seat numbers are also adjusted according to the aircraft type.

This service uses RESTful API communication, built with Go programming language using the Gin framework.

## Tech Stacks

- **Language**: Go
- **Framework**: Gin
- **ORM**: GORM
- **Database**: SQLite
- **Architecture**: RESTful API


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

**Response (Voucher exists):**
```json
{
  "exist": true
}
```

**Response (Voucher not exists):**
```json
{
  "exist": false
}
```

**Response (Error):**
```json
{
  "error": "Invalid request body",
  "message": "Key: 'CheckVoucherRequest.FlightNumber' Error:Field validation for 'FlightNumber' failed on the 'required' tag"
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

**Response (New voucher created):**
```json
{
  "success": true,
  "seats": ["12A", "15C", "8B"]
}
```

**Response (Voucher already exists):**
```json
{
  "success": false,
  "seats": ["12A", "15C", "8B"]
}
```

**Response (Error - Invalid date format):**
```json
{
  "error": "Invalid date format",
  "message": "Flight date must be in YYYY-MM-DD format"
}
```

**Response (Error - Missing required field):**
```json
{
  "error": "Invalid request body",
  "message": "Key: 'GenerateVoucherRequest.CrewName' Error:Field validation for 'CrewName' failed on the 'required' tag"
}
```

## Additional Notes
- Aircraft types are limited to 3 types: ATR, Airbus 320, and Boeing 737 Max.
- This service automatically creates sqlite db files on specified path in env file.