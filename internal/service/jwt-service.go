package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string, Name string, Otorisasi string, Email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Otoritas string `json:"otoritas"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "_megono",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "_megono"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string, Name string, Otoritas string, Email string) string {
	claims := &jwtCustomClaim{
		UserID,
		Name,
		Email,
		Otoritas,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token_catch string) (*jwt.Token, error) {
	token := ExtractTokenMidleware(token_catch)
	// println(token)
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func ExtractTokenMidleware(bearerToken string) string {
	// keys := r.URL.Query()
	// token := keys.Get("token")
	// if token != "" {
	// 	return token
	// }
	// bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
