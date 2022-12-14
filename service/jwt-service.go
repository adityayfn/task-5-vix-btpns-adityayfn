package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string, username string, email string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}
func NewJWTService() JWTService {
	return &jwtService{
		issuer: "adityayfn",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "adityayfn"
	}
	return secretKey
}
func (j *jwtService) GenerateToken(UserID string, Username string, Email string) string {
	claims := &jwtCustomClaim{
		UserID,
		Username,
		Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
