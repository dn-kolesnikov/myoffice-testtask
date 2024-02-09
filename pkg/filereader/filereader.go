package filereader

import (
	"bufio"
	"context"
	"io"
	"sync"
)

// Reader -.
type Reader struct {
	sc *bufio.Scanner
	ch chan string
}

// New -.
func New(r io.Reader) *Reader {
	return &Reader{
		sc: bufio.NewScanner(r),
		ch: make(chan string),
	}
}

// Read -.
func (r *Reader) Read(ctx context.Context) chan string {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for r.sc.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
				r.ch <- r.sc.Text()
			}
		}
	}()

	go func() {
		wg.Wait()
		close(r.ch)
	}()

	return r.ch
}
