# High-Level Design Document for "Hello World" Go Web Server

## Overview

The "Hello World" Go Web Server is a simple web application that listens for incoming HTTP requests and responds with the text "Hello, World!". This application serves as an introductory example of how to create a basic web server using the Go programming language.

## Objectives

- **Demonstrate** how to set up a basic HTTP server in Go.
- **Respond** to HTTP GET requests with a simple text message.
- **Utilize** standard Go libraries to keep the application lightweight and straightforward.

## Libraries Used

The application relies on the following libraries:

- **net/http**: Provides HTTP client and server implementations.
- **zerolog**: Offers logging capabilities for recording server activities and errors.

## General Concepts

### HTTP Server Initialization

- **Server Creation**: The application initializes an HTTP server using the net/http package.
- **Port Listening**: It listens on a specified port (commonly 8080) for incoming HTTP requests.

### Request Handling

- **Handler Functions**: A handler function is defined to manage incoming requests. When a request is received at the root URL path (/), the handler responds with "Hello, World!".
- **Routing**: The http.HandleFunc method is used to associate URL paths with handler functions.

### Logging

- **Startup Logs**: Upon server startup, a log message confirms that the server is running and listening on the specified port.
- **Error Handling**: Errors encountered during server execution are logged using the zerolog package.

## Execution Flow

1. **Import Necessary Packages**: The application starts by importing the required Go libraries.
2. **Define the Handler Function**:
   - A function is created to handle HTTP requests.
   - This function writes "Hello, World!" to the HTTP response.
3. **Set Up Routing**:
   - The handler function is linked to the root URL path (/) using http.HandleFunc.
4. **Start the Server**:
   - The server begins listening for incoming requests on the specified port using http.ListenAndServe.
   - Startup confirmation is logged.
5. **Graceful Shutdown**:
   - The server can gracefully shut down upon receiving an interrupt signal, ensuring no requests are left unprocessed.
6. **Handle Incoming Requests**:
   - When a request is received at the root path, the server invokes the handler function.
   - The response "Hello, World!" is sent back to the client.

## Configuration

- **Port Number**: The server listens on port 8080 by default. This can be changed by modifying the port value in the server setup.
- **Network Interface**: The server binds to all available network interfaces (0.0.0.0), making it accessible from any network interface on the host machine.
- **Timeouts**: The server includes configurable read and write timeouts to ensure requests are handled within a reasonable time frame.
- **Log Level**: The logging level (e.g., debug, info) can be configured using command-line flags.

## Deployment

- **Local Deployment**: Suitable for running on a local machine for testing and development purposes.
- **Containerization**: Can be containerized using Docker for deployment in different environments or cloud platforms.
- **Cloud Deployment**: Adaptable for deployment on cloud services like AWS, GCP, or Azure with minimal adjustments.

## Testing

To verify that the server is functioning correctly:

1. **Start the Server**: Run the application to start the server.
2. **Access the Server**:
   - Open a web browser and navigate to http://localhost:8080/.
   - Alternatively, use a command-line tool like curl:
     
     ```
     curl http://localhost:8080/
     ```
3. **Verify the Response**: Ensure that the response is "Hello, World!".

## Usage Example

To run the server with all available flags:

```sh
go run main.go -addr ":8080" -read-timeout 10s -write-timeout 10s -log-level "info"
```

- **-addr**: Specifies the address and port to listen on (e.g., ":8080").
- **-read-timeout**: Sets the maximum duration for reading the request (e.g., "10s" for 10 seconds).
- **-write-timeout**: Sets the maximum duration for writing the response (e.g., "10s" for 10 seconds).
- **-log-level**: Sets the logging level (e.g., "info", "debug", "warn", etc.).

## Summary

This "Hello World" web server serves as a foundational example of creating a web server in Go. It demonstrates the basic concepts of handling HTTP requests and responses, setting up routing, utilizing Go's standard libraries, and implementing graceful shutdown and enhanced logging to build a simple yet functional application.

