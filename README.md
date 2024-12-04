# Golang Hexagonal Architecture Example

This repository demonstrates the **Hexagonal Architecture** pattern in Go (Golang), providing a clean and modular approach to building scalable applications. The project focuses on maintaining separation of concerns between business logic, application services, and infrastructure.

## Features

- **Hexagonal Architecture**:
  - Core domain logic is isolated from frameworks and external dependencies.
  - Clear distinction between ports (interfaces) and adapters (implementations).
- **Docker Integration**:
  - Pre-configured `Dockerfile` for containerized deployment.
- **Modular Design**:
  - Separate packages for domain, application, and infrastructure layers.

## Project Structure
- /application
  - Application services and use cases
- /config
  - Configuration files
- /domain
  - Domain entities and business rules 
- /infrastructure
  - Adapters for persistence, external APIs, etc.
- /presentation 
  - API or UI-facing components
- main.go 
  - Application entry point
 
## Prerequisites

- Go 1.20 or newer
- Docker (optional for containerization)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/RenzoReccio/golang-hexagonal-architecture.git
   cd golang-hexagonal-architecture
    ```

2. Install dependencies:
   ```bash
   go mod tidy
    ```

3. Run the application:
   ```bash
   go run main.go
    ```

## Contributing
Feel free to submit issues or pull requests. Contributions are welcome!


## Contact
Author: Renzo Reccio

Let me know if you'd like to add more details or modify anything!
