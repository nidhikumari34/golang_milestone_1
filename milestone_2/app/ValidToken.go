package app

import (
	"log"
)

//token verification
func ValidToken(header_token string) bool {
	log.Println(header_token)
	log.Println(TokenString)
	if header_token == TokenString {
		return true
	} else {
		return false
	}
}
