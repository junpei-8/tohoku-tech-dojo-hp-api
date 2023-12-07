# Cloud SQL へ DB をマイグレーションするための Dockerfile


## 本番環境の DB へマイグレーションを実行（開発中）
FROM arigaio/atlas:latest as migration

COPY ./migrations ./migrations

CMD []
