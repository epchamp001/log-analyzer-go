package log

import (
	"log-analyzer-go/internal/models"
)

func AggregateLogs(entries []models.LogEntry) map[string]int {
	stats := make(map[string]int)
	for _, entry := range entries {
		stats[entry.IP]++
	}
	return stats
}
