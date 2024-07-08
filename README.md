# Datetime Server

This project demonstrates the implementation of a basic HTTP server in Go that returns the current date and time. The server is implemented using net/http library and Gin frameworks and containerized using Docker and Docker Compose.

## Installation

To install the project use:

```bash
go get github.com/codescalersinternships/datetime-server-amr
```

## Usage
Use the Makefile to build and run the project.

to build the HTTP server
```
make build http
```
to build the Gin server
```
make build gin
```
to build the images
```
make build images
```
to launch the containers
```
make launch containers
```
Access the (net/http) server endpoint:

```
Curl http://localhost:8081/datetime
```
Access the Gin server endpoint:
```
Curl http://localhost:8082/datetime
```

## Testing
Run the tests using Go's testing package.
```
go test ./...
```
## Contribution
Feel free to open issues or submit pull requests with improvements.

