FROM golang:1.23.4-bullseye
WORKDIR /app/

COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o app main.go

CMD ["./app"]