# Docker file for recipestats
# Start from a Golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Change to the directory containing the main.go file
WORKDIR /app/cmd/recipestats

# Build the Go app
RUN go build -v -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/app/cmd/assetmgmt/main"]