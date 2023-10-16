#build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o cmd cmd/main.go

#run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/cmd .
COPY app.env .
COPY internal/repository/pgstore/migration ./internal/repository/pgstore/migration

EXPOSE 8080
CMD [ "/app/cmd" ]