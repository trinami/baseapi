package utils

import (
	"testing"
	"strings"
	"encoding/base64"
)

func TestGenerateHash(t *testing.T) {
	password := "mySecurePassword"
	hash, err := GenerateHash(password)
	if err != nil {
		t.Fatalf("GenerateHash returned an error: %v", err)
	}

	if len(hash) == 0 {
		t.Fatalf("GenerateHash returned an empty hash")
	}

	parts := strings.Split(hash, ".")
	if len(parts) != 2 {
		t.Fatalf("GenerateHash returned an invalid hash format: %v", hash)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		t.Fatalf("GenerateHash returned an invalid salt: %v", err)
	}

	if len(salt) != saltLength {
		t.Fatalf("GenerateHash returned a salt of invalid length: got %d, want %d", len(salt), saltLength)
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	password := "mySecurePassword"
	hash, err := GenerateHash(password)
	if err != nil {
		t.Fatalf("GenerateHash returned an error: %v", err)
	}

	match, err := CompareHashAndPassword(hash, password)
	if err != nil {
		t.Fatalf("CompareHashAndPassword returned an error: %v", err)
	}

	if !match {
		t.Fatalf("CompareHashAndPassword did not match the correct password")
	}

	wrongPassword := "wrongPassword"
	match, err = CompareHashAndPassword(hash, wrongPassword)
	if err != nil {
		t.Fatalf("CompareHashAndPassword returned an error: %v", err)
	}

	if match {
		t.Fatalf("CompareHashAndPassword matched an incorrect password")
	}
}
