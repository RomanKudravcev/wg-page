#build phase
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ensure /db directory is created in the /app directory
RUN mkdir -p ./db

# Copy the source code, including the database.sql file
COPY ./db/database.sql ./db/
RUN CGO_ENABLED=0 GOOS=linux go build -o /backendApp

# final phase
FROM alpine:latest

WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /backendApp .

# Copy the database.sql file from the builder stage
COPY --from=builder /app/db/database.sql /app/db/

EXPOSE 8080

CMD ["./backendApp"]
