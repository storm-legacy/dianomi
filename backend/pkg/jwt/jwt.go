package jwt

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
)

var (
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey

	jwtExpiredIn time.Duration
	// jwtMaxAge    time.Duration
)

func init() {

	// Check for ed25519 private key and generate if needed
	privateKeyBase64 := config.GetString("APP_JWT_EDDSA_PRIVATE_KEY")
	privateKey, _ = base64.StdEncoding.DecodeString(privateKeyBase64)
	publicKeyBase64 := config.GetString("APP_JWT_EDDSA_PUBLIC_KEY")
	publicKey, _ = base64.StdEncoding.DecodeString(publicKeyBase64)

	// Generate key pair if any empty
	if len(privateKey) != 64 || len(publicKey) != 32 {
		publicKey, privateKey, _ = ed25519.GenerateKey(rand.Reader)
		publicKeyBase64 = base64.StdEncoding.EncodeToString(publicKey)
		privateKeyBase64 = base64.StdEncoding.EncodeToString(privateKey)
		config.Set("APP_JWT_EDDSA_PUBLIC_KEY", publicKeyBase64)
		log.WithFields(log.Fields{"privateKey": privateKeyBase64, "publicKey": publicKeyBase64}).Warn("Wrong or missing values for JWT, keys were auto-generated")
	} else {
		log.WithFields(log.Fields{"privateKey": privateKeyBase64, "publicKey": publicKeyBase64}).Debug("Loaded private and public keys")
	}

	// Load durations
	jwtExpiredIn, _ = time.ParseDuration(config.GetString("APP_JWT_EXPIRED_IN"))
	// jwtMaxAge, _ = time.ParseDuration(config.GetString("APP_JWT_MAXAGE"))
}

func ParseToken(token string) (claims jwt.MapClaims, err error) {
	// Check if provided token was signed with ed25519
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	// Check if any error occured
	if err != nil {
		return nil, err
	}

	// Extract claims
	var ok bool
	claims, ok = parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("claims could not be extracted")
	}

	return claims, nil
}

func GenerateToken(id int64, email string, role string) (result string, err error) {
	// Generate token string for user
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":   id,
		"exp":   now.Add(jwtExpiredIn).Unix(),
		"iat":   now.Unix(),
		"nbf":   now.Unix(),
		"email": email,
		"role":  role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	result, err = token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return result, nil
}
