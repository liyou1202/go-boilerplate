# API Backend

基於 Go + Clean Architecture + DDD 的後端 API 服務初始框架

- 核心框架: Go 1.25 · Gin · Wire · Viper
- 資料庫: MySQL (GORM) · MongoDB · Redis
- 即時通訊: MQTT (Mosquitto) · WebSocket · Socket.IO
- 其他: Zap (日誌) · JWT · AWS S3

## 快速啟動

### 本地開發（推薦）

```bash
# 1. 啟動基礎設施（MySQL, MongoDB, Redis, MQTT）
make loc-infra

# 2. 運行 API
make loc-backend
```

服務啟動後訪問：`http://localhost:8080/health`

### 完整環境（全部用 Docker）

```bash
make loc-all
```

### 停止服務

```bash
make down
```

## 環境說明

專案支援三種環境配置：

| 環境 | 配置文件 | 啟動命令 |
|------|---------|---------|
| **本地** (loc) | `configs/loc.toml` | `make loc-backend` / `make loc-infra` / `make loc-all` |
| **開發** (dev) | `configs/dev.toml` | `make dev-backend` / `make dev-infra` / `make dev-all` |
| **生產** (prod) | `configs/prod.toml` | `make prod-backend` / `make prod-infra` / `make prod-all` |

## 環境變量（可選）

```bash
cd deployments
cp .env.example .env
# 編輯 .env 設置密碼（不設置則使用默認值）
```

## 常用命令

```bash
make help      # 查看所有命令
make wire      # 生成依賴注入代碼
make test      # 運行測試
make logs      # 查看服務日誌
make clean     # 清理編譯產物
```

## 目錄結構

```
├── cmd/api/              # 應用入口與 Wire 配置
├── configs/              # 環境配置文件
├── internal/
│   ├── common/           # 共用層（中介層、工具）
│   ├── core/             # 業務核心（auth, fleet, order, telemetry）
│   └── infrastructure/   # 基礎設施（資料庫、MQTT、S3）
├── pkg/                  # 公共套件（logger, crypto, jwt, errors）
├── deployments/          # Docker 配置
├── docs/                 # 文檔
└── Makefile
```

## 架構說明

採用 Clean Architecture + DDD 四層架構：
- **Domain Layer**: 領域實體與業務邏輯（純淨，無標籤）
- **Application Layer**: 用例編排與事務處理
- **Interface Layer**: HTTP/MQTT/WebSocket 處理
- **Infrastructure Layer**: 資料庫與外部服務實作

詳細說明參考 [docs/develop_rule.md](docs/develop_rule.md)

## 服務端口

- API Server: `8080`
- MySQL: `3306`
- MongoDB: `27017`
- Redis: `6379`
- MQTT: `1883`
