package misc

// TODO Replace with slog? Or something better?
import (
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

func LogMessage(logger log.Logger, message string) {
	for _, token := range strings.Split(message, "\n") {
		logger.Log("message", token)
	}
}

func ShortenUUID(id uuid.UUID, prefixLen, suffixLen int) string {
	hex := strings.ReplaceAll(id.String(), "-", "")
	if len(hex) <= prefixLen+suffixLen {
		return id.String()
	}
	return hex[:prefixLen] + "…" + hex[len(hex)-suffixLen:]
}
