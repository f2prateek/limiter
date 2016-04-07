# semaphore

Provides a semaphore synchronization primitive. A semaphore controls access to a finite number of resources.

# Usage

```go
s := semaphore.New(5)

for {
  s.Acquire()
  go func() {
    defer s.Release()

    // Only 5 go-routines will ever be executing simultaneously.
  }()
}
```
