package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger

func InitLogger(path string) error {
	config := zap.NewDevelopmentEncoderConfig()
	//jsonEncoder := zapcore.NewJSONEncoder(config) // alternative encoder for log to file (JSON format)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(logFile), logLevel),   // log to file
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel), // log to console
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar() // only trace at ERROR level
	return nil
}
