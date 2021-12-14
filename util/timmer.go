package util

import (
	"time"
)

// NewTicker new a wrapped tick for control
func NewTicker(tickFunc func(), duration time.Duration) chan bool {
	ticker := time.NewTicker(duration)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				tickFunc()
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker)

	return stopChan
}

func After(callback func(), millisecond int) {
	go func() {
		select {
		case <-time.After(time.Duration(millisecond) * time.Millisecond):
			callback()
		}
	}()
}
