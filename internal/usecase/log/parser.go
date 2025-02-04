package log

import (
	"fmt"
	"log-analyzer-go/internal/models"
	"strconv"
	"strings"
)

func ParseLog(line string) (models.LogEntry, error) {
	parts := strings.Fields(line)
	if len(parts) < 9 {
		return models.LogEntry{}, fmt.Errorf("invalid log format")
	}

	timestamp := strings.Trim(parts[3], "[") + " " + strings.Trim(parts[4], "]")

	method := strings.Trim(parts[5], `"`)
	path := parts[6]

	statusCode, err := strconv.Atoi(parts[8])
	if err != nil {
		return models.LogEntry{}, fmt.Errorf("invalid status code: %v", err)
	}

	entry := models.LogEntry{
		IP:         parts[0],
		Timestamp:  timestamp,
		Method:     method,
		Path:       path,
		StatusCode: statusCode,
	}

	return entry, nil
}
