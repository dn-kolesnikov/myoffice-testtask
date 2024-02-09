package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"

	"myoffice-test-task/internal/config"
	"myoffice-test-task/pkg/filereader"
	"myoffice-test-task/pkg/httpclient"
)

func Run(cfg config.Config) error {
	// TODO:
	ctx := context.TODO()

	// open file
	fh, err := os.Open(cfg.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer fh.Close()

	client := httpclient.New()

	lines := filereader.New(fh).Read(ctx)

	wg := sync.WaitGroup{}

	ncpu := runtime.NumCPU()

	for i := 0; i < ncpu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := range lines {
				select {
				case <-ctx.Done():
					return
				default:
					if client.ValidateURL(l) != nil {
						continue
					}
					// Magic happens here
					r, err := client.Get(l)
					slog.Info("URL Parsed",
						slog.String("url", l),
						slog.Duration("time", r.HandleTime()),
						slog.Int("length", r.ContentLength()),
						slog.Any("error", err),
					)
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
