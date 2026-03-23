# ===== Fase 1: BUILD =====
FROM golang:1.25-alpine AS build
RUN apk add --no-cache build-base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
ENV APP_PORT=8080
ENV DB_PATH=/data/app.db
RUN mkdir -p /data
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/api ./cmd/api

# ===== Fase 2: RUNTIME =====
FROM alpine:3.20
RUN apk add --no-cache sqlite-libs ca-certificates
WORKDIR /app
COPY --from=build /app/bin/api /app/api
RUN mkdir -p /data
EXPOSE 8080
ENV APP_PORT=8080
ENV DB_PATH=/data/app.db
CMD ["/app/api"]
