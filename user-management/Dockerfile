# Start from a minimal golang base image
FROM golang:alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o user-management .

# Start from a fresh image
FROM alpine

# Copy the built binary from the previous stage
COPY --from=builder /app/user-management /app/user-management

# Set the working directory in the container
WORKDIR /app

# Expose the port the application uses
EXPOSE 4000

# Command to run the application
CMD ["./user-management"]
