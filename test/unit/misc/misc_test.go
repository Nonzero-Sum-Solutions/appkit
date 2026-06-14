package misc

import (
	"os"
	"testing"

	klog "github.com/go-kit/kit/log"

	"github.com/mattmunz/appkit/misc"
)

func TestMultiLineLog(t *testing.T) {
	logger := klog.NewLogfmtLogger(klog.NewSyncWriter(os.Stderr))
	misc.LogMessage(logger, "Testing line1\nLine2\nLine3")

	// Uncomment this to see the expected output (3 log messages).
	// require.Fail(t, "See output")
}
