package time

import (
	"context"
	"time"
)

func Ticker(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

Loop: // more about label: https://juejin.cn/post/7030996392487157773
	for {
		select {
		case <-ticker.C:
			println("hi")
		case <-ctx.Done():
			break Loop
		}
	}
	println("finished")
}
