package application

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger menginisialisasi logger dengan konfigurasi tertentu
func InitLogger() {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			MessageKey:     "msg",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// SyncLogger melakukan sinkronisasi dan memastikan log dibersihkan dengan baik
func SyncLogger() {
	defer Logger.Sync()
}
