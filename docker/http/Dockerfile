FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o main ./cmd/mutual-http

CMD ["/app/main"]
