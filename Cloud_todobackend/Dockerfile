# Use a lightweight Go image
FROM golang:1.22

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./main"]
