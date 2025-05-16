package logging

import (
	"fmt"
	"shorty_api/internal/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func NewZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *ZapLogger) getLevel() zapcore.Level {
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.InfoLevel
	}
	return level
}
func (l *ZapLogger) Init() {
	fileName := fmt.Sprintf("%s%s.%s", l.cfg.Logger.FilePath, time.Now().Format("2006-01-02"), "log")

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,
		MaxAge:     20,
		LocalTime:  true,
		MaxBackups: 5,
		Compress:   true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		l.getLevel(),
	)

	l.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func (l *ZapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}

func (l *ZapLogger) Debugf(template string, args ...any) {
	l.logger.Debugf(template, args...)
}

func (l *ZapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Infow(msg, params...)
}

func (l *ZapLogger) Infof(template string, args ...any) {
	l.logger.Infof(template, args...)
}

func (l *ZapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}

func (l *ZapLogger) Warnf(template string, args ...any) {
	l.logger.Warnf(template, args...)
}

func (l *ZapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}

func (l *ZapLogger) Errorf(template string, args ...any) {
	l.logger.Errorf(template, args...)
}

func (l *ZapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Fatalw(msg, params...)
}

func (l *ZapLogger) Fatalf(template string, args ...any) {
	l.logger.Fatalf(template, args...)
}

func prepareLogKeys(extra map[ExtraKey]any, cat Category, sub SubCategory) []any {
	if extra == nil {
		extra = make(map[ExtraKey]any)
	}

	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}
