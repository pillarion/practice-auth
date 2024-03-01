package password

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// Hash takes a password string and returns the hashed password, the salt used for hashing, the hash value, and an error if any.
//
// to DB string, Salt string, Hash valuer string, error
func Hash(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	stringSalt := hex.EncodeToString(salt)

	toHash := fmt.Sprintf("%s%s", stringSalt, password)

	sha512 := sha512.New()
	sha512.Write([]byte(toHash))
	hashValue := hex.EncodeToString(sha512.Sum(nil))

	toDb := fmt.Sprintf("%s$%s", stringSalt, hashValue)

	return toDb, nil
}

// Check checks if the provided password matches the given hash.
//
// It takes in two parameters: password of type string, and hash of type string.
// It returns a boolean value.
func Check(password string, hash string) bool {
	salt := strings.Split(hash, "$")[0]
	hashValue := strings.Split(hash, "$")[1]

	toHash := fmt.Sprintf("%s%s", salt, password)
	sha512 := sha512.New()
	sha512.Write([]byte(toHash))
	hashedPass := hex.EncodeToString(sha512.Sum(nil))

	return hashedPass == hashValue
}
