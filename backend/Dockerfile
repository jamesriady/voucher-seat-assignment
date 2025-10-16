FROM golang:1.22-alpine AS builder

# Install necessary CGO dependencies for SQLite
RUN apk add --no-cache gcc musl-dev

# Set the current working directory inside the container 
WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

ENV CGO_ENABLED=1
RUN go build -ldflags "-linkmode external -extldflags '-static'" -o main .

FROM alpine:3.20.2

RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container
WORKDIR /root/

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 8000

# Command to run the executable
CMD ["./main"]