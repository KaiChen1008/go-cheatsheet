package context

import (
	"net/http"
	"time"
)

// ref: https://pjchender.dev/golang/pkg-context/

/*
diff between WithTimeout and WithDeadline:
- WithTimeout: input is a duration (-> WithTimout call WithDeadline internally)
- WithDeadline: input is a specific time
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("ok"))
	case <-ctx.Done():
		http.Error(w, "canceled", http.StatusRequestTimeout)
	}
}
