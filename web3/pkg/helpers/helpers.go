package helpers

import "log"

func ErrorCheck(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}
