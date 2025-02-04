package models

type LogEntry struct {
	IP         string
	Timestamp  string
	Method     string
	Path       string
	StatusCode int
}
