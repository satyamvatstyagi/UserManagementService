# Multi-stage build to compile the Go application
FROM golang:1.20 AS builduserserv

WORKDIR /application

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Set the Gin mode to debug (or release, or test)
ENV GIN_MODE=debug

# Set the default value for the database environment variables
ENV DATABASE_HOST=localhost
ENV DATABASE_PORT=5432
ENV DATABASE_USER=postgres
ENV DATABASE_PASSWORD=nagarro
ENV DATABASE_NAME=mtn

#Test go application
RUN go test ./...

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.57.1

# Run linting
RUN golangci-lint run -v

RUN CGO_ENABLED=0 GOOS=linux go build -o /user-management-serv cmd/server/main.go

RUN rm -rf /application/*

# Build a minimal image for running the application
FROM alpine:3.11.3

RUN addgroup -g 1000 noroot
RUN adduser -u 1000 -G noroot -h /home/noroot -D noroot
RUN mkdir /home/noroot/app
WORKDIR /home/noroot/app

COPY --from=builduserserv /user-management-serv /home/noroot/app/

ARG SERVICE_PORT
EXPOSE $SERVICE_PORT

# Set the default value for the SERVICE_PORT environment variable
ENV SERVICE_PORT=$SERVICE_PORT

ENTRYPOINT ["./user-management-serv"]