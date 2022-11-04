package models

import (
	"net/mail"
	"regexp"
)

func (inv *Invoice) HasValidEmail() bool {
	_, err := mail.ParseAddress(inv.Buyer.Contact.Email)
	return err == nil
}

func (inv *Invoice) IsTestMod() bool {
	if inv.TestMod {
		return true
	} else {
		return false
	}
}

func (inv *Invoice) IsValidPhone() bool {
	regex := regexp.MustCompile("^[0-9]{5,10}$")
	match := regex.Match([]byte(inv.Buyer.Contact.Phone))
	if match {
		return true
	} else {
		return false
	}
}

func IsValidDate(date string) bool {
	regex := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
	match := regex.Match([]byte(date))
	if match {
		return true
	} else {
		return false
	}
}
