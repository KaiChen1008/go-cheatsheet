package context

import (
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("ok"))
	case <-ctx.Done():
		http.Error(w, "canceled", http.StatusRequestTimeout)
	}
}
