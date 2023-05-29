package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/surysatriah/go-dashboard-app/pkg"
)

var privateKey = []byte(pkg.GetDotEnvVariable("JWT_PRIVATE_KEY"))

func GenerateJWT(id uint, email string) string {
	tokenTTL, _ := strconv.Atoi(pkg.GetDotEnvVariable("JWT_TOKEN_TTL"))
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"iat":   time.Now().Unix(),
		"eat":   time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString(privateKey)

	return signedToken
}

func ValidateJWT(c *gin.Context) error {

	token, err := getToken(c)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid jwt provided")

}

// func CurrentUser(c *gin.Context) (model.User, error) {
// 	err := ValidateJWT(c)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	token, _ := getToken(c)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := uint(claims["id"].(float64))

// 	user, err := model.FindUserById(userId)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, nil
// }

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}
