package internal

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logging struct {
	instance *zap.Logger
}

var (
	once        = sync.Once{}
	logInstance *Logging
)

func GetLogger(config *Config) *Logging {

	once.Do(func() {
		logInstance = &Logging{
			instance: Logger(config),
		}
	})
	return logInstance

}

func Logger(config *Config) *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeLevel:    shortLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		config.DefaultLogLevel,
	)
	return zap.New(core, zap.AddCaller())
}

func shortLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("[DBG]")
	case zapcore.InfoLevel:
		enc.AppendString("[INF]")
	case zapcore.WarnLevel:
		enc.AppendString("[WRN]")
	case zapcore.ErrorLevel:
		enc.AppendString("[ERR]")
	default:
		enc.AppendString("[???]")
	}
}

func (l *Logging) logInfo(msg string, fields ...zap.Field) {
	l.instance.Info(msg, fields...)
}

func (l *Logging) logDebug(msg string, fields ...zap.Field) {
	l.instance.Debug(msg, fields...)
}
func (l *Logging) logError(msg string, fields ...zap.Field) {
	l.instance.Error(msg, fields...)
}
func (l *Logging) logFatal(msg string, fields ...zap.Field) {
	l.instance.Fatal(msg, fields...)
}

func (l *Logging) logErrorException(err error, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	l.instance.Error(msg, fields...)
}
