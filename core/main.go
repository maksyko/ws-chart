package core

import (
	"go.uber.org/zap"
)

var zapLogger, err = zap.NewDevelopment()

// Logger is the global logger used by application
var Logger = zapLogger.Sugar()
