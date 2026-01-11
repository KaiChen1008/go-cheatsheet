package concurrency

import "sync"

// see how sync.Once implements
// the following reference explains why we need both slow and fast paths.
// ref: https://zhuanlan.zhihu.com/p/683760105 !important

var (
	once   sync.Once
	loaded bool
)

func Once() {
	once.Do(func() {
		loaded = true
	})
}
