# Project Overview

This document provides a comprehensive overview of the technologies, tools, and project structure employed in the development of the backend, focusing on APIs for file upload and serving, with an emphasis on handling large files (100MB+) and utilizing IPFS for storage.

## Technologies

1. **Golang**: The primary programming language used for implementing the backend. Golang is known for its performance and efficiency.

2. **Gin**: The Gin web framework for Golang was used to create and manage the HTTP server, handle routing, and facilitate middleware integration.

3. **IPFS (InterPlanetary File System)**: IPFS was chosen as the storage solution for file management. IPFS is a distributed, peer-to-peer hypermedia protocol designed to make the web faster, safer, and more open.

4. **Docker**: Docker is used for containerization, making it easy to package the application and its dependencies into a consistent and isolated environment.

## APIs

### 1. File Upload API

- **HTTP Method**: POST
- **Endpoint**: /upload
- **Description**: Allows users to upload a file. Handles large files, including those larger than 100MB.
- **Request Payload**: Form data with the file to be uploaded.
- **Response Payload**:

  ```json
  {
    "fileId": "",
    "size": "",
    "timestamp": ""
  }

### 2. Serving a File

- **HTTP Method**: GET
- **Endpoint**: /file/:fileId
- **Description**: Allows users to get  a file using fileId.
- **Request Param**: FileID.
- **Response Payload**: `<File Content>`
  
## Project Structure

- `app.env`: Configuration file that holds environment variables used by the application.

- `assets/`: Directory for storing files and assets. This can include uploaded files, static resources, or any other data required by the application.

- `cmd/`: Directory that contains the main entry point of the application.

  - `main.go`: The main application file that initializes and runs the web server.

- `docker-compose.yaml`: Configuration file for Docker Compose, which defines services, networks, and volumes for your application.

- `Dockerfile`: The Dockerfile used to build a Docker image for your application.

- `go.mod` and `go.sum`: Go modules files that manage dependencies for your Go project.

- `initiator/`: Directory that contains the initialization code for your application.

  - `initiator.go`: Initialization code for setting up the application, including loading configuration, connecting to IPFS, and starting the web server.

- `internal/`: The internal directory contains components that are specific to your application but not exposed as public APIs.

  - `constant/`: Constants and configuration models.

    - `model/`: Model definitions, such as the `config.go` file that defines the project's configuration.

  - `glue/`: Components responsible for integrating and routing requests.

    - `routing/`: Routing configuration, such as `file_routing.go`, which handles file-related routes.

  - `handlers/`: Request and response handling components.

    - `rest/`: RESTful API handlers, including `file_handler.go` for file-related operations.

  - `pkg/`: Reusable and shareable components.

    - `storage/`: Components related to storage, including `local_storage.go` for managing local file storage.

- `platform/`: Platform-specific components and libraries.

  - `ipfs/`: Components related to IPFS integration.

    - `ipfs.go`: IPFS client for interacting with the IPFS network.

  - `logger/`: Logging configurations and utilities.

    - `logrus.go`: Logging setup using Logrus.

  - `routers/`: Router configurations for the web server.

    - `routers.go`: Definition of routing interfaces and configurations.

- `start.sh`: A shell script used to start the application. This script may handle environment variable setup and application execution.

## Project Overview

This project is organized to separate concerns and maintain a clean and structured codebase. Key components, such as file handling, IPFS integration, routing, and configuration, are organized into their respective directories. The use of Docker and Docker Compose allows for easy containerization and deployment of the application.

## Running the Application

To run the application and set up the development environment, follow these steps:

1. Clone the project repository to your local machine:

   ```bash
   git clone https://github.com/aleale2121/ipfs-fileverse.git
   cd ipfs-fileverse
   docker-compose up --build

2. Open the postman collection from your agent and test the endpoints

# Writing Unit Tests and End-to-End Tests

This section provides guidelines on writing unit tests and end-to-end tests for your project. Writing tests ensures the reliability and correctness of your application.

## Unit Tests

Unit tests are written to test small, isolated units of your code, such as functions, methods, or individual components. They focus on verifying that these units perform as expected. To write unit tests for your project:

1. **Organize Your Tests**: Create a directory for your tests, typically named `tests` or `unit_tests`, in your project's root directory.

2. **Use the Go Testing Framework**: Go has a built-in testing framework that makes it easy to write and run unit tests. To write a test file for a Go package, create a file with a name ending in `_test.go`.

3. **Test Function Naming**: Test functions should start with `Test` and describe what is being tested. For example, if you are testing a function `CalculateSum`, you can name your test function `TestCalculateSum`.

4. **Use Test Cases**: Write test cases to cover different scenarios and edge cases for your code. Use the `t.Run` function to define sub-tests within a test function.

5. **Assertions**: Use the `t.Errorf` or `t.Fail` functions to report test failures. Ensure that your test cases include assertions to verify the expected results.

6. **Run Tests**: To run the tests, use the `go test` command followed by the package or test file you want to test. For example, to run tests in the `tests` directory, use:

   ```bash
   go test ./tests

Here are some of Golang Testing Lbraries includin the builtin.

1. **Testing Package (builtin)**:
   Go's built-in `testing` package provides a comprehensive testing framework that includes testing functions like `testing.T`, `testing.B`, and `testing.M`. It supports writing unit tests and benchmarks and is a standard choice for most Go projects.

2. **Testify**:
   [Testify](https://github.com/stretchr/testify) is a popular testing toolkit that adds additional functionality and assertions to the standard Go testing library. It provides features like assertions, mock objects, and suites for organizing tests.

3. **Ginkgo and Gomega**:
   [Ginkgo](https://github.com/onsi/ginkgo) is a BDD-style testing framework for Go, and [Gomega](https://github.com/onsi/gomega) is its corresponding assertion library. Ginkgo allows you to write expressive and readable tests using BDD-style descriptions, making it suitable for more complex tests.

4. **GoConvey**:
   [GoConvey](https://github.com/smartystreets/goconvey) is a tool for writing behavioral tests in Go. It provides real-time test results and has a web-based UI for visually displaying test output. It's ideal for BDD-style testing and test-driven development (TDD).

5. **GoMock**:
   [GoMock](https://github.com/golang/mock) is a mocking framework for Go. It allows you to create and use mock objects in your tests, which is useful for isolating components and dependencies during unit testing.

6. **TestSuite**:
   [TestSuite](https://github.com/stretchr/testify/suite) is a part of the Testify toolkit and provides a framework for organizing tests into test suites. It makes it easier to set up and tear down common test contexts for a group of tests.

## End-to-End (E2E) Tests

End-to-End (E2E) tests validate your application's behavior across multiple components and systems, ensuring it functions correctly from start to finish.

### What are E2E Tests?

E2E tests, also known as integration tests, simulate real user scenarios and validate an application's behavior across multiple components, from the user interface to APIs and databases.

### E2E Testing in Golang

Golang offers several E2E testing frameworks and libraries. Some popular options include:

- **Selenium:** A versatile framework for web application testing with Golang bindings.

- **Ginkgo:** A BDD-style testing framework for writing expressive E2E tests in Golang.

Choose a framework that aligns with your project's requirements and technology stack.

### E2E Testing in Node.js

Node.js developers have access to a variety of E2E testing tools and frameworks, such as:

- **Cypress:** A JavaScript-based E2E testing framework known for its simplicity and speed.

- **Puppeteer:** A headless Chrome browser that can be controlled programmatically for E2E testing.

- **Jest:** A popular testing framework for JavaScript applications that can be extended for E2E testing.

Select a Node.js testing tool that best suits your project's needs.

## Writing E2E Tests

E2E tests involve creating test scenarios, defining user actions, and making assertions about the application's behavior. They are often written using a behavior-driven development (BDD) approach, using plain language that is easy to understand for both developers and non-developers.

Organize your E2E tests in a maintainable and structured manner to ensure they remain up-to-date as your application evolves.
