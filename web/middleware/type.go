package middleware

import "log"

// LogError .
func LogError(e string) string {
	log.Printf("web fail  [Msg]: %s", e)
	return e
}
