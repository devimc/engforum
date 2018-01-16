package api

var (
	// WriteLogFuncImpl implementation
	WriteLogFuncImpl = writeLog

	// ValidUserFuncImpl implementation
	ValidUserFuncImpl = validUser
)

// WriteLog write a log
func WriteLog(data string) error {
	return WriteLogFuncImpl(data)
}

// ValidUser returns true if the current user is valid
func ValidUser() bool {
	return ValidUserFuncImpl()
}
