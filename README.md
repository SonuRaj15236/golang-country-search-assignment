# Golang Country Search Assignment

This project is a simple REST API built using **Go (Golang)** that fetches country details by name and exposes them via a backend service.

The project is implemented using **only the Go standard library**. No external frameworks or third‑party libraries are used.

---

## 🚀 Features

- Search country details by country name
- Clean and simple REST API
- Uses public country data API internally
- Proper project structure (handler, service, model)
- Unit tests included
- No external dependencies

---

## 🧱 Project Structure

```
Assignment/
├── main.go
├── go.mod
├── handler/
│   └── handler.go
├── service/
│   ├── service.go
│   └── service_test.go
├── model/
│   └── countryModel.go
└── README.md
```

---

## 📦 API Details

### Endpoint
```
GET /country?name={country_name}
```

### Full Browser URL
After running the application:
```
http://localhost:8000/country?name=India
```

### Example Response
```json
{
  "name": "India",
  "capital": "New Delhi",
  "currency": "INR",
  "population": 1380004385
}
```

> Response field order:
> 1. name  
> 2. capital  
> 3. currency  
> 4. population  

---

## ▶️ How to Run the Project

### Prerequisites
- Go installed (Go 1.20+ recommended)

Check Go version:
```bash
go version
```

### Run the application
```bash
go run main.go
```

Server will start on:
```
http://localhost:8000
```

---

## 🧪 Run Tests

### Run all tests
```bash
go test ./...
```

### Run tests with coverage
```bash
go test ./... -cover
```

### Generate detailed coverage report
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 🔧 Notes

- Only Go standard library is used
- No external frameworks
- Clean separation of concerns
- Suitable for backend assignment submission

---

## 👨‍💻 Author

**Sonu Raj**  
Golang Backend Developer

---

## 📌 Assignment Objective

This project demonstrates:
- REST API development in Go
- Clean project architecture
- External API consumption
- Unit testing and coverage using Go tooling