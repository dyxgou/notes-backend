# Use the official Go image as build stage
FROM golang:1.24.5-alpine AS builder

# Install build dependencies for CGO and SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Enable CGO for SQLite
ENV CGO_ENABLED=1
ENV GOOS=linux

# Create bin directory and build the application
RUN mkdir -p bin
RUN go build -a -installsuffix cgo -o bin/server ./cmd/api

# Production stage
FROM alpine:latest

# Install SQLite runtime
RUN apk --no-cache add ca-certificates sqlite

# Create app directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/bin/server .
COPY --from=builder /app/db .

# Set environment variables
RUN mkdir -p data
ENV DB_PATH=data/database.db
VOLUME [ "/data" ]

ENV PORT=:5000

# Expose port
EXPOSE 5000

RUN ls -la
RUN sqlite3 ${DB_PATH} < migrations/delete_tables.down.sql
RUN sqlite3 ${DB_PATH} < migrations/create_tables.up.sql

# Run the application
CMD ["./server"]
