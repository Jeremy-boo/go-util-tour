package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger zap log
var Logger *zap.Logger

// InitLogger Initialization log
func InitLogger(param *NewLogger) (err error) {
	param = buildInitParams(param)
	// set logger level info
	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(param.Level))
	if err != nil {
		return err
	}
	// init zao core
	core := zapcore.NewCore(getEncoder(), getLoggerWriter(param.FileName,
		param.MaxSize, param.MaxBackups, param.MaxAge), l)
	// set the log initialization field
	filed := zap.Fields(zap.String("ServiceName", param.ServiceName))
	Logger = zap.New(core, zap.AddCaller(), filed)
	return nil
}

func buildInitParams(param *NewLogger) *NewLogger {
	if param.Level == "" {
		param.Level = LogLevel
	}
	if param.ServiceName == "" {
		param.ServiceName = "defaultService"
	}
	if param.FileName == "" {
		param.FileName = FileName
	}
	if param.MaxSize == 0 {
		param.MaxSize = MaxSize
	}
	if param.MaxAge == 0 {
		param.MaxAge = MaxAge
	}
	if param.MaxBackups == 0 {
		param.MaxBackups = MaxBackups
	}
	return param
}

// getLoggerWriter use lumberJackLogger to  cut logs by size
func getLoggerWriter(fileName string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	jackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}
	return zapcore.AddSync(jackLogger)
}

// getEncoder set zap default encode
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
