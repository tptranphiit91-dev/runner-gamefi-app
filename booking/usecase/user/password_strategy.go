package user

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher defines the strategy interface for password hashing
// Strategy Pattern: Allows different password hashing algorithms
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

// BcryptHasher implements PasswordHasher using bcrypt
type BcryptHasher struct {
	cost int
}

// NewBcryptHasher creates a new BcryptHasher
func NewBcryptHasher(cost int) *BcryptHasher {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	return &BcryptHasher{cost: cost}
}

// Hash hashes a password using bcrypt
func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

// Compare compares a hashed password with a plain password
func (h *BcryptHasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// SHA256Hasher implements PasswordHasher using SHA256 (less secure, for demo)
type SHA256Hasher struct {
	salt string
}

// NewSHA256Hasher creates a new SHA256Hasher
func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt: salt}
}

// Hash hashes a password using SHA256
func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(password + h.salt))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Compare compares a hashed password with a plain password
func (h *SHA256Hasher) Compare(hashedPassword, password string) error {
	newHash, err := h.Hash(password)
	if err != nil {
		return err
	}
	if newHash != hashedPassword {
		return bcrypt.ErrMismatchedHashAndPassword
	}
	return nil
}

