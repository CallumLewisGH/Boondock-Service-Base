package authenticationHelper

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Memory      uint32 // Memory cost in KiB (e.g., 64 * 1024 = 64MB)
	Iterations  uint32 // Time cost (number of passes)
	Parallelism uint8  // Number of threads
	SaltLength  uint32 // Salt length (16-32 bytes recommended)
	KeyLength   uint32 // Output hash length (32 bytes recommended)
}

var (
	// Default Argon2 parameters (adjust based on your hardware)
	defaultParams = Argon2Params{
		Memory:      64 * 1024, // 64MB
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}
)

// GenerateHash hashes a password with Argon2 and a random salt.
func GenerateHash(password string) (string, error) {
	// Generate a random salt
	salt := make([]byte, defaultParams.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Generate the hash using Argon2id (recommended variant)
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultParams.Iterations,
		defaultParams.Memory,
		defaultParams.Parallelism,
		defaultParams.KeyLength,
	)

	// Encode salt and hash in a string (format: argon2id$params$salt$hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a standardized string format
	return fmt.Sprintf(
		"argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultParams.Memory,
		defaultParams.Iterations,
		defaultParams.Parallelism,
		encodedSalt,
		encodedHash,
	), nil
}

// VerifyPassword checks if a password matches a stored hash.
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Split the encoded hash into parts
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// Parse parameters
	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, errors.New("unsupported Argon2 version")
	}

	var params Argon2Params
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Recompute the hash with the same parameters
	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		uint32(len(storedHash)),
	)

	return subtle.ConstantTimeCompare(computedHash, storedHash) == 1, nil
}
