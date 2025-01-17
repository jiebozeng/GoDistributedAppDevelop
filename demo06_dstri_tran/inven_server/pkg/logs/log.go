package logs

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
)

type LogType string

const (
	LogType_stdout LogType = "stdout" // stdout
	LogType_file   LogType = "file"   // file
)

type LogEnvType string

const (
	LogEnv_debug   LogEnvType = "debug"
	LogEnv_release LogEnvType = "release"
)

var logTypes []LogType
var logPath string
var logEnv LogEnvType
var saveDays int

var ZapLogger *zap.Logger

func InitLogger(_logTypes []string, _logPath string, _logEnv LogEnvType, _saveDays int) {
	for _, v := range _logTypes {
		if v == string(LogType_file) {
			logTypes = append(logTypes, LogType_file)
		} else if v == string(LogType_stdout) {
			logTypes = append(logTypes, LogType_stdout)
		}
	}
	logPath = _logPath
	logEnv = _logEnv
	saveDays = _saveDays
	var cores []zapcore.Core
	for _, logType := range logTypes {
		if logType == LogType_stdout {
			// Console print
			cores = append(cores, getStdoutCore())
			continue
		}
		if logType == LogType_file {
			// file print
			cores = append(cores, getFileCore())
			continue
		}
	}
	tee := zapcore.NewTee(cores...)
	// logger Output identification of calling code line.
	ZapLogger = zap.New(tee, zap.AddCaller())
	ZapLogger.Info("log server start")
}

func CloseLog() {
	err := ZapLogger.Sync()
	if err != nil {
		err = errors.Wrap(err, "Failed to close the log.")
		if err != nil {
			log.Printf("Failed to close the log: %v", err)
		}
	}
}

func getStdoutCore() zapcore.Core {
	// Restrict log output level, logs of all levels will be printed if >= DebugLevel.
	// Generally, >= ErrorLevel is used in a production environment.
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	// Use JSON format for logging.
	encoder := getConfig()
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), levelEnablerFunc)
}

func getFileCore() zapcore.Core {
	dir := filepath.Dir(logPath)
	_, err := os.Stat(dir)
	if err != nil {
		// Check if the folder exists.
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0o777)
			if err != nil {
				err = errors.Wrap(err, "Failed to create log directory.")
				panic(err)
			}
		}
	}

	err = os.Chmod(dir, 0o777)
	if err != nil {
		err = errors.Wrap(err, "Failed to modify log directory permissions.")
		panic(err)
	}

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    20,
		MaxAge:     saveDays,
		MaxBackups: 100,
		Compress:   true,
	})
	encoder := getConfig()
	return zapcore.NewCore(encoder, zapcore.Lock(writeSyncer), zapcore.DebugLevel)
}

func getConfig() zapcore.Encoder {
	var encoder zapcore.Encoder
	switch logEnv {
	case LogEnv_debug:
		developmentEncoderConfig := zap.NewDevelopmentEncoderConfig()
		//developmentEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		developmentEncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		developmentEncoderConfig.MessageKey = "message"
		developmentEncoderConfig.CallerKey = "caller"
		developmentEncoderConfig.LevelKey = "level"
		developmentEncoderConfig.TimeKey = "created_at"
		encoder = zapcore.NewJSONEncoder(developmentEncoderConfig)
	case LogEnv_release:
		productionEncoderConfig := zap.NewProductionEncoderConfig()
		//productionEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		productionEncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		productionEncoderConfig.MessageKey = "message"
		productionEncoderConfig.CallerKey = "caller"
		productionEncoderConfig.LevelKey = "level"
		productionEncoderConfig.TimeKey = "created_at"
		encoder = zapcore.NewJSONEncoder(productionEncoderConfig)
	default:
		panic(errors.Errorf("unknow env value:%s,please fix it", logEnv))
	}
	return encoder
}
