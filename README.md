# 🧮 Dockerized Go Calculator

A simple command-line calculator built with Go and containerized using Docker.

---

## 📁 Project Structure
```bash
.
├── calculator.go # Main Go source file
├── dockerfile # Dockerfile to build the image
└── go.mod # Go module (for dependency and version management)
```

---

## 🚀 Features

- Basic CLI-based calculator
- Accepts two numbers and an operator (+, -, *, /)
- Built with Go
- Dockerized for portability and consistency

---

## 🛠️ Technologies Used

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

---

## 🐳 Getting Started with Docker

### 🔧 Build the Docker Image

```bash
docker build -t go-calculator .
```
▶️ Run the Docker Container

```bash
docker run -it go-calculator
```


🧪 Example Input/Output
```bash
Enter first number: 10
Enter operator (+, -, *, /): *
Enter second number: 3
Result: 30
```
📌 Notes
Make sure you have Go installed locally if not using Docker.
A go.mod file is already included for module support.
