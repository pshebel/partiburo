package utils

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"

	"github.com/pshebel/partiburo/backend/models"
)

func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}


// ToHashString converts the struct to a Base64 string
func ToHashString(p models.Token) (string, error) {
	// 1. Convert struct to JSON bytes
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	// 2. Encode bytes to Base64 string
	return base64.StdEncoding.EncodeToString(b), nil
}

// FromHashString recovers the struct from a Base64 string
func FromHashString(s string) (models.Token, error) {
	var p models.Token

	// 1. Decode Base64 string back to bytes
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return p, err
	}

	// 2. Unmarshal JSON bytes back into the struct
	err = json.Unmarshal(b, &p)
	return p, err
}