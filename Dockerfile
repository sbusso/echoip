# Stage 1: Build the Go binary
FROM golang:1.21 as builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echoip .

# Stage 2: Build a small image
FROM scratch

# Set a default environment variable for the port
ENV PORT 3000

# Copy the binary from the builder stage
COPY --from=builder /app/echoip /echoip

# Run the binary
ENTRYPOINT ["/echoip"]
