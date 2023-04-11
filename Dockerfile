# Use the official Golang base image
FROM golang:1.20.3-alpine3.17

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8000

# Run the application
CMD ["./main"]

