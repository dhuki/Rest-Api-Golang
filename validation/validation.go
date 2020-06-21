package validation

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id, role string) (string, error) {
	expired := time.Now().Add(time.Hour * 2) // expired 2 hour from at the time

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "dhukidwir.herokuapp.com",     // identify token issued by
		Audience:  "api.dhukidwir.herokuapp.com", // identify more specific from issuer for client validation
		Id:        id,                            // id user logged
		Subject:   role,                          // information about user who logged
		IssuedAt:  time.Now().Local().UnixNano(), // issued at the time and convert it to nano second
		ExpiresAt: expired.Local().UnixNano(),    // expiration
	})

	tokenString, err := token.SignedString(SECRET) // complete generate token with secret
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) { // authorization process, verifying if you've access ?
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		// SECRET is a []byte containing your secret, e.g. []byte("my_secret_key") to validate secret_key in token
		return []byte(SECRET), nil // conversion to byte
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token is expired")
	}

	return token, nil
}
