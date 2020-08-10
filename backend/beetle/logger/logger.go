package logger

import (
	"github.com/ShintaNakama/yukibu-channel/backend/beetle/env"
	"go.uber.org/zap"
)

// NewLogger returns logger
func NewLogger(e *env.ENV) *zap.Logger {
	var logger *zap.Logger
	if env.IsProd(e) {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)

	return logger
}
