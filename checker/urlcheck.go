package checker

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Result struct {
	Url      string
	Err      error
	Size     int
	Duration time.Duration
}

type Service struct {
	logger   *slog.Logger
	fileName string
}

func New() *Service {
	return &Service{
		logger:   slog.New(slog.NewTextHandler(os.Stdout, nil)),
		fileName: "results.txt",
	}
}

func (s *Service) Do(ctx context.Context, urls []url.URL) error {
	f, err := os.OpenFile(s.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	oks, errs := uint64(0), uint64(0)

	s.logger.InfoContext(ctx, "проверка запущена")
	defer s.logger.InfoContext(ctx, "проверка завершена",
		slog.Uint64("oks", oks),
		slog.Uint64("errs", errs),
	)

	for _, u := range urls {
		chk := s.checkUrl(u)
		if chk.Err != nil {
			errs++
			_, _ = fmt.Fprintf(f, "ts=%s url=%s size=%d duration=%s\n", time.Now(), u.String(), chk.Size, chk.Duration)
			continue
		}
		oks++
		_, _ = fmt.Fprintf(f, "ts=%s url=%s err= size=%d duration=%s\n", time.Now(), u.String(), chk.Size, chk.Duration)
	}

	return nil
}

func (s *Service) checkUrl(u url.URL) Result {
	start := time.Now()
	resp, err := http.Get(u.String())
	if err != nil {
		return Result{
			Url: u.String(),
			Err: err,
		}
	}

	body, _ := io.ReadAll(resp.Body)

	return Result{
		Url:      u.String(),
		Size:     len(body),
		Duration: time.Since(start),
	}
}
