package errors

// 通用錯誤碼
const (
	CodeSuccess          = 0
	CodeInternalError    = 10001
	CodeInvalidParams    = 10002
	CodeNotFound         = 10003
	CodeUnauthorized     = 10004
	CodeForbidden        = 10005
	CodeAlreadyExists    = 10006
	CodeDatabaseError    = 10007
	CodeCacheError       = 10008
	CodeExternalAPIError = 10009
)

// 認證相關錯誤碼
const (
	CodeAuthInvalidToken  = 20001
	CodeAuthExpiredToken  = 20002
	CodeAuthInvalidUser   = 20003
	CodeAuthWrongPassword = 20004
	CodeAuthUserExists    = 20005
)

// 訂單相關錯誤碼
const (
	CodeOrderNotFound       = 30001
	CodeOrderInvalidStatus  = 30002
	CodeOrderAlreadyPaid    = 30003
	CodeOrderCannotCancel   = 30004
)

// 車隊相關錯誤碼
const (
	CodeFleetNotFound         = 40001
	CodeVehicleNotFound       = 40002
	CodeDriverNotFound        = 40003
	CodeVehicleNotAvailable   = 40004
	CodeDriverNotAvailable    = 40005
)

// 遙測相關錯誤碼
const (
	CodeTelemetryInvalidData = 50001
	CodeTelemetryDeviceOffline = 50002
)

// 錯誤訊息對應表
var messages = map[int]string{
	CodeSuccess:               "success",
	CodeInternalError:         "internal server error",
	CodeInvalidParams:         "invalid parameters",
	CodeNotFound:              "resource not found",
	CodeUnauthorized:          "unauthorized",
	CodeForbidden:             "forbidden",
	CodeAlreadyExists:         "resource already exists",
	CodeDatabaseError:         "database error",
	CodeCacheError:            "cache error",
	CodeExternalAPIError:      "external api error",
	CodeAuthInvalidToken:      "invalid token",
	CodeAuthExpiredToken:      "token expired",
	CodeAuthInvalidUser:       "invalid user",
	CodeAuthWrongPassword:     "wrong password",
	CodeAuthUserExists:        "user already exists",
	CodeOrderNotFound:         "order not found",
	CodeOrderInvalidStatus:    "invalid order status",
	CodeOrderAlreadyPaid:      "order already paid",
	CodeOrderCannotCancel:     "order cannot be cancelled",
	CodeFleetNotFound:         "fleet not found",
	CodeVehicleNotFound:       "vehicle not found",
	CodeDriverNotFound:        "driver not found",
	CodeVehicleNotAvailable:   "vehicle not available",
	CodeDriverNotAvailable:    "driver not available",
	CodeTelemetryInvalidData:  "invalid telemetry data",
	CodeTelemetryDeviceOffline: "device offline",
}

// GetMessage 根據錯誤碼取得錯誤訊息
func GetMessage(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "unknown error"
}
