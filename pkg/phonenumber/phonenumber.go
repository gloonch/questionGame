package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {

	// TODO : technical debt: we can use regex to support +98
	if len(phoneNumber) != 11 {
		return false
	}

	// 0912
	if phoneNumber[0:2] != "09" {
		return false
	}

	// 09827hgsdjb
	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false
	}

	return true
}
