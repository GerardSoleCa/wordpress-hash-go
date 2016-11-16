package wphash

// HashPassword function to hash a given password
func HashPassword(password string) string {
	random := sixCharRandom()
	salt := generatePrivateSalt(random)
	return cryptPrivate(password, salt)
}

// CheckPassword function to check if a password and its hash are valid or not
func CheckPassword(password string, storedHash string) bool {
	hash := cryptPrivate(password, storedHash)
	return hash == storedHash
}
