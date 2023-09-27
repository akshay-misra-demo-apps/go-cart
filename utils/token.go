package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/akshay-misra-demo-apps/go-cart/constants"
	"github.com/akshay-misra-demo-apps/go-cart/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenWithRefresh(email string,
	firstName string,
	userType string,
	uid string) (signedToken string, signedRefreshToken string, err error) {

	claims := &models.SignedDetails{
		Email:     email,
		FirstName: firstName,
		Uid:       uid,
		UserType:  userType,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Local().Add(time.Second * time.Duration(30)).Unix(),
			"iat": time.Now().Local().Unix(),
			"iss": "go-cart-auth",
			"sub": email,
		},
	}

	refreshClaims := &models.SignedDetails{
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Local().Add(time.Minute * time.Duration(5)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(constants.SECRET))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(constants.SECRET))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *models.SignedDetails, err error) {
	fmt.Println("Validate Token: signedToken: ", signedToken)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		err = fmt.Errorf("invalid token")
		return
	}

	expTime, err := claims.MapClaims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expTime.Unix() < time.Now().Local().Unix() {
		err = fmt.Errorf("token is expired")
		return
	}
	return claims, err
}
