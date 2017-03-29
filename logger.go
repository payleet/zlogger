package zlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wraper struct
type Logger struct {
	logger *zap.Logger
}

// Fileds key/val
type Fileds map[string]interface{}

//NewLogger Returns Logger
func NewLogger() *Logger {

	conf := zap.NewProductionConfig()
	conf.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	conf.Encoding = "json"
	conf.Development = false
	l, _ := conf.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{
		logger: l,
	}
}

// Error logs at error level
func Error(message string, value interface{}) {
	NewLogger().logger.Error(message, valuesToZapFields(value)...)
}

// Info loges at info level
func Info(message string, value interface{}) {
	NewLogger().logger.Info(message, valuesToZapFields(value)...)
}

// Fatal logs at fatal level and calls os.Exit(1) after
func Fatal(message string, value interface{}) {
	NewLogger().logger.Fatal(message, valuesToZapFields(value)...)
}

// Debug logs at debug level
func Debug(message string, value interface{}) {
	NewLogger().logger.Debug(message, valuesToZapFields(value)...)
}

// Panic logs at panic level
func Panic(message string, value interface{}) {
	NewLogger().logger.Panic(message, valuesToZapFields(value)...)
}

// DPanic logs at dpanic level
func DPanic(message string, value interface{}) {
	NewLogger().logger.Panic(message, valuesToZapFields(value)...)
}

// WithError returns a new logger with field error set to error
func (l *Logger) WithError(err error) *Logger {

	return &Logger{
		logger: l.logger.WithOptions(zap.Fields(zap.Error(err))),
	}
}

// WithField returns Logger with Filed key set to value
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		logger: l.logger.WithOptions(zap.Fields(keyValeToZapfield(key, value))),
	}
}

// WithFields returns Logger with Filed key set to value
func (l *Logger) WithFields(f Fileds) *Logger {
	return &Logger{
		logger: l.logger.WithOptions(zap.Fields(fieldsToZapFields(f)...)),
	}
}

// Error logs at error level
func (l *Logger) Error(message string, val ...interface{}) {
	l.logger.Error(message, valuesToZapFields(val)...)
}

// Info logs at info level
func (l *Logger) Info(message string, val ...interface{}) {
	l.logger.Info(message, valuesToZapFields(val)...)
}

// Debug logs at debug level
func (l *Logger) Debug(message string, val ...interface{}) {
	l.logger.Info(message, valuesToZapFields(val)...)
}

// Fatal logs at fatal level then calls os.Exit(1)
func (l *Logger) Fatal(message string, val ...interface{}) {
	l.logger.Fatal(message, valuesToZapFields(val)...)
}

// Panic logs at panic level
func (l *Logger) Panic(message string, val ...interface{}) {
	l.logger.Panic(message, valuesToZapFields(val)...)
}

// Warn logs at warn level
func (l *Logger) Warn(message string, val ...interface{}) {
	l.logger.Warn(message, valuesToZapFields(val)...)
}

// DPanic logs at panic level
func (l *Logger) DPanic(message string, val ...interface{}) {
	l.logger.DPanic(message, valuesToZapFields(val)...)
}

func valuesToZapFields(values ...interface{}) []zapcore.Field {
	if len(values) == 0 {
		return []zapcore.Field{}
	}

	fields := []zapcore.Field{}
	val := values[0].([]interface{})
	for _, v := range val {
		fields = append(fields, keyValeToZapfield("extra", v))
	}

	return fields
}

func fieldsToZapFields(f Fileds) []zapcore.Field {
	zf := []zapcore.Field{}

	for k, v := range f {
		zf = append(zf, keyValeToZapfield(k, v))
	}
	return zf
}

func keyValeToZapfield(key string, val interface{}) zapcore.Field {
	switch val.(type) {
	case error:
		return zap.Error(val.(error))
	case string:
		return zap.String(key, val.(string))
	case bool:
		return zap.Bool(key, val.(bool))
	case int:
		return zap.Int(key, val.(int))
	case int16:
		return zap.Int16(key, val.(int16))
	case int32:
		return zap.Int32(key, val.(int32))
	case int64:
		return zap.Int64(key, val.(int64))
	case uint:
		return zap.Uint(key, val.(uint))
	default:
		return zap.Any(key, val)
	}

}
