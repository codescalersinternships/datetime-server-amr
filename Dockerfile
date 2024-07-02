# BUild Stage
FROM golang:1.22.4-alpine AS BuildStage

WORKDIR /datetime-server

COPY go.mod .

RUN go mod download

COPY . .

EXPOSE 8000

RUN go build -o /dateTimeServer main.go

# Deploy Stage
FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /dateTimeServer /dateTimeServer

EXPOSE 8000

ENTRYPOINT ["/dateTimeServer"]