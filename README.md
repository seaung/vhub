# VHub

VHub是一个使用Golang开发的Web漏洞靶场环境，提供多种常见Web安全漏洞的演示和学习环境。通过实践来提升Web安全测试能力。

## 功能特性

- 支持多种常见Web安全漏洞环境：
  - 反射型XSS漏洞
  - 存储型XSS漏洞
  - DOM型XSS漏洞
  - XXE外部实体注入漏洞
  - SQL注入漏洞
  - SSRF服务端请求伪造漏洞
  - 命令执行漏洞

## 环境要求

- Go 1.16+
- MySQL 5.7+

## 快速开始

### 使用Docker Compose（推荐）

1. 克隆项目
```bash
git clone https://github.com/seaung/vhub.git
cd vhub
```

2. 启动服务
```bash
docker-compose up -d
```

3. 访问应用
```
打开浏览器访问：http://localhost:8080
```

### 手动部署

1. 克隆项目
```bash
git clone https://github.com/seaung/vhub.git
cd vhub
```

2. 修改配置文件
```bash
cp config/config.yml.example config/config.yml
# 编辑config.yml，配置数据库连接信息
```

3. 启动应用
```bash
go mod tidy
go run main.go
```

4. 访问应用
```
打开浏览器访问：http://localhost:8080
```

## 注意事项

- 本项目仅用于学习和测试目的
- 请勿在生产环境使用
- 使用过程中遵守相关法律法规
