package wphash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	const password = "test"
	hashedPwd := HashPassword(password)
	print(hashedPwd)
	assert.Equal(t, hashedPwd, "", "Hashed password of %s should be equal to %s", password , "taka")
}

func TestCheckPassword(t *testing.T) {

}