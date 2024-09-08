package utils

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Generate Random Hash
func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// HashPassword hashes the given password with the provided salt and pepper
func HashPassword(password, salt, pepper string) string {
	hash := sha256.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(password))
	hash.Write([]byte(pepper))
	return hex.EncodeToString(hash.Sum(nil))
}

func VerifyPassword(password, salt, hashedPassword string) bool {
	globalConf := config.GetGlobalConfig()

	pepper := globalConf.Login.PasswordPepper
	return HashPassword(password, salt, pepper) == hashedPassword
}

func GeneratePassword(password string) (string, string, error) {
	globalConf := config.GetGlobalConfig()

	// Generate a salt
	salt, err := GenerateSalt()
	if err != nil {
		return "", "", err
	}

	// Get hash from env
	pepper := globalConf.Login.PasswordPepper

	// Generate Hash the password with the salt and pepper
	hashedPassword := HashPassword(password, salt, pepper)

	return salt, hashedPassword, nil
}
