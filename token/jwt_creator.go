package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTCreator struct {
	secretKey string
}

func NewJWTCreator(secretKey string) (Creator, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("key size should atleast be %d", minSecretKeySize)
	}

	return &JWTCreator{secretKey}, nil
}




func (creator *JWTCreator) CreateToken(username string, duration time.Duration) (string, *Payload, error){
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload,  err
	}
	
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token,err := jwtToken.SignedString([]byte(creator.secretKey))
	return token, payload, err
}


func (creator *JWTCreator) VerifyToken(token string) (*Payload, error){
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(creator.secretKey), nil
	}
	
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{},keyFunc)

	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return payload, nil

}