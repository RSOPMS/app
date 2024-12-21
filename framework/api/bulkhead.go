package api

import (
	"errors"
	"sync"
)

type Bulkhead struct {
	maxConcurrent int
	current       int
	mutex         sync.Mutex
}

func NewBulkhead(maxConcurrent int) *Bulkhead {
	return &Bulkhead{
		maxConcurrent: maxConcurrent,
	}
}

func (b *Bulkhead) Execute(request func() error) error {
	b.mutex.Lock()
	if b.current >= b.maxConcurrent {
		b.mutex.Unlock()
		return errors.New("bulkhead limit reached")
	}
	b.current++
	b.mutex.Unlock()

	defer func() {
		b.mutex.Lock()
		b.current--
		b.mutex.Unlock()
	}()

	return request()
}
