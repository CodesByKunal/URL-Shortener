FROM golang:latest

WORKDIR /src

COPY go.mod .

RUN go mod tidy

EXPOSE 3000

CMD ["go","run","cmd/urlshortener/main.go"]