package service

import (
	"avito_test_assingment/pkg/repository"
	"avito_test_assingment/types"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"time"
)

const (
	salt            = "jldsajlf%sfldj#dfsf"
	signingUserKey  = "ajfas#user#ajldfj32"
	signingAdminKey = "fjallknnei%admin$laslnrlk32ncmldsj3ndfosfklm3"
	tokenTTL        = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user types.UserType) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	if user.Role == 0 {
		user.Role = types.User
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(user types.UserType) (string, error) {
	slog.Info("GenerateTOken Role: ", user.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})
	if user.Role == types.User {
		return token.SignedString([]byte(signingUserKey))
	}
	return token.SignedString([]byte(signingAdminKey))
}

func (s *AuthService) ParserToken(accessToken string) (*types.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &types.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		if claims, ok := token.Claims.(*types.TokenClaims); ok {
			slog.Info("Claims.Role = ", claims.Role)
			if claims.Role == types.Admin {
				return []byte(signingAdminKey), nil
			} else if claims.Role == types.User {
				return []byte(signingUserKey), nil
			}
		}

		return nil, errors.New("token claims are not of type *tokenClaims")
	})
	if err != nil {
		slog.Info("Invalid token", err)
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("токен невалиден")
	}

	claims, _ := token.Claims.(*types.TokenClaims)

	return claims, nil
}

func (s *AuthService) CheckAuthData(username, password string) (types.UserType, error) {
	return s.repo.GetUser(username, s.generatePasswordHash(password))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
