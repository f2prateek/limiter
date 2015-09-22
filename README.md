# limiter

Easily ensure only a limited number of requests are run at a time.

# Usage
```
l := limiter.New(5)

for {
    l.Acquire()
    go func() {
      defer l.Release()
      // do work; only 5 go-routines will be
      // executing simultaneously in this case.
    }()
}

```