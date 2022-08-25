package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createDir 创建缓存目录
func createDir() {
	rootPath := utils.GetRootPath()
	runtimeDir := fmt.Sprintf("%s%s", rootPath, global.CONFIG.Zap.RuntimeDir)
	if ok, _ := utils.PathExists(runtimeDir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", runtimeDir)
		_ = os.Mkdir(runtimeDir, os.ModePerm)
	}

	logDir := fmt.Sprintf("%s%s", rootPath, global.CONFIG.Zap.LogDir)
	if ok, _ := utils.PathExists(logDir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logDir)
		_ = os.Mkdir(logDir, os.ModePerm)
	}
}

func Zap() (logger *zap.Logger) {
	// 创建缓存/日志 目录
	createDir()

	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("%s%s/%s/server_debug.log", utils.GetRootPath(), global.CONFIG.Zap.LogDir, time.Now().Format("20060102")), debugPriority),
		getEncoderCore(fmt.Sprintf("%s%s/%s/server_info.log", utils.GetRootPath(), global.CONFIG.Zap.LogDir, time.Now().Format("20060102")), infoPriority),
		getEncoderCore(fmt.Sprintf("%s%s/%s/server_warn.log", utils.GetRootPath(), global.CONFIG.Zap.LogDir, time.Now().Format("20060102")), warnPriority),
		getEncoderCore(fmt.Sprintf("%s%s/%s/server_error.log", utils.GetRootPath(), global.CONFIG.Zap.LogDir, time.Now().Format("20060102")), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    2,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,   // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if global.CONFIG.Zap.LogInConsole { // 是否打印在控制台 Stdout
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := getWriteSyncer(fileName) // 使用lumberjack进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
