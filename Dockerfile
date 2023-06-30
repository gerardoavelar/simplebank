#Build stage
FROM golang:1.19.10-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Final run
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]