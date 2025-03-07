package main

import "log"

func LogPrintln(message ...string) {
	log.Println(LOG_PREFIX, message)
}

func LogFatal(message ...string) {
	log.Fatal(LOG_PREFIX, message)
}
