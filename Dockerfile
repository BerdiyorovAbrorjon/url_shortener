#build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

#run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY internal/repository/pgstore/migration ./internal/repository/pgstore/migration

EXPOSE 8080
CMD [ "/app/main" ]