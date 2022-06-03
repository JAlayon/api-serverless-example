package validators

import (
	"net/mail"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	if len(email) < 3 || len(email) > 254 || err != nil {
		return false
	}
	return true
}
