# Soybean Admin Go

Go语言实现的Soybean Admin后端服务。

## 数据库支持

本项目支持以下数据库：

- PostgreSQL (默认)
- MySQL ([使用指南](./docs/mysql_usage.md))

## 配置

复制 `config.example.json` 文件并重命名为 `config.json`，然后根据需要修改配置。

```bash
cp config.example.json config.json
```

## 运行

### 开发环境

```bash
go run .
```

### 生产环境

```bash
go build -tags=prod -o server .
./server
```

### Docker

```bash
docker build -t soybean-admin-go .
docker run -d -p 8080:8080 -v /path/to/config/config.json:/app/config/config.json soybean-admin-go
```

## API文档

启动服务后访问 `/swagger/index.html` 查看API文档。 