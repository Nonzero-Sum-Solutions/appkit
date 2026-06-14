package misc

// TODO Replace with slog? Or something better?
import (
	"strings"

	"github.com/go-kit/kit/log"
)

func LogMessage(logger log.Logger, message string) {
	for _, token := range strings.Split(message, "\n") {
		logger.Log("message", token)
	}
}
