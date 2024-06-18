# Use the official Golang image for building the application
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /project/otoo/

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and the .env file into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /project/otoo/build/myapp .

# Use a minimal image for the final stage
FROM alpine:latest

# Copy the built application from the builder stage
COPY --from=builder /project/otoo/build/myapp /project/otoo/build/myapp

# # Copy the .env file to the final image (if needed by the application)
# COPY --from=builder /project/otoo/.env /project/otoo/.env

# # Set the environment variable for the .env file path (if necessary)
# ENV ENV_FILE_PATH=/project/otoo/.env

# Expose the application port
EXPOSE 8081

# Set the entry point for the container to run the application
ENTRYPOINT [ "/project/otoo/build/myapp" ]
