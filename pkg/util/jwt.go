/**
 * @Author : jinchunguang
 * @Date : 19-10-31 下午3:45
 * @Project : bingo-micro
 */
package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "bingo-micro/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour)

    claims := Claims{
        username,
        password,
        jwt.StandardClaims {
            ExpiresAt : expireTime.Unix(),
            Issuer : "bingo-micro",
        },
    }
    // 签名方案
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // 生成签名字符串，再用于获取完整、已签名的token
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {

    // 解码和校验
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    // 验证
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}