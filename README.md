# semaphore

Provides a semaphore synchronization primitive. A semaphore controls access to a finite number of resources.

# Usage

```go
s := semaphore.New(5)

for {
  s.Acquire()
  // Only 5 go-routines will run simultaneously.
  go func() {
    defer s.Release()
  }()
}
```
