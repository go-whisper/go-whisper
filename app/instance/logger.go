package instance

import (
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

var onceLogger sync.Once

func Logger() *zap.Logger {
	path := viper.GetString("logger.path")
	onceLogger.Do(func() {
		var level zapcore.Level
		logger = zap.New(zapcore.NewCore(jsonEncoder(), writeSyncer(path), level))
	})
	return logger
}
func writeSyncer(f string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   f,
		MaxSize:    20,
		MaxBackups: 200,
		MaxAge:     15,
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func jsonEncoder() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("20060102 15:04:05.000"))
	}
	return zapcore.NewJSONEncoder(encoder)
}

func consoleEncoder() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("20060102 15:04:05.000"))
	}
	return zapcore.NewConsoleEncoder(encoder)
}
