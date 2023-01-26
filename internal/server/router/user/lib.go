package user

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slog"
	"neutron-star-api/internal/database/dictdb"
	"neutron-star-api/internal/database/userdb"
	"neutron-star-api/internal/server/model/user/usermodel"
	"os"
	"time"
)

func getSecret() (string, error) {
	jwtSecretKeyInEnv := os.Getenv("jwt_secret_key")
	if jwtSecretKeyInEnv != "" {
		return jwtSecretKeyInEnv, nil
	}
	kv, err := dictdb.Get("jwt_secret_key")
	if err != nil {
		return "", err
	}
	err = os.Setenv(kv.Key, kv.Key)
	if err != nil {
		slog.Error("Set jwt_secrete_key env failed", err)
	}
	return kv.Value, nil
}

func createToken(username string) (string, error) {
	secret, err := getSecret()
	if err != nil {
		return "", err
	}
	claims := usermodel.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 15)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	secret, err := getSecret()
	if err != nil {
		return "", nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &usermodel.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	claims, ok := token.Claims.(*usermodel.Claims)
	if !token.Valid || !ok {
		return "", nil
	}
	if err := claims.Valid(); err != nil {
		return "", err
	}
	return claims.Username, nil
}

func verifyPassword(username, password string) error {
	userInDB, err := userdb.Get(username)
	if err != nil {
		return err
	}
	encryptedPassword := encryptPassword(password)
	if encryptedPassword != userInDB.EncryptedPassword {
		return fmt.Errorf("error username or password")
	}
	return nil
}

func encryptPassword(password string) string {
	sum := sha256.New()
	sum.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(sum.Sum(nil))
}
