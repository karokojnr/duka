package helper

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/karokojnr/duka/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.New(`password should contain at least 6 characters`)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		// todo: log actual error and report to logging tool
		return "", errors.New(`server error, please try again later`)
	}

	return string(hashedPassword), nil
}

func (a Auth) GenerateToken(id uint, email, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		return "", errors.New(`unable to generate token`)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New(`server error, please try again later`)
	}
	return tokenString, nil
}

func (a Auth) VerifyPassword(plainPassword, hashedPassword string) error {
	if len(plainPassword) < 6 {
		return errors.New(`password should contain at least 6 characters`)
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New(`password doesn't match`)
	}
	return nil
}

func (a Auth) VerifyToken(t []string) (domain.User, error) {
	tkn := t[0]
	tokenArr := strings.Split(tkn, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]
	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New(`invalid token`)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unkown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New(`invalid signing method`)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token has expired")
		}
		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserRole = claims["role"].(string)

		return user, nil
	}
	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]
	usr, err := a.VerifyToken(authHeader)
	if err == nil && usr.ID > 0 {
		ctx.Locals("user", usr)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "unauthorized",
			"data":    err,
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
}
