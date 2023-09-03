# Set the base image to Go version 1.21.0
# FROM baseImage:tag
FROM golang:1.21.0

# Labels are metadata that provide additional information about image
# LABEL key="value"
LABEL maintainer="Vishnuprakash P <vishkash.k@gmail.com>"
LABEL project="go-docker"
LABEL version="v1.0"
LABEL description="A Go web server using Gin framework"

# Set the working directory inside the container to /app
# WORKDIR /the/workdir/path
WORKDIR /app

# Copy Go module files (go.mod and go.sum) to the container
# COPY source dest
COPY go.mod go.sum ./

# Download and verify Go module dependencies
# RUN command
RUN go mod download && go mod verify

# Copy the entire local directory into the container
# COPY source dest
COPY . .

# Build the Go application and create an executable binary named "main"
# RUN command
RUN go build -o main .

# Inform Docker that the container will listen on port 8080 at runtime
# EXPOSE port
EXPOSE 8080

# Use docker's secret management, etc to handle sensitive information securely
# Set the environment variables(not sensitive) needed for production
# ENV key=value
ENV GIN_MODE=release

# Define the default command to run the application
# CMD [ "executable" ]
CMD [ "./main" ]
