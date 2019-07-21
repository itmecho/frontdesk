package types_test

import (
	"testing"

	. "github.com/itmecho/frontdesk/pkg/types"
)

func TestUserPasswordHashing(t *testing.T) {
	passwords := [][]byte{
		[]byte("testpassword"),
	}

	u := &User{}

	blankPassword := []byte("")

	for _, pwd := range passwords {

		if err := u.SetPassword(blankPassword); err != nil {
			t.Error("failed to set user's password: ", err)
		}

		if u.CheckPassword(pwd) {
			t.Error("expected password check to fail but it succeeded")
		}

		if err := u.SetPassword(pwd); err != nil {
			t.Error("failed to set user's password: ", err)
		}

		if !u.CheckPassword(pwd) {
			t.Error("expected password check to succeed but it failed")
		}
	}
}
