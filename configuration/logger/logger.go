package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// variável que vai guardar as informações do logger
var (
	log *zap.Logger
)

// todas as vezes que a aplicação iniciar, vai adicionar informações na variável de logger.
func init() {
	logConfiguration := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"tmp/auction.log", "stdout"},
		ErrorOutputPaths: []string{"tmp/auction-errors.log", "stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfiguration.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}
