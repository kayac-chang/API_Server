package utils

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func CheckMail(email string) error {

	err := checkmail.ValidateFormat(email)

	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {

		return fmt.Errorf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
	}

	return nil
}
