FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/server cmd/server/main.go
RUN go build -o /app/refresh_news cmd/refresh_news/main.go


FROM alpine:latest

COPY --from=builder /app/server /server

COPY --from=builder /app/refresh_news /refresh_news

CMD ["./server"]