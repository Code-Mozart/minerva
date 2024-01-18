# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy only the go.mod and go.sum files to take advantage of caching
COPY src/go.mod src/go.sum /app/src/

# Copy the rest of the files in the current directory into the container at /app
COPY . /app

# Set the GOPATH environment variable
ENV GOPATH=/go

# Install any needed dependencies
WORKDIR /app/src
RUN go mod download
WORKDIR /app

# Build the Go app
# go build [options] [build flags] [packages]
#   -C ./src    change directory to ./src
#   -o ../main  output to ../main
#   ./cmd       path to the main package
RUN go build -C ./src -o ../main ./cmd

# Remove everyting except "main" and ".env"
# for hardening
# RUN find . ! -name 'main' ! -name '.env' ! -name '.' ! -name '..' -exec rm -rf {} +

# Expose port 8134 to the outside world
EXPOSE 8134

# Run the executable
ENTRYPOINT ["./main"]
