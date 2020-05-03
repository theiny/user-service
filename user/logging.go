package user

import "go.uber.org/zap"

const service = "users"

func (s *server) NewLogger() {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true

	logger, _ := cfg.Build()
	defer logger.Sync() // flushes buffer, if any
	s.Log = logger.Sugar().With(
		"service", service,
	)
}
