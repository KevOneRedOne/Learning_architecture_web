FROM golang:1.21.3

# Set the Current Working Directory inside the container
WORKDIR /my-app-api_go

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY go.mod go.sum ./
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


# #Docker Build
# #docker build -t m1_custom_api_go .

# #Docker Run
# #docker run -d --name m1_custom_api_go --env-file .env -p 8080:8080 -d m1_custom_api_go