# Fiber MVC Boilerplate

A boilerplate project for building web applications in Go using the Fiber framework. This project follows the **MVC (Model-View-Controller)** pattern to provide a clean and modular structure for your application.

## Features

- **Fiber Framework**: High-performance and lightweight Go web framework.
- **MVC Structure**: Organized into models, views, and controllers for scalability and maintainability.
- **Ready to Extend**: Perfect starting point for building RESTful APIs or web applications.

---

## Project Structure

    ```bash
    go-messenger/
    ├── controllers/               <-- Contains application controllers (e.g., user controller)
    ├── database/                  <-- Database configuration files
    ├── models/                    <-- Contains application models (e.g., user model)
    ├── routes/                    <-- Handles route definitions
    ├── services/                  <-- Contains the business logic
    ├── main.go                    <-- Entry point of the application
    └── README.md                  <-- Project documentation
    ```

---

## Getting Started

### Prerequisites

- **Golang**: Ensure you have Go installed on your machine. [Download Go](https://golang.org/dl/)
- **Fiber**: Fiber framework is included in the dependencies.

---

### Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:Mlcarvalho1/go-messenger.git
   cd go-messenger
   ```

2. Install dependencies:

    ```
    go mod tidy
    ```

3. Run the application:

    ```
    go run main.go
    ```

4. Access the application:

    Open your browser and navigate to http://localhost:3000/user

## Running tests

- **Integration Tests**
    In te terminal run:
    ```
        make test-integration
    ```

    To receive the coverage report in html run:

    ```
        make test-integration-report
    ```
