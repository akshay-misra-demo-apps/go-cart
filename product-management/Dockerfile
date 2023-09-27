# Start from a minimal golang base image
FROM golang:alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o product-management .

# Start from a fresh image
FROM alpine

# Copy the built binary from the previous stage
COPY --from=builder /app/product-management /app/product-management

# Set the working directory in the container
WORKDIR /app

# Expose the port the application uses
EXPOSE 4100

# Command to run the application
CMD ["./product-management"]
