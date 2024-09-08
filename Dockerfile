# Use a specific version of Go for consistency (e.g., go1.20)
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests first
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Verify Go modules and build the application
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main .

# Expose port 8080 (adjust if your app uses a different port)
EXPOSE 3000

# Set the entry point to the compiled binary
ENTRYPOINT ["/app/main"]


