package argon2

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	password := ".Qwert_y123!$#@%*&^)"

	encodedPassword, err := EncodePassword(&password)
	if err != nil {
		t.Errorf(`Password could not be encoded: %s`, err.Error())
	}

	if result, err := ComparePasswordAndHash(&password, &encodedPassword); !result {
		if err != nil {
			t.Errorf("Error occured while trying to compare encoded and plain password: %s", err.Error())
			return
		}
		t.Errorf(`Provided plaintext password is not the same as encoded one`)
	}
}
