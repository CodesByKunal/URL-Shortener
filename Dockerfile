FROM golang:latest

WORKDIR /src

COPY . .

RUN go mod tidy

ENV HOST=0.0.0.0
ENV PORT=3000

EXPOSE 3000

CMD ["go","run","cmd/urlshortener/main.go"]