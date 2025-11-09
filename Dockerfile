FROM golang:1.24.5-alpine AS builder

RUN apk add --no-cache git gcc musl-dev sqlite-dev

RUN git clone https://github.com/pressly/goose /goose

WORKDIR /goose

RUN go mod tidy

RUN go build \
  -ldflags="-s -w" \
  -tags='no_postgres no_clickhouse no_mssql no_mysql' \
  -o goose ./cmd/goose
RUN ls -la

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
ENV GOOS=linux

RUN mkdir -p bin
RUN go build -a -installsuffix cgo -o bin/server ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite

WORKDIR /root/

COPY --from=builder /app/bin/server .
COPY --from=builder /app/migrations migrations
COPY --from=builder /goose/goose .

RUN mkdir -p data
ENV DB_PATH=data/database.db
COPY /app/data/database.db data/database.db
VOLUME [ "/data" ]

ENV PORT=:5000

EXPOSE 5000

CMD ["./server"]
