package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Dir            string
	Level          string
	RetentionDays  int
	ConsoleEnabled bool
}

func Init(cfg Config) (*slog.Logger, error) {
	if cfg.Dir == "" {
		return nil, fmt.Errorf("log dir is empty")
	}

	if err := os.MkdirAll(cfg.Dir, 0o755); err != nil {
		return nil, err
	}

	if err := cleanupOldFiles(cfg.Dir, cfg.RetentionDays); err != nil {
		return nil, err
	}

	level := parseLevel(cfg.Level)
	fileHandler := slog.NewTextHandler(newDailyFileWriter(cfg.Dir), &slog.HandlerOptions{Level: level})
	handlers := []slog.Handler{fileHandler}

	if cfg.ConsoleEnabled {
		consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handlers = append(handlers, consoleHandler)
	}

	logger := slog.New(newMultiHandler(handlers...))
	slog.SetDefault(logger)
	return logger, nil
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

type dailyFileWriter struct {
	dir         string
	mu          sync.Mutex
	currentDate string
	file        *os.File
}

func newDailyFileWriter(dir string) io.Writer {
	return &dailyFileWriter{dir: dir}
}

func (w *dailyFileWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	date := time.Now().In(time.Local).Format("2006-01-02")
	if w.file == nil || w.currentDate != date {
		if w.file != nil {
			_ = w.file.Close()
		}

		if err := os.MkdirAll(w.dir, 0o755); err != nil {
			return 0, err
		}

		path := filepath.Join(w.dir, date+".log")
		file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return 0, err
		}
		w.file = file
		w.currentDate = date
	}

	return w.file.Write(p)
}

type multiHandler struct {
	handlers []slog.Handler
}

func newMultiHandler(handlers ...slog.Handler) slog.Handler {
	return &multiHandler{handlers: handlers}
}

func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *multiHandler) Handle(ctx context.Context, record slog.Record) error {
	var firstErr error
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, record); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	cloned := make([]slog.Handler, 0, len(h.handlers))
	for _, handler := range h.handlers {
		cloned = append(cloned, handler.WithAttrs(attrs))
	}
	return &multiHandler{handlers: cloned}
}

func (h *multiHandler) WithGroup(name string) slog.Handler {
	cloned := make([]slog.Handler, 0, len(h.handlers))
	for _, handler := range h.handlers {
		cloned = append(cloned, handler.WithGroup(name))
	}
	return &multiHandler{handlers: cloned}
}

var logFilePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\.log$`)

func cleanupOldFiles(dir string, retentionDays int) error {
	if retentionDays <= 0 {
		return nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	now := time.Now().In(time.Local)
	cutoff := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	cutoff = cutoff.AddDate(0, 0, -retentionDays+1)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !logFilePattern.MatchString(name) {
			continue
		}

		datePart := strings.TrimSuffix(name, ".log")
		parsed, err := time.ParseInLocation("2006-01-02", datePart, time.Local)
		if err != nil {
			continue
		}

		if parsed.Before(cutoff) {
			path := filepath.Join(dir, name)
			if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
	}

	return nil
}
