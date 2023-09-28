# Start from a minimal golang base image
FROM golang:alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o order-management .

# Start from a fresh image
FROM alpine

# Copy the built binary from the previous stage
COPY --from=builder /app/order-management /app/order-management

# Set the working directory in the container
WORKDIR /app

# Expose the port the application uses
EXPOSE 4300

# Command to run the application
CMD ["./order-management"]
