package util

import "log"

func CopeError(e error, message string) {
	if e != nil {
		log.Println(message)
	}
}
