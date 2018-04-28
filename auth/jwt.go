package auth

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

const (
	AuthParameter = "key"
)

// CreateToken creates a token with the given claims
func CreateToken(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func Middleware(secret string) *jwtmiddleware.JWTMiddleware {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Extractor:     jwtmiddleware.FromFirst(jwtmiddleware.FromAuthHeader, jwtmiddleware.FromParameter(AuthParameter)),
	})

	return jwtMiddleware
}
