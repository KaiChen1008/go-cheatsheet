package time

import (
	"context"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Ticker(ctx)
}
