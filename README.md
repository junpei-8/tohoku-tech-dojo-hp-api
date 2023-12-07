# 東北 TECH 道場 HP API

<br>

## セットアップ

```
go mod download
```

<br>

## 開発

### DB

```
docker compose up db
```

### DB GUI Manager

```
docker compose up db-gui
```

### API

```
docker compose up api
```

<br>

## デプロイ

### API

```
make deploy-api
```

### DB

```
make deploy-db
```
