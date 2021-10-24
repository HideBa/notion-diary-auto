package rand

func GenerateState() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randStr := GenerateRandomStr(letters, 100)
	return randStr
}
