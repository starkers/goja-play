package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func StartLogger(debug bool) (zapLog *zap.SugaredLogger) {
	// debug mode is colorful and for human reading
	// https://blog.sandipb.net/2018/05/02/using-zap-simple-use-cases/
	if debug {
		logCfg := zap.Config{
			Encoding:         "console",
			Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",
				// TimeKey:      "time",
				// EncodeTime:   zapcore.ISO8601TimeEncoder,
				CallerKey: "caller",
				// EncodeCaller: zapcore.FullCallerEncoder, // use this if u want the full path to the files in logs
				EncodeCaller: zapcore.ShortCallerEncoder, // instead use package/file:line format
				LevelKey:     "level",
				EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			},
		}
		logger, _ := logCfg.Build()
		return logger.Sugar()
	}

	// non-debug is json
	logConfig := zap.NewProductionConfig()
	// customise the "time" field to be ISO8601
	logConfig.EncoderConfig.TimeKey = "time"
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// main message data into the key "msg"
	logConfig.EncoderConfig.MessageKey = "msg"

	// always send to stdout
	logConfig.OutputPaths = []string{"stdout"}
	logConfig.ErrorOutputPaths = []string{"stdout"}
	zapLogger, err := logConfig.Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return zapLogger.Sugar()
}
