FROM golang:1.22-alpine

ENV GO111MODULE=on

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o online_store_app

RUN echo "Files in /app:" && ls -la /app
RUN echo "Environment variables:" && env

RUN chmod +x /app/online_store_app

EXPOSE 3000

CMD ["./online_store_app"]
