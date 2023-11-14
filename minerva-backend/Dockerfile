# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed dependencies
RUN cd ./src && go mod tidy && cd ..

# Build the Go app
RUN go build -C ./src -o ../main ./cmd

# Remove everyting except "main" and ".env"
RUN find . ! -name 'main' ! -name '.env' ! -name '.' ! -name '..' -exec rm -rf {} +

# Expose port 8080 to the outside world
EXPOSE 8134

# Run the executable
CMD ["./main"]
