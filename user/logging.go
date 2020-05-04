package user

import "go.uber.org/zap"

const service = "users"

func NewLogger(label string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true

	logger, _ := cfg.Build()
	defer logger.Sync() // flushes buffer, if any
	return logger.Sugar().With("service", service, "action", label)
}
