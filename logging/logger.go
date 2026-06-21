package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	slogcontext "github.com/PumpkinSeed/slog-context"
	"github.com/SladkyCitron/slogcolor"
)

var (
	defwr               = os.Stdout
	deflvl              = slog.LevelInfo
	lg     *slog.Logger = NewJSONLogger(deflvl)
)

func NewJSONLogger(l slog.Level) *slog.Logger {
	return slog.New(
		slogcontext.NewHandler(
			slog.NewJSONHandler(
				defwr,
				&slog.HandlerOptions{
					Level: l,
				},
			),
		),
	)
}

func NewColoredLogger(l slog.Level) *slog.Logger {
	return slog.New(
		slogcontext.NewHandler(
			slogcolor.NewHandler(
				defwr,
				&slogcolor.Options{
					Level:         l,
					TimeFormat:    slogcolor.DefaultOptions.TimeFormat,
					SrcFileMode:   slogcolor.Nop,
					SrcFileLength: 0,
					MsgPrefix:     slogcolor.DefaultOptions.MsgPrefix,
					MsgColor:      slogcolor.DefaultOptions.MsgColor,
					MsgLength:     slogcolor.DefaultOptions.MsgLength,
					NoColor:       slogcolor.DefaultOptions.NoColor,
					LevelTags:     slogcolor.DefaultOptions.LevelTags,
				},
			),
		),
	)
}

func WithValue(ctx context.Context, key string, value any) context.Context {
	return slogcontext.WithValue(ctx, key, value)
}

func GetLogger() *slog.Logger {
	return lg
}

func SetDefaultLogger(l *slog.Logger) {
	lg = l
}

func Debug(ctx context.Context, msg string) {
	lg.DebugContext(ctx, msg)
}

func Debugf(ctx context.Context, format string, args ...any) {
	Debug(ctx, fmt.Sprintf(format, args...))
}

func DebugKV(ctx context.Context, msg string, attrs ...any) {
	lg.DebugContext(ctx, msg, attrs...)
}

func Info(ctx context.Context, msg string) {
	lg.InfoContext(ctx, msg)
}

func Infof(ctx context.Context, format string, args ...any) {
	Info(ctx, fmt.Sprintf(format, args...))
}

func InfoKV(ctx context.Context, msg string, attrs ...any) {
	lg.InfoContext(ctx, msg, attrs...)
}

func Warn(ctx context.Context, msg string) {
	lg.WarnContext(ctx, msg)
}

func Warnf(ctx context.Context, format string, args ...any) {
	Warn(ctx, fmt.Sprintf(format, args...))
}

func WarnKV(ctx context.Context, msg string, attrs ...any) {
	lg.WarnContext(ctx, msg, attrs...)
}

func Error(ctx context.Context, msg string) {
	lg.ErrorContext(ctx, msg)
}

func Errorf(ctx context.Context, format string, args ...any) {
	Error(ctx, fmt.Sprintf(format, args...))
}

func ErrorKV(ctx context.Context, msg string, attrs ...any) {
	lg.ErrorContext(ctx, msg, attrs...)
}

func Fatal(ctx context.Context, msg string) {
	lg.ErrorContext(ctx, msg)
	os.Exit(1)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	Error(ctx, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func FatalKV(ctx context.Context, msg string, attrs ...any) {
	lg.ErrorContext(ctx, msg, attrs...)
	os.Exit(1)
}
