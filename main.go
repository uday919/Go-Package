package mylogger

import "log"

func Loginfo(message string) {
	log.Printf("INFO - %v", message)
}
func LogWarn(message string) {
	log.Printf("Warn - %v", message)
}
func LogError(message string) {
	log.Printf("Err - %v", message)
}
