Go HTTP Servers with Docker

This project contains two HTTP server implementations in Go. One uses the standard net/http library, and the other uses the gin-gonic framework.

server_standard.go

A basic HTTP server using the net/http library.
server_gin.go

An HTTP server using the gin-gonic framework.
Prerequisites

    Docker
    Docker Compose
    Make

Building and Running the Project
Using Docker Compose

To build and run the services using Docker Compose, follow these steps:

Build the Docker images and start the containers:

   docker-compose up --build

This command will build the Docker images for both servers and start the containers. The docker-compose.yml file specifies the services:

    httpServerGin: Runs on port 8081 and maps to the gin-gonic based server.
    httpServer: Runs on port 8082 and maps to the standard HTTP server.

Using Makefile

The provided Makefile can be used to streamline the build and run process.

To build the binaries:

make build

To format the code:

make fmt

To lint the code:

make lint

To build the Docker images:

make docker-build

To run the Docker containers:

make docker-run

Accessing the Services

Once the containers are running, you can access the services via the following URLs:

    http://localhost:8081/datetime: Access the gin-gonic based server.
    http://localhost:8082/datetime: Access the standard HTTP server.

Example Requests
curl

To get the current date and time from the servers, you can use curl:

sh

curl http://localhost:8081/datetime
curl http://localhost:8082/datetime

Browser

Alternatively, you can open your web browser and navigate to the above URLs to see the responses.
