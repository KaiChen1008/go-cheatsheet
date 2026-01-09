package errorgroup

import (
	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	eg.Go(func() error {
		return nil
	})

	if err := eg.Wait(); err != nil {
		println(err)
	}
}
