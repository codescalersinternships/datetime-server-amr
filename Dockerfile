# Build Stage 1 for HTTP Server
FROM golang:1.22.4-alpine AS img1

WORKDIR /datetime-server-http

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY http .

EXPOSE 8000

RUN go build -o /dateTimeServerhttp cmd/main.go

# Deploy Stage 1 for HTTP Server
FROM alpine:latest AS deploy1

WORKDIR /

COPY --from=img1 /dateTimeServerhttp /dateTimeServerhttp

EXPOSE 8000

ENTRYPOINT ["/dateTimeServerhttp"]

# Build Stage 2 for Gin Server
FROM golang:1.22.4-alpine AS img2

WORKDIR /datetime-server-gin

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY gin .

EXPOSE 8000

RUN go build -o /dateTimeServergin cmd/main.go

# Deploy Stage 2 for Gin Server
FROM alpine:latest AS deploy2

WORKDIR /

COPY --from=img2 /dateTimeServergin /dateTimeServergin

EXPOSE 8000

ENTRYPOINT ["/dateTimeServergin"]
