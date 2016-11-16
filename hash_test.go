package wphash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestHashingAndChecking(t *testing.T) {
	const password = "test"
	hashedPassword := HashPassword(password)
	result := CheckPassword(password, hashedPassword)
	assert.Equal(t, result, true, "Hashing the password and then verifiying it should be equal")
}
