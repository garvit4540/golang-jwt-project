package helpers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(password1 string, password2 string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		return false, "Password does not match"
	}
	return true, ""
}

func CheckUserType(ctx *gin.Context, role string) error {
	userType := ctx.GetString("user_type")
	if userType != role {
		err := errors.New("Unauthorized to access this resource")
		return err
	}
	return nil
}

func MatchUserTypeToUid(ctx *gin.Context, userId string) error {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	if userType == "USER" && uid != userId {
		err := errors.New("Unauthorized to access this resource")
		return err
	}
	if err := CheckUserType(ctx, userType); err != nil {
		return err
	}
	return nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
