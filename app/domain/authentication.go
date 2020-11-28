package domain

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SantiagoZuluaga/GoAuth/app/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	var publicBytes []byte
	var privateBytes []byte
	var err error

	publicBytes, err = ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		fmt.Println(err)
	}

	privateBytes, err = ioutil.ReadFile("./private.rsa")
	if err != nil {
		fmt.Println(err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateJWT(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Prueba",
		},
	})
	authToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func ValidateToken(r *http.Request) (bool, string) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			validateError := err.(*jwt.ValidationError)
			switch validateError.Errors {
			case jwt.ValidationErrorExpired:
				return false, "EXPIRED TOKEN"
			case jwt.ValidationErrorSignatureInvalid:
				return false, "INVALID SIGNATURE"
			default:
				return false, "INVALID TOKEN"
			}
		default:
			return false, "INVALID TOKEN"
		}
	}

	if token.Valid {
		return true, "VALID TOKEN"
	}

	return false, "INVALID TOKEN"
}
