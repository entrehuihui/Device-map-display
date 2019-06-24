package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// exp JWT过期时间 默认6小时=21600秒
var exp = 21600
var hmacSampleSecret = []byte("mapPlatformtokenhmacSampleSecret")

// SetExp 设置过期和解密字符串时间
func SetExp(sec int, hmacSample string) {
	if sec > 0 {
		exp = sec
	}
	if hmacSample != "" {
		hmacSampleSecret = []byte(hmacSample)
	}
}

// NewJWT .
func NewJWT(id, permisson int) string {
	now := int(time.Now().Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.Itoa(id),
		"yi":  strconv.Itoa(now<<6 | permisson<<4 | (now & 0xf)),
		"exp": time.Now().Add(time.Second * time.Duration(exp)).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		log.Println(err)
		return ""
	}
	return tokenString
}

// ParseJWT .
func ParseJWT(tokenString string) ([]int, error) {
	// if tokenString == "" {
	// 	return "", errors.New()
	// }
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("not authorized")
	}
	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return nil, err
	}
	yi, err := strconv.Atoi(claims["yi"].(string))
	if err != nil {
		return nil, err
	}
	yi = yi >> 4 & 0x03
	return []int{id, yi}, nil
}
