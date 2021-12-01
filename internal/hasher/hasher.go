// Here can be some copyrights.

// Package hasher provide functionality to dial with passwords. It is meant to
// help everyone work with passwords and passwords hashes.
//
// Notice that you should not extend it and place some other functionality.
//
// If you have any suggestion or comment, please feel free to open an issue on
// this tutorial's GitHub page!
//
// By Mykola Mateichuk
package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword create hash for password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}