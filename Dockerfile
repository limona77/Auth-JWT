FROM golang:latest

WORKDIR /back

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go


CMD ["/back/main"]