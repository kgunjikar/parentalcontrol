package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log *zap.SugaredLogger

func LoggerInit() error {
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := config.Build()
	if err != nil {
		return err
	}
	Log = logger.Sugar()
	return nil
}
