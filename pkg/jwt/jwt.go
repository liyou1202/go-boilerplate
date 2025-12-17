package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT 自定義聲明
type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	RoleId   string `json:"roleId"`
	jwt.RegisteredClaims
}

// Generate 生成 JWT Token
func Generate(userID, username, roleID, secret string, expireHours int) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RoleId:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Parse 解析 JWT Token
func Parse(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

// Verify 驗證 JWT Token 是否有效
func Verify(tokenString, secret string) bool {
	_, err := Parse(tokenString, secret)
	return err == nil
}
