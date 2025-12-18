# 專案開發規範與架構說明文件
## 1. 使用的技術與核心架構

### 1.1 使用套件

| 類別          | 技術選型                                                                                                                                                                         | 用途說明                                                          |
|:------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------------------------------------------------------------|
| **核心語言/框架** | Go 1.25<br>Gin<br>Wire                                                                                                                                                       | 開發語言<br>Web 框架<br>依賴注入（編譯期生成）                                 |
| **資料持久化**   | MySQL 8.0+ (GORM)<br>MongoDB 6.0+<br>Redis 7.0+                                                                                                                              | 關聯式資料庫與 ORM<br>文件型資料庫（日誌、事件）<br>快取、會話管理、限流                    |
| **即時通訊**    | MQTT (Eclipse Mosquitto)<br>Socket.IO<br>WebSocket (Gorilla)                                                                                                                 | IoT 設備通訊、輕量級訊息<br>Web 即時通訊<br>雙向即時連接                          |
| **雲端服務**    | AWS S3 SDK v2                                                                                                                                                                | 物件儲存（檔案、圖片）<br>私有雲儲存方案                                        |
| **配置管理**    | Viper                                                                                                                                                                        | 多環境配置管理（支援 TOML）                                              |
| **核心依賴**    | `gorm.io/gorm`<br>`github.com/google/wire`<br>`github.com/spf13/viper`<br>`github.com/go-playground/validator/v10`<br>`github.com/google/uuid`<br>`github.com/tidwall/gjson` | ORM 框架<br>依賴注入代碼生成<br>配置管理<br>結構體驗證<br>UUID 生成<br>高效能 JSON 解析 |
| **日誌**      | Zap                                                                                                                                                                          | 結構化高效能日誌                                                      |
| **部署**      | Docker<br>Docker Compose                                                                                                                                                     | 容器化<br>本地開發與部署編排                                              |

### 1.2 核心架構原則

採用 **Clean Architecture + Domain-Driven Design (DDD)** 混合架構：

- **Clean Architecture**: 確保依賴方向正確，業務邏輯與技術實現解耦
- **DDD**: 以 Bounded Context (BC) 劃分業務領域，明確領域邊界
- **分層設計**: 嚴格的四層架構（Domain、Application、Interface、Infrastructure）
- **依賴注入**: 使用 Wire 在編譯期生成依賴注入代碼，避免運行時反射

---

## 2. 專案目錄結構

```text
project-root/
├── cmd/
│   ├── api/
│       ├── main.go                    # API Server 主程式入口
│       ├── wire.go                    # Wire 依賴注入定義
│       └── wire_gen.go                # Wire 自動生成（不提交）
│
├── configs/
│   ├── dev.toml                       # 開發環境配置
│   ├── prod.toml                      # 生產環境配置
│   └── kernel.go                      # Viper 配置載入與解析
│
├── internal/
│   │
│   ├── common/                        # 【跨領域共用層】
│   │   ├── consts/                    # 全局常數定義
│	│   │   └── role.go                # 角色常數
│	│	│
│   │   ├── middleware/                # Gin 中介層
│   │   │   ├── auth/
│   │   │   │   ├── jwt.go             # JWT Token 驗證
│   │   │   │   └── permission.go      # API 存取權限控制
│   │   │   ├── logging/
│   │   │   │   └── logger.go          # 請求日誌記錄
│   │   │   └── request/
│   │   │       ├── request_id.go      # 請求 ID 追蹤
│   │   │       └── rate_limit.go      # API 限流
│   │   └── util/                      # 通用工具函數
│   │       ├── array.go
│   │       └── string.go
│   │
│   ├── core/                          
│   │   │
│   │   ├── auth/                      # [BC] 身份驗證與授權
│   │   │   ├── domain/
│   │   │   │   ├── entity/            # 領域實體
│   │   │   │   │   ├── user.go
│   │   │   │   │   └── account.go
│   │   │   │   ├── repository/        # 儲存庫介面定義
│   │   │   │   │   └── user_repository.go
│   │   │   │   └── service/           # 領域服務
│   │   │   │       └── auth_domain_service.go
│   │   │   ├── application/           # 應用服務
│   │   │   │   ├── service.go         # AuthService
│   │   │   │   └── dto/               # 請求/響應 DTO
│   │   │   │       ├── request.go
│   │   │   │       └── response.go
│   │   │   └── interface/             # 介面層
│   │   │       ├── http/
│   │   │       │   ├── controller.go  # Auth Controller
│   │   │       │   └── router.go      # 路由註冊
│   │   │       └── middleware/
│   │   │           └── jwt.go         # JWT 中介軟體
│   │   │
│   │   ├── fleet/                     # [BC] 車隊車輛管理
│   │   │   ├── domain/
│   │   │   │   ├── entity/            # Fleet, Vehicle, Driver
│   │   │   │   ├── repository/
│   │   │   │   └── service/
│   │   │   ├── application/
│   │   │   │   ├── service.go
│   │   │   │   └── dto/
│   │   │   └── interface/
│   │   │       └── http/
│   │   │
│   │   ├── order/                     # [BC] 訂單與任務派遣
│   │   │   ├── domain/
│   │   │   │   ├── entity/            # Order, OrderItem, Delivery
│   │   │   │   │   ├── order.go
│   │   │   │   │   └── delivery.go
│   │   │   │   ├── repository/
│   │   │   │   │   ├── order_repository.go
│   │   │   │   │   └── delivery_repository.go
│   │   │   │   ├── service/
│   │   │   │       └── order_service.go
│   │   │   ├── application/
│   │   │   │   ├── service.go
│   │   │   │   └── dto/
│   │   │   └── interface/
│   │   │       └── http/
│	│	│		    └── controller.go
│   │   │
│   │   └── telemetry/                 # [BC] 遙測數據處理
│   │       ├── domain/
│   │       │   ├── entity/            # GPSData, SensorData
│   │       │   │   ├── gps_data.go
│   │       │   │   └── sensor_data.go
│   │       │   ├── repository/
│   │       │   │   └── telemetry_repository.go
│   │       │   └── service/
│   │       │       └── telemetry_domain_service.go
│   │       ├── application/
│   │       │   ├── service.go         # TelemetryService
│   │       │   └── dto/
│   │       └── interface/
│   │           ├── mqtt/              # MQTT Handler (接收車輛遙測數據)
│   │           │   └── handler.go
│   │           ├── websocket/         # WebSocket Handler (即時推送)
│   │           │   └── handler.go
│   │           └── socketio/          # Socket.IO Handler (即時推送)
│   │               └── handler.go
│   │
│   └── infrastructure/                #【基礎設施層】
│       ├── persistence/               # 資料持久化
│       │   ├── mysql/
│       │   │   ├── init.go            # MySQL 連接初始化
│       │   │   ├── record/            # GORM 資料模型（含標籤）
│       │   │   │   ├── user.go
│       │   │   │   ├── order.go
│       │   │   │   └── vehicle.go
│       │   │   └── repository/        # Repository 具體實作
│       │   │       ├── user_mpl.go
│       │   │       ├── order_impl.go
│       │   │       └── vehicle_impl.go
│       │   ├── mongodb/
│       │   │   ├── init.go
│       │   │   ├── record/            # MongoDB 文件模型
│       │   │   │   ├── log.go
│       │   │   │   └── event.go
│       │   │   └── repository/
│       │   │       └── log_impl.go
│       │   └── redis/
│       │       ├── init.go
│       │       └── cache.go           # Redis 快取操作封裝
│       ├── broker/                    # 訊息中介
│       │   ├── mqtt_client.go         # MQTT 客戶端封裝
│       │   ├── publisher.go           # 訊息發佈者
│       │   └── subscriber.go          # 訊息訂閱者
│       ├── socket/                    # 即時通訊
│       │   ├── websocket.go           # WebSocket Server 封裝
│       │   └── socketio.go            # Socket.IO Server 封裝
│       ├── event/                     # 事件基礎設施
│       │   ├── publisher.go           # 領域事件發佈
│       │   └── subscriber.go          # 領域事件訂閱
│       ├── webserver/                 # Web 伺服器
│       │   ├── router.go              # Gin Router 總註冊
│       │   └── health/
│       │       └── health.go          # 健康檢查端點
│       └── storage/                   # 雲端儲存
│           └── s3.go                  # AWS S3 客戶端
│
├── pkg/                               # 【公共套件】
│   ├── logger/
│   │   └── zap.go                     # Zap Logger 封裝
│   ├── crypto/
│   │   ├── hash.go                    # Hash 工具（MD5, SHA256）
│   │   └── password.go                # 密碼加密
│   ├── errors/                        # 統一錯誤處理
│   │   ├── errors.go                  # 自定義錯誤類型
│   │   ├── codes.go                   # 錯誤碼定義
│   │   └── handler.go                 # 錯誤處理器
│   ├── jwt/
│   │   └── jwt.go                     # JWT Token 生成與驗證
│   └── tools/
│       ├── gjson.go                   # GJSON 工具封裝
│       ├── uuid.go                    # UUID 生成工具
│       └── validator.go               # 驗證工具封裝
│
├── deployments/                       # 部署配置
│   ├── docker-compose.yml
│   ├── docker-compose.dev.yml
│   ├── docker-compose.prod.yml
│   ├── Dockerfile
│   └── mosquitto/
│       └── mosquitto.conf             # MQTT Broker 配置
│
├── docs/                              # 文件
│   ├── architecture.md                # 架構說明
│   ├── api/                           # API 文件
│   └── development.md                 # 開發指南
│
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

## 3. Clean Architecture + DDD 架構說明

### 3.1 核心設計原則

#### 依賴規則 (Dependency Rule)

**依賴只能由外向內，內層不能依賴外層**

```text
┌─────────────────────────────────────────┐
│  Interface Layer (最外層)               │ ← 依賴 Application
│  Controllers, Handlers, Routers         │
├─────────────────────────────────────────┤
│  Application Layer                      │ ← 依賴 Domain
│  Use Cases, Service, DTOs               │
├─────────────────────────────────────────┤
│  Domain Layer (核心，最內層)            │ ← 不依賴任何層
│  Entities, Repository Interfaces        │
├─────────────────────────────────────────┤
│  Infrastructure Layer (最外層)          │ ← 依賴 Domain（實作介面）
│  DB, MQTT, S3, External Services        │
└─────────────────────────────────────────┘
```

**依賴方向檢查表**：

| 從哪一層               | 可以依賴                                               | 禁止依賴                                             |
|--------------------|----------------------------------------------------|--------------------------------------------------|
| **Domain**         | 無                                                  | ❌ Application<br>❌ Interface<br>❌ Infrastructure |
| **Application**    | ✅ Domain（Entity、Repository Interface）              | ❌ Interface<br>❌ Infrastructure（具體實作）            |
| **Interface**      | ✅ Application<br>✅ Domain                          | ❌ Infrastructure（具體實作）                           |
| **Infrastructure** | ✅ Domain（實作 Repository Interface）<br>✅ Application | ❌ Interface                                      |

#### 實體獨立性 (Entity Independence)

**Domain Entity 必須完全純淨，不包含任何技術標籤**

```go
// ✅ 正確：Domain Entity（無任何標籤）
// internal/core/order/domain/entity/order.go
package entity

import (
	"time"
	"github.com/google/uuid" // ✅ 純函式庫 OK
)

type Order struct {
	ID          string // ✅ 無標籤
	CustomerID  string
	Items       []OrderItem
	Status      OrderStatus
	TotalAmount float64
	CreatedAt   time.Time
}

// 業務方法
func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (o *Order) CanCancel() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}

func (o *Order) Cancel() error {
	if !o.CanCancel() {
		return errors.New("cannot cancel order in current status")
	}
	o.Status = OrderStatusCancelled
	return nil
}
```

```go
// ❌ 錯誤：Entity 包含技術標籤
package entity

type Order struct {
	ID     string `gorm:"primaryKey" json:"id"` // ❌ 不可以！
	Status string `gorm:"index" json:"status"`
}
```

#### 資料模型分離

**三種模型嚴格分離，各司其職**

```
┌─────────────────────────────────────────────────────┐
│ Request/Response DTO (interface/application layer)  │
│ - 包含 json, binding 標籤                            │
│ - 用於 API 請求/響應                                  │
└─────────────────────────────────────────────────────┘
                        ↓ 轉換
┌─────────────────────────────────────────────────────┐
│ Domain Entity (domain layer)                        │
│ - 無任何標籤                                         │
│ - 包含業務邏輯方法                                    │
└─────────────────────────────────────────────────────┘
                        ↓ 轉換
┌─────────────────────────────────────────────────────┐
│ Database Record (infrastructure/persistence)        │
│ - 包含 gorm, bson 標籤                               │
│ - 對應資料庫表結構                                    │
└─────────────────────────────────────────────────────┘
```

**範例**：

```go
// 1. Request DTO (application/dto/request.go)
type CreateOrderRequest struct {
CustomerID string   `json:"customer_id" binding:"required"`
Items      []string `json:"items" binding:"required"`
}

// 2. Domain Entity (domain/entity/order.go)
type Order struct {
ID         string // 無標籤
CustomerID string
Items      []OrderItem
Status     OrderStatus
}

// 3. Database Record (infrastructure/persistence/mysql/record/order.go)
type Order struct {
ID         string    `gorm:"primaryKey"`
CustomerID string    `gorm:"index"`
Status     string    `gorm:"index"`
CreatedAt  time.Time `gorm:"autoCreateTime"`
}
```

#### Repository Pattern（儲存庫模式）

**Domain 定義介面，Infrastructure 實作介面**

```go
// Domain 層：定義介面（不實作）
// internal/core/order/domain/repository/order_repository.go
package repository

import (
	"context"
	"your-project/internal/core/order/domain/entity"
)

type IOrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	FindByID(ctx context.Context, id string) (*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id string) error
	FindByCustomerID(ctx context.Context, customerID string) ([]*entity.Order, error)
}
```

```go
// Infrastructure 層：實作介面
// internal/infrastructure/persistence/mysql/repository/order_repository_impl.go
package repository

import (
	"context"
	"gorm.io/gorm"
	"your-project/internal/core/order/domain/entity"
	"your-project/internal/core/order/domain/repository"
	"your-project/internal/infrastructure/persistence/mysql/record"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

// 確保實作介面
var _ repository.IOrderRepository = (*OrderRepositoryImpl)(nil)

func NewOrderRepository(db *gorm.DB) repository.IOrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(ctx context.Context, order *entity.Order) error {
	rec := toRecord(order) // Entity → Record 轉換
	return r.db.WithContext(ctx).Create(rec).Error
}

func (r *OrderRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	var rec record.Order
	if err := r.db.WithContext(ctx).First(&rec, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toEntity(&rec), nil // Record → Entity 轉換
}

// Entity ↔ Record 轉換函數
func toRecord(e *entity.Order) *record.Order {
	return &record.Order{
		ID:         e.ID,
		CustomerID: e.CustomerID,
		Status:     string(e.Status),
		CreatedAt:  e.CreatedAt,
	}
}

func toEntity(r *record.Order) *entity.Order {
	return &entity.Order{
		ID:         r.ID,
		CustomerID: r.CustomerID,
		Status:     entity.OrderStatus(r.Status),
		CreatedAt:  r.CreatedAt,
	}
}
```

#### **層級職責分離**

**每一層有明確且單一的職責**

| 層級                 | 職責                                   | 禁止事項                                         |
|--------------------|--------------------------------------|----------------------------------------------|
| **Domain**         | 定義業務規則、實體行為、儲存庫介面                    | ❌ 不可包含技術實現<br>❌ 不可依賴外層<br>❌ 不可有 GORM/Gin 等標籤 |
| **Application**    | 編排用例、協調多個 Domain Service、處理事務        | ❌ 不可包含業務規則<br>❌ 不可直接操作資料庫<br>❌ 不可依賴具體實作      |
| **Interface**      | 處理 HTTP/MQTT/WebSocket 請求、資料驗證、響應格式化 | ❌ 不可包含業務邏輯<br>❌ 不可直接訪問 Repository            |
| **Infrastructure** | 實作 Repository、連接外部服務（DB、MQTT、S3）     | ❌ 不可包含業務邏輯<br>❌ 不可被 Domain 依賴                |

### 3.2 DDD 核心概念

#### **Bounded Context (界限上下文)**

本專案按業務領域劃分為多個 BC，每個 BC 是一個獨立的業務模組：

- **`auth`**: 身份驗證與授權（登入、註冊、權限管理）
- **`fleet`**: 車隊管理（車輛、班表）
- **`order`**: 訂單管理（訂單創建、派遣、追蹤）
- **`telemetry`**: 遙測數據（GPS、感測器數據處理）

**BC 之間透過以下方式通訊**：

1. **應用層直接調用**（同步）
2. **領域事件**（非同步，選配）
3. **MQTT 訊息**（鬆耦合）

#### **Domain Service vs Application Service**

```
Domain Service (domain/service/)
├─ 職責：處理單一 BC 內跨實體的業務邏輯
├─ 範例：計算訂單優先級、驗證車輛可用性
└─ 特點：純業務邏輯，不依賴外部服務

Application Service (application/service.go)
├─ 職責：編排用例，協調多個 Domain Service 和 Repository
├─ 範例：建立訂單並分配司機（跨 Order 和 Fleet BC）
└─ 特點：可處理事務、調用外部服務（MQTT、S3）
```

**範例**：

```go
// Domain Service：訂單優先級計算（純業務邏輯）
// internal/core/order/domain/service/order_domain_service.go
package service

import "your-project/internal/core/order/domain/entity"

type OrderDomainService struct {
	// 不依賴任何外部服務
}

func (s *OrderDomainService) CalculatePriority(order *entity.Order) int {
	priority := 0
	if order.IsUrgent {
		priority += 10
	}
	if order.TotalAmount > 1000 {
		priority += 5
	}
	return priority
}
```

```go
// Application Service：建立訂單並分配司機（編排用例）
// internal/core/order/application/service.go
package application

import (
	"context"
	orderDomain "your-project/internal/core/order/domain/repository"
	orderService "your-project/internal/core/order/domain/service"
	fleetDomain "your-project/internal/core/fleet/domain/repository"
	"your-project/pkg/logger"
)

type OrderService struct {
	orderRepo      orderDomain.IOrderRepository
	driverRepo     fleetDomain.IDriverRepository
	orderDomainSvc *orderService.OrderDomainService
	mqttPublisher  MQTTPublisher
}

func (s *OrderService) CreateOrderWithAssignment(ctx context.Context, dto CreateOrderDTO) error {
	// 1. 建立 Domain Entity
	order := entity.NewOrder(dto.CustomerID, dto.Items)

	// 2. 呼叫 Domain Service 計算優先級
	order.Priority = s.orderDomainSvc.CalculatePriority(order)

	// 3. 開始事務
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 4. 儲存訂單
	if err := s.orderRepo.CreateWithTx(ctx, tx, order); err != nil {
		tx.Rollback()
		return err
	}

	// 5. 尋找可用司機（跨 BC 調用）
	driver, err := s.driverRepo.FindAvailableDriver(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 6. 分配司機
	order.AssignDriver(driver.ID)
	if err := s.orderRepo.Update(ctx, order); err != nil {
		tx.Rollback()
		return err
	}

	// 7. 提交事務
	if err := tx.Commit(); err != nil {
		return err
	}

	// 8. 發送 MQTT 通知（基礎設施服務）
	s.mqttPublisher.Publish("order/assigned", order.ID)

	logger.Info("Order created and assigned", "orderID", order.ID, "driverID", driver.ID)
	return nil
}
```

---

## 附錄：關鍵規則速查表

### ✅ DO（應該做的）

1. **Domain Entity**
    - 包含業務邏輯方法
    - 使用值物件（Value Object）
    - 保持純淨（無標籤）

2. **Repository**
    - 在 Domain 定義介面
    - 在 Infrastructure 實作介面
    - 使用 `context.Context`

3. **Application Service**
    - 編排用例
    - 處理事務
    - 協調多個 Domain Service

### ❌ DON'T（禁止做的）

1. **Domain Entity**
    - 不可包含 `gorm`、`json`、`bson` 標籤
    - 不可依賴框架（Gin、GORM）
    - 不可依賴外層（Application、Infrastructure）

2. **Application Service**
    - 不可直接操作資料庫（必須透過 Repository）
    - 不可包含業務規則（應在 Domain）
    - 不可依賴具體實作（只依賴介面）

3. **Controller**
    - 不可包含業務邏輯
    - 不可直接訪問 Repository
    - 不可處理事務

4. **全域**
    - 避免循環依賴
    - 避免全域變數（使用依賴注入）