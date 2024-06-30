FROM golang:1.22.0

WORKDIR /app

RUN go mod init teste

COPY . .

RUN go build -o math

CMD ["./math"]