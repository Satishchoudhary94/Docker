FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o calc

ENTRYPOINT ["./calc"]
