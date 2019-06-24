package middleware

import "log"

// Resqonse ..
type Resqonse struct {
	Error int
}

// LogError .
func LogError(e string) string {
	log.Printf("web fail  [Msg]: %s", e)
	return e
}
