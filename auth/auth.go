package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserAccessClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

const (
	secretKey  = "sample_secret_key"
	cookieName = "access_token"
)

func GenerateAndSet(c echo.Context, user db.User) error {
	expiresAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newUserAccessClaims(user.ID.String(), user.Role, expiresAt.Unix()))
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = tokenString
	cookie.Expires = expiresAt
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)

	return nil
}

func Delete(c echo.Context) {
	expiresAt := time.Now().Add(-1 * time.Hour)

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = ""
	cookie.Expires = expiresAt
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func Parse(c echo.Context, auth string) (interface{}, error) {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return UserAccessClaims{}, err
	}

	if token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	}

	return UserAccessClaims{}, fmt.Errorf("invalid token: %v", token)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err == nil
}

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func newUserAccessClaims(id, role string, exp int64) *UserAccessClaims {
	return &UserAccessClaims{
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
}
