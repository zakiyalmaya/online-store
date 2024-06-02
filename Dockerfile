FROM golang:1.22-alpine

ENV GO111MODULE=on

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

COPY . .

RUN go build -o main .

# Expose port
EXPOSE 3000

CMD ["./main"]
