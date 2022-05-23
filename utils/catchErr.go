package utils

import (
	"log"
)

func CatchErr(logInfo string, err error) {
	if err != nil {
		log.Printf(logInfo+": %s\n", err)
	}
}
