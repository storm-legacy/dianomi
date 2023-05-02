// To be replaced by other solution
// Revokation is not persistent with this implementation
package jwt

import (
	"time"

	"golang.org/x/exp/slices"
)

var (
	revokedTokens []string
)

func IsTokenRevoked(jti string) bool {
	return slices.Contains(revokedTokens, jti)
}

func RevokeToken(jti string, validUntil time.Time) {
	revokedTokens = append(revokedTokens, jti)
}
