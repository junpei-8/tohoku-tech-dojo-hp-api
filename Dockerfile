# DB のマイグレーションを実行
FROM golang:1.19.3-alpine as migrate-db

WORKDIR /app

COPY . .

CMD ["go", "run", "./cmd/migrate_db"]


# 開発用 API サーバーの立ち上げ
FROM golang:1.19.3-alpine as dev-api

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", "./.air.toml"]


# 本番環境用の API サーバーを実行するための exe ファイルの作成
FROM golang:1.19.3-alpine as api-builder

WORKDIR /app

COPY . .

RUN go generate ./...

RUN go build -o ./dist cmd/api/main.go


# 本番環境用の API サーバーで生成された exe ファイルの実行
FROM alpine as api-runner

COPY --from=api-builder ./dist .

CMD /app $PORT
