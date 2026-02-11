package logger

import (
	"context"
	"log/slog"
	"os"
)

// WailsAdapter routes Wails logs into the core logger.
type WailsAdapter struct {
	logger *slog.Logger
}

// NewWailsAdapter creates a Wails logger adapter. If l is nil, slog.Default() is used.
func NewWailsAdapter(l *slog.Logger) *WailsAdapter {
	if l == nil {
		l = slog.Default()
	}
	return &WailsAdapter{logger: l.With("source", "wails")}
}

func (w *WailsAdapter) Print(message string) {
	w.logger.Log(context.Background(), slog.LevelInfo, message)
}

func (w *WailsAdapter) Trace(message string) {
	w.logger.Log(context.Background(), slog.LevelDebug, message)
}

func (w *WailsAdapter) Debug(message string) {
	w.logger.Log(context.Background(), slog.LevelDebug, message)
}

func (w *WailsAdapter) Info(message string) {
	w.logger.Log(context.Background(), slog.LevelInfo, message)
}

func (w *WailsAdapter) Warning(message string) {
	w.logger.Log(context.Background(), slog.LevelWarn, message)
}

func (w *WailsAdapter) Error(message string) {
	w.logger.Log(context.Background(), slog.LevelError, message)
}

func (w *WailsAdapter) Fatal(message string) {
	w.logger.Log(context.Background(), slog.LevelError, message)
	os.Exit(1)
}
