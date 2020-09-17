package utils

import (
	"LearnLogin/config"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//token设置
var JwtKey = []byte("www.zero.com")

type Client struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//密码加密
func EncryptionPwd(password string) string {
	salt := time.Now().Unix()
	md5 := md5.New()
	md5.Write([]byte(password))
	md5.Write([]byte(string(salt)))
	st := md5.Sum(nil)
	pwd := hex.EncodeToString(st)
	return pwd
}

//生成JWT
func GetToken(username string, role string) (string, error) {
	expiresTime := time.Now().Add(10 * time.Hour)
	claime := &Client{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			Issuer:    "zero",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claime)
	return token.SignedString(JwtKey)
}

//解析JWT
func CheckToken(tokenString string) (*Client, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Client{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	//验证token
	if claims, ok := token.Claims.(*Client); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效令牌！")
}

//Jwt中间件
func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    config.NilToken,
				"message": "token为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code":    config.NilToken,
				"message": "token格式错误",
			})
			c.Abort()
			return
		}
		mc, err := CheckToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("role", mc.Role)
		c.Next()
	}
}
