package api

import (
	"time"
)

func Retry(request func() error, attempts int, delay time.Duration) error {
	var err error
	for i := 0; i < attempts; i++ {
		// rety the request
		if err = request(); err == nil {
			return nil
		}
		time.Sleep(delay) // wait before retrying
		delay *= 2        // exponential backoff
	}
	return err
}
