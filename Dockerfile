# Use official Golang image as a build stage
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /  # This sets the working directory to the root directory

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o pt_aka_tech_test

# Use a minimal base image
FROM alpine:latest

# Set the working directory in the new container
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /pt_aka_tech_test .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./pt_aka_tech_test"]
