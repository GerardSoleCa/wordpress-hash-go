package wphash

import "fmt"

func HashPassword(password string) string {
	salt := generatePrivateSalt(sixCharRandom())
	fmt.Println(salt)
	return cryptPrivate(password, salt)
}

func CheckPassword(password string, storedHash string) bool {
	hash := cryptPrivate(password, storedHash)
	return hash == storedHash
}