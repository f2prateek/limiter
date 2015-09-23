package semaphore

type empty struct{}
type Semaphore chan empty

func New(n int) Semaphore {
	return Semaphore(make(chan empty, n))
}

func (s Semaphore) Acquire() {
	s <- empty{}
}

func (s Semaphore) Release() {
	<-s
}
