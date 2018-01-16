package api

// MockWriteLog mock version of writeLog
func MockWriteLog(data string) error {
	return nil
}

// MockValidUser mock version of validUser
func MockValidUser() bool {
	return true
}
