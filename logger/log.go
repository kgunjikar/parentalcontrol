package logger

import (
	"encoding/json"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func LoggerInit() error {
	rawJSONConfig := []byte{
		"level": "info",
		"encoderConfig": {
			"timeEncoder": "iso8601",
		},
	}
	config := zap.Config()
	if err := json.Unmarshal(rawJSONConfig, &config); err != nil {
		return err
	}
	logger, err := config.Build()
	if err != nil {
		return err
	}
	Log = logger.SugaredLogger()
}
