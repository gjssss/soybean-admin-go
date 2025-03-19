# MySQL 数据库支持指南

本项目现已支持MySQL数据库，您可以按照以下步骤使用MySQL替代PostgreSQL。

## 配置方法

1. 复制配置文件模板：

```bash
cp config.mysql.example.json config.json
```

2. 编辑 `config.json` 文件，填入您的MySQL数据库信息：

```json
"db": {
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "user": "your_mysql_username",
    "password": "your_mysql_password",
    "name": "soybean_admin",
    "timezone": "Asia/Shanghai"
}
```

## MySQL 数据库配置说明

| 字段 | 说明 |
|------|------|
| type | 数据库类型，填写 `mysql` |
| host | MySQL 服务器地址 |
| port | MySQL 端口，默认 3306 |
| user | 数据库用户名 |
| password | 数据库密码 |
| name | 数据库名称 |
| sslmode | SSL 模式，MySQL 可以留空 |
| timezone | 时区，建议设置为 `Asia/Shanghai` 或 `Local` |

## 预先准备

在使用之前，请确保已创建对应的MySQL数据库：

```sql
CREATE DATABASE soybean_admin CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## Docker 环境使用 MySQL

如果您使用Docker部署，可以通过挂载卷的方式传入MySQL配置：

```bash
docker run -d \
  --name soybean-admin \
  -p 8080:8080 \
  -v /path/to/mysql_config.json:/app/config/config.json \
  soybean-admin-go
```

或者使用Docker Compose：

```yaml
version: '3'

services:
  mysql:
    image: mysql:8.0
    container_name: soybean-admin-mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: your_password
      MYSQL_DATABASE: soybean_admin
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"
    
  soybean-admin:
    build: .
    container_name: soybean-admin
    depends_on:
      - mysql
    ports:
      - "8080:8080"
    volumes:
      - ./config.mysql.json:/app/config/config.json
    restart: unless-stopped
    environment:
      - GIN_MODE=release

volumes:
  mysql_data:
```

## 注意事项

1. MySQL 与 PostgreSQL 之间存在一些语法和功能差异，如果您编写了自定义 SQL 查询，可能需要针对 MySQL 进行调整
2. 第一次运行时，系统会自动创建所需的表结构
3. 如果遇到编码问题，请确保 MySQL 配置使用 utf8mb4 字符集 