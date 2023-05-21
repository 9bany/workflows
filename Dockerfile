FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod tidy
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/main .

COPY start.sh .
RUN ["chmod", "+x", "/app/start.sh"]

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]