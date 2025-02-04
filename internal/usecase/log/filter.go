package log

import (
	"strconv"
	"strings"
)

func FilterLog(line string) bool {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return false
	}

	lastPart := parts[len(parts)-1]

	statusCode, err := strconv.Atoi(lastPart)
	if err != nil {
		return false
	}

	return statusCode >= 400 && statusCode <= 599
}
