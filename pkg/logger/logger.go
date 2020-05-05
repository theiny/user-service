package logger

import "go.uber.org/zap"

const service = "users"

// New instantiates a new logger instane. Logs for all behaviours (e.g. adding, listing etc) will have the label declared by 'service', and an additional label passed in to describe the subsystem e.g. 'user-add'.
func New(label string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true

	logger, _ := cfg.Build()
	defer logger.Sync() // flushes buffer, if any
	return logger.Sugar().With("service", service, "event", label)
}
