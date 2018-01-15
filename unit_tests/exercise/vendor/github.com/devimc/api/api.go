package api

// WriteLog write a log
func WriteLog(data string) error {
	return writeLog(data)
}

// ValidUser returns true if the current user is valid
func ValidUser() bool {
	return validUser()
}
