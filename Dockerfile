FROM golang:1.18.2 as development
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]

# ---------------------------------------------------

# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app ./cmd

# ---------------------------------------------------

FROM debian:bullseye-slim as production

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]
