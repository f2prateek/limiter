// Package semaphore provides a semaphore synchronization primitive.
package semaphore

// Semaphore controls access to a finite number of resources.
type Semaphore chan struct{}

// New creates a Semaphore that controls access to `n` resources.
func New(n int) Semaphore {
	return Semaphore(make(chan struct{}, n))
}

// Acquire `n` resources.
func (s Semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

// Release `n` resources.
func (s Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
