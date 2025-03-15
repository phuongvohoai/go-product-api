FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /product-api .

EXPOSE 8080

CMD ["/product-api"]