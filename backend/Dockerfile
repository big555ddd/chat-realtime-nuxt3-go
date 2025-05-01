FROM golang:alpine as builder
RUN apk --no-cache add build-base tzdata ca-certificates
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum /app/
RUN go mod download

# Copy the source code
COPY . /app

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=go_json -o /app/dist/app .
RUN mkdir -p /app/storage

FROM gcr.io/distroless/static as serve
WORKDIR /app

# Copy the built application and storage directory
COPY --from=builder /app/dist/app /app/dist/app
COPY --from=builder /app/storage /app/storage

# Set the PORT environment variable
ENV PORT 8080

ENTRYPOINT [ "/app/dist/app", "http" ]