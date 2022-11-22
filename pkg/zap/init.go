package zap

import (
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/gmp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Logger *zap.SugaredLogger

func InitLogger(name string) error {
	path, err := gmp.GetModPath()
	if err != nil {
		log.Fatalln(err)
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	//jsonEncoder := zapcore.NewJSONEncoder(encoderConfig) // alternative encoder for log to file (JSON format)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	filepath := path + viper.GetString("zap.path") + name + ".log"
	logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logLevel := zapcore.DebugLevel
	cores := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(logFile), logLevel),   // log to file
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel), // log to console
	)
	Logger = zap.New(cores, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar() // only trace at ERROR level
	return nil
}
