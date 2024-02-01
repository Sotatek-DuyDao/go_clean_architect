package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
}

func GenerateJWT(userId uint, email, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		ID:       userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString("DUYDAO")
	fmt.Printf("TOKEN : %s\n", tokenString)
	return tokenString, err
}

func ValidateToken(signedToken string) (claimData *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("DUYDAO"), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return claims, nil

}
