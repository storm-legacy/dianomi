package jwt

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	// "github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/xid"
	"github.com/storm-legacy/dianomi/pkg/config"
)

var (
	privateKey ed25519.PrivateKey
	jwtExp     time.Duration
	jwtIssuer  string
)

func init() {
	var err error
	jwtIssuer = config.GetString("APP_JWT_ISSUER", "development")

	jwtExp, err = time.ParseDuration(config.GetString("APP_JWT_EXPIRED_IN", "60m"))
	if err != nil {
		panic(err)
	}

	privateKeyBase64 := config.GetString("APP_JWT_EDDSA_PRIVATE_KEY", "")
	// No key in .env
	if len(privateKeyBase64) == 0 {
		_, privateKey, _ = ed25519.GenerateKey(rand.Reader)
		return
	}
	// decode
	privateKey, err = base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		panic(err)
	}
}

func GenerateToken(userID uint64, customClaims ...map[string]string) (string, error) {
	// Baseline claims
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"jti": xid.New().String(),
		"iss": jwtIssuer,
		"aud": jwtIssuer,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(jwtExp).Unix(),
		"sub": userID,
	}
	// Additional claims
	for _, customClaimsArr := range customClaims {
		for key, claim := range customClaimsArr {
			claims[key] = claim
		}
	}
	// Sign, encrypt and return token
	res := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return res.SignedString(privateKey)
}

func ExtractClaims(token string) (*jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey.Public(), nil
	})

	// Check if any error occured
	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("claims could not be extracted")
	}

	return &claims, nil
}

// func EncryptToken(token string) (string, error) {
// 	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.ED25519, Key: privateKey.Public()}, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	object, err := encrypter.Encrypt([]byte(token))
// 	if err != nil {
// 		return "", err
// 	}
// 	return object.CompactSerialize()
// }

// func DecryptToken(encryptedToken string) (string, error) {
// 	object, err := jose.ParseEncrypted(encryptedToken)
// 	if err != nil {
// 		return "", err
// 	}
// 	decrypted, err := object.Decrypt(privateKey)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return string(decrypted), nil
// }
