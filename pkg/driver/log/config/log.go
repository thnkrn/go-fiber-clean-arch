package config

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ProvidZapLogger() (*zap.Logger, error) {
	rawJSON := []byte(`{
		"level": "info",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)
	var zcfg zap.Config
	if err := json.Unmarshal(rawJSON, &zcfg); err != nil {
		panic(err)
	}

	zcfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger := zap.Must(zcfg.Build())
	defer logger.Sync()

	return logger, nil
}
