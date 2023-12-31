# Use an official Golang runtime as a base image
FROM golang:1.17-alpine as builder

# Set the working directory to /app
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o golang-task-quiz .

# Use a minimal Alpine image as the final base image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage to the final image
COPY --from=builder /app/golang-task-quiz /app/golang-task-quiz

# Expose the port the application runs on
EXPOSE 8080

# Run the application when the container starts
CMD ["./golang-task-quiz"]
