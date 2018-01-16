package api

import (
	"io/ioutil"
	"os"
)

func writeLog(data string) error {
	return ioutil.WriteFile("/var/log/api.log", []byte(data), 0644)
}

func validUser() bool {
	return os.Getuid() == 0
}
