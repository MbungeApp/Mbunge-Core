/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/MbungeApp/mbunge-core/config"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// GenerateHash -
func GenerateHash(password string) (string, error) {

	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt, err := generateRandomSalt(p.saltLength)
	if err != nil {
		config.ErrorReporter(err.Error())
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d, t=%d, p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil

}

func generateRandomSalt(n uint32) ([]byte, error) {
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// ComparePasswordAndHash -
func ComparePasswordAndHash(password string, encodedHash string) (bool, error) {
	// Extract the parameters, salt and derived key from the encoded password hash
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		config.ErrorReporter(err.Error())
		return false, err
	}

	// Derive the key from the other password using the same parameters
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical.
	// subtle.ConstantTimeCompare() function prevents timing attacks
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil

}

func decodeHash(encodedHash string) (p *params, salt []byte, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		config.ErrorReporter(err.Error())
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}

	_, err = fmt.Sscanf(vals[3], "m=%d, t=%d, p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		config.ErrorReporter(err.Error())
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		config.ErrorReporter(err.Error())
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		config.ErrorReporter(err.Error())
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

var (
	// ErrInvalidHash -
	ErrInvalidHash = errors.New("the encoded hash is not in the correct format")
	// ErrIncompatibleVersion -
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)
