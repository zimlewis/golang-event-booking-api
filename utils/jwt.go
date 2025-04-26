package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "aaaaaa"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected token signing type")
		}
		
		return []byte(SECRET_KEY), nil	
	})

	if err != nil {
		return -1, err
	}

	if !parsed.Valid {
		return -1, errors.New("parse failed")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid token claims")
	}

	var exp int64;

	switch v := claims["exp"].(type) {
	case float64:
		exp = int64(v)
	case float32:
		exp = int64(v)
	case int:
		exp = int64(v)
	case int64:
		exp = v
	default:
		return -1, errors.New("invalid exp type")
	}

	var uid int64;

	switch v := claims["userId"].(type) {
	case float64:
		uid = int64(v)
	case float32:
		uid = int64(v)
	case int:
		uid = int64(v)
	case int64:
		uid = v
	default:
		return -1, errors.New("invalid user id type")
	}



	if time.Unix(exp, 0).Before(time.Now()) {
		return -1, errors.New("expired token")
	}

	return uid, nil
}