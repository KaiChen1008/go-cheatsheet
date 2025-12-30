package polling

import (
	"context"
	"math/rand/v2"
	"time"
)

// polling -> select
// Prefer using select over polling or other busy-waiting techniques for more efficient resource utilization.

// forever polling
func ForeverPolling(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		// it is okay if duration of checkStatus > 1sec.
		// <-ticker.C and checkStatus are done simultaneously
		case <-ticker.C:
			checkStatus()

		case <-ctx.Done():
			return
		}
	}
}

// polling five times
func Polling(ctx context.Context, maxAttampts int) {
	// poll directly
	if status := checkStatus(); status == 1 {
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range maxAttampts - 1 {
		select {
		case <-ticker.C:
			if status := checkStatus(); status == 1 {
				// break <- pitfall, only beak select
				return
			}
		case <-ctx.Done():
			return
		}
	}

}

func checkStatus() int {
	return rand.Int()
}
