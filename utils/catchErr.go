package utils

import (
	"log"
	"os"
)

// CatchErr 捕捉错误
func CatchErr(logInfo string, err error) bool {
	if err != nil {
		logFile, _ := os.OpenFile("../logs/ByteDance.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		log.SetOutput(logFile)
		log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
		log.Printf(logInfo+": %s\n", err)
		return false
	}
	return true
}
