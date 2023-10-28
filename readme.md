
# gRPC Blog Platform

This is a gRPC-based Blog Platform. It provides gRPC APIs for creating, retrieving, updating, and deleting blog posts. The project is implemented in Go and follows standard design patterns for gRPC services.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Running the Server](#running-the-server)
  - [Using the Client](#using-the-client)
- [Project Details](#project-details)
  - [Proto Definitions](#proto-definitions)
  - [gRPC Service](#grpc-service)
- [Testing](#testing)

## Project Structure

The project is organized as follows:

- `api/`: Contains the Protobuf definitions for the gRPC service.
- `client/`: The client application.
- `server/`: The server application.
- `README.md`: The project's README file.

## Getting Started

### Prerequisites

Before you start, make sure you have the following software installed:

- Go: [https://golang.org/doc/install](https://golang.org/doc/install)
- Protocol Buffers: [https://developers.google.com/protocol-buffers/docs/gotutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/b0a04gl/gRPCBlogX.git
   ```

2. Change to the project directory:

   ```bash
   cd gRPCBlogX
   ```

3. Generate Go code from the .proto file:

   ```bash
   protoc --go_out=. --go-grpc_out=. api/blog.proto
   ```

4. Install the project dependencies:

   ```bash
   go mod tidy
   ```

## Usage

### Running the Server

To run the gRPC server, navigate to the `server` directory and execute the following command:

```bash
go run main.go
```

The server should start and listen on port 50051.

### Using the Client

To interact with the gRPC server, navigate to the `client` directory and execute the following commands:

- Create a blog:

  ```bash
  go run main.go create "My First Blog" "This is the content of my first blog."
  ```

- Get a blog by ID:

  ```bash
  go run main.go get <blog_id>
  ```

- Update a blog by ID:

  ```bash
  go run main.go update <blog_id> "Updated Title" "Updated Content"
  ```

- Delete a blog by ID:

  ```bash
  go run main.go delete <blog_id>
  ```

- List all blogs:

  ```bash
  go run main.go list
  ```

## Project Details

### Proto Definitions

The project uses Protocol Buffers (Protobuf) for defining the gRPC service. You can find the definitions in `api/blog.proto`. These definitions specify the message formats and service methods.

### gRPC Service

The gRPC service, implemented in Go, follows the standard design patterns for gRPC services. The server and client code can be found in the respective `server` and `client` directories.

## Testing

The project includes unit tests to verify the functionality of the gRPC service. You can run the tests using the `go test` command.

```bash
go test ./...
```
