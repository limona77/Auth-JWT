FROM golang:1.22.0-alpine

WORKDIR /back

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go
RUN go mod tidy

CMD ["/back/main"]