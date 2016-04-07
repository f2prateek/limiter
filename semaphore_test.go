package semaphore_test

import (
	"testing"
	"time"

	"github.com/f2prateek/semaphore"
)

func TestSemaphoreAcquires(t *testing.T) {
	s := semaphore.New(1)

	acquired := make(chan struct{})
	go func() {
		s.Acquire(1)
		acquired <- struct{}{}
	}()

	select {
	case <-acquired:
	case <-time.After(1 * time.Second):
		t.Errorf("semaphore not acquired")
	}
}

func TestSemaphoreDoesNotAcquireAtCapacity(t *testing.T) {
	s := semaphore.New(2)
	s.Acquire(2)
	acquired := make(chan struct{})
	go func() {
		s.Acquire(1)
		acquired <- struct{}{}
	}()

	select {
	case <-acquired:
		t.Errorf("semaphore acquired")
	case <-time.After(1 * time.Second):
	}
}

func TestSemaphoreReleases(t *testing.T) {
	s := semaphore.New(2)
	s.Acquire(2)
	acquired := make(chan struct{})
	go func() {
		s.Acquire(1)
		acquired <- struct{}{}
	}()

	s.Release(1)

	select {
	case <-acquired:
	case <-time.After(1 * time.Second):
		t.Errorf("semaphore not acquired")
	}
}
