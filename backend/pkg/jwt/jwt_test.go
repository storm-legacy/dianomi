package jwt

import (
	"testing"
	"time"

	"github.com/rs/xid"
)

func TestTokenGeneration(t *testing.T) {
	var userID uint64 = 332984
	// Token generation
	token, err := GenerateToken(userID)
	if err != nil || token == "" {
		t.Errorf("Token could not be generated: %s", err.Error())
	}
	// Get claims and check if generated token is valid
	claims, err := ExtractClaims(token)
	if err != nil || claims == nil {
		t.Errorf("Claims could not be extracted from token: %s", err.Error())
	}
	sub := uint64((*claims)["sub"].(float64))
	iss := (*claims)["iss"]
	aud := (*claims)["aud"]
	iat := time.Unix(int64((*claims)["iat"].(float64)), 0)
	nbf := time.Unix(int64((*claims)["iat"].(float64)), 0)
	exp := time.Unix(int64((*claims)["iat"].(float64)), 0)

	if sub != userID {
		t.Errorf("UserID doesn't overlap with claims subject: %d:%d", userID, sub)
	}
	if iss != jwtIssuer || aud != jwtIssuer {
		t.Errorf("Claims issuer is not the same as specified")
	}

	now := time.Now()
	if iat.Unix() > now.Unix() {
		t.Errorf("Issued_at is specified into the future")
	}
	if nbf.Unix() < exp.Unix() {
		t.Errorf("Token will never be valid (nbf < exp)")
	}
	if now.Unix() > exp.Unix() {
		t.Errorf("Token will never be valid (exp > now)")
	}
}

func TestCustomClaims(t *testing.T) {
	customClaims1 := make(map[string]string)
	customClaims1["test1"] = "02193u410u41ufqj0q-823u80"
	customClaims1["test2"] = "asdf65438@##$r!^@$(u)sa09"

	customClaims2 := make(map[string]string)
	customClaims2["test3"] = "a0u99r13u-iU{R(#@U)(*@#U)}"
	customClaims2["test4"] = "0JKA:LA@$U:(@j093w2rj90wj)"

	token, err := GenerateToken(123, customClaims1, customClaims2)
	if err != nil || token == "" {
		t.Errorf("Token could not be generated: %s", err.Error())
	}

	// Get claims and check if generated token is valid
	claims, err := ExtractClaims(token)
	if err != nil || claims == nil {
		t.Errorf("Claims could not be extracted from token: %s", err.Error())
	}
	test1 := (*claims)["test1"].(string)
	test2 := (*claims)["test2"].(string)
	test3 := (*claims)["test3"].(string)
	test4 := (*claims)["test4"].(string)

	if test1 != customClaims1["test1"] ||
		test2 != customClaims1["test2"] ||
		test3 != customClaims2["test3"] ||
		test4 != customClaims2["test4"] {
		t.Error("Failed one or more custom claim check")
	}

}

func TestJtiRevoke(t *testing.T) {
	guid := xid.New()
	jti := guid.String()

	// Insert value to cache database
	RevokeToken(jti, time.Now())

	// Check if token is revoked
	status := IsTokenRevoked(jti)
	if !status {
		t.Error("Problem with checking if Jti was revoked")
	}

}

func TestJtiFalsePositive(t *testing.T) {
	guid := xid.New()
	jti := guid.String()

	// Check if token is revoked
	if IsTokenRevoked(jti) {
		t.Error("Jti wasn't revoked, but returned positive")
	}
}
