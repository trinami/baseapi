package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"golang.org/x/crypto/argon2"
)

const (
	time       = 4
	memory     = 64 * 1024
	threads    = 4
	keyLength  = 32
	saltLength = 16
)

func GenerateHash(password string) (string, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s.%s", b64Salt, b64Hash), nil
}

func CompareHashAndPassword(hash, password string) (bool, error) {
	parts := strings.Split(hash, ".")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	expectedHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	expectedHashB64 := base64.RawStdEncoding.EncodeToString(expectedHash)

	return parts[1] == expectedHashB64, nil
}
