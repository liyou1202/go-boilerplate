# 部署配置說明

## 快速開始

### 1. 設置環境變量（可選）

```bash
# 複製範例文件
cp .env.example .env

# 編輯 .env，設置您的密碼
# 如果不設置，將使用默認密碼（適合開發環境）
```

### 2. 啟動服務

查看所有可用命令：
```bash
make help
```

## 使用場景

### 場景 1：機器 A（基礎建設） + 機器 B（本地開發）

**機器 A（192.168.1.100）：**
```bash
make loc-infra
# 啟動：MySQL, MongoDB, Redis, MQTT
```

**機器 B（開發機）：**
```bash
# 1. 修改 configs/loc.toml，把 localhost 改成 192.168.1.100
# 2. 運行 API
make loc-backend
```

### 場景 2：完整本地開發環境

```bash
make loc-all
# 啟動所有服務（基礎設施 + API）
```

### 場景 3：部署到開發服務器

```bash
make dev-all
# 使用開發環境配置啟動所有服務
```

### 場景 4：部署到生產環境

```bash
# 1. 設置 .env 文件（必須）
cp .env.example .env
vim .env  # 設置強密碼

# 2. 啟動所有服務
make prod-all
```

## Profiles 說明

服務分為兩個 profile：

- **infra**: 基礎設施（MySQL, MongoDB, Redis, MQTT）
- **backend**: API 服務器

| 命令 | 啟動的服務 |
|------|-----------|
| `make loc-infra` | 只啟動基礎設施 |
| `make loc-backend` | 本地運行 API（go run，不用 Docker） |
| `make loc-all` | 啟動基礎設施 + API（都用 Docker） |

## 環境變量

`.env` 文件（gitignore，不會提交）控制：
- 數據庫密碼
- 端口號
- 環境設置（ENV, GIN_MODE）

如果沒有 `.env`，將使用默認值（見 `.env.example`）。

## 停止服務

```bash
make down
```

## 查看日誌

```bash
make logs
```
