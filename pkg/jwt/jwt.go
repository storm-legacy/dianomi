package jwt

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/storm-legacy/dianomi/pkg/config"
)

var (
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey

	jwtExpiredIn time.Duration = time.Duration(config.GetInt("APP_JWT_EXPIRED_IN"))
	jwtMaxAge    time.Duration = time.Duration(config.GetInt("APP_JWT_MAXAGE"))

	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

func init() {
	// Check for ed25519 private key and generate if needed
	privateKeyBase64 := viper.GetString("APP_JWT_EDDSA_PRIVATE_KEY")
	privateKey, _ = base64.StdEncoding.DecodeString(privateKeyBase64)
	publicKeyBase64 := viper.GetString("APP_JWT_EDDSA_PUBLIC_KEY")
	publicKey, _ = base64.StdEncoding.DecodeString(publicKeyBase64)

	// Generate key pair if any empty
	if len(privateKey) != 64 || len(publicKey) != 32 {
		publicKey, privateKey, _ = ed25519.GenerateKey(rand.Reader)
		publicKeyBase64 = base64.StdEncoding.EncodeToString(publicKey)
		privateKeyBase64 = base64.StdEncoding.EncodeToString(privateKey)
		log.WithFields(log.Fields{"privateKey": privateKeyBase64, "publicKey": publicKeyBase64}).Warn("Wrong or missing values for JWT, keys were auto-generated")
	} else {
		log.WithFields(log.Fields{"privateKey": privateKeyBase64, "publicKey": publicKeyBase64}).Debug("Loaded private and public keys")
	}
}
