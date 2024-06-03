package utils

import (
	"sync"
)

func RunInTransaction(txFunc func() error) error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	err := txFunc()
	if err != nil {
		return err
	}

	return nil
}
