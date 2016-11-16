package wphash

func HashPassword(password string) string {
	random := sixCharRandom()
	salt := generatePrivateSalt(random)
	return cryptPrivate(password, salt)
}

func CheckPassword(password string, storedHash string) bool {
	hash := cryptPrivate(password, storedHash)
	return hash == storedHash
}