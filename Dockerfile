# Build stage
FROM golang:latest AS builder


# Copy the application code
COPY . /app

# Set the working directory
WORKDIR /app

# Build the Golang application
RUN go build -o server_app .


# Production stage
FROM debian

# Make a directory for the app
RUN mkdir /app

# Set the working directory
WORKDIR /app

# Copy built binary and configuration from build stage
COPY --from=builder /app/server_app .


# Expose port
EXPOSE 8080

CMD ["./server_app"]