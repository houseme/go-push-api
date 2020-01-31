/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: sugarLogger
 * @Version: 1.0.0
 * @Date: 2020/1/31 16:47
 *
 */
package util

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

var logLevel = zap.NewAtomicLevel()

func init() {
    filePath := getFilePath()
    
    fmt.Println("filePath: ", filePath)
    
    w := zapcore.AddSync(&lumberjack.Logger{
        Filename:  filePath,
        MaxSize:   1024, //MB
        LocalTime: true,
        Compress:  true,
    })
    
    config := zap.NewProductionEncoderConfig()
    config.EncodeTime = zapcore.ISO8601TimeEncoder
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(config),
        w,
        logLevel,
    )
    
    loggers := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
    log = loggers.Sugar()
}

type Level int8

const (
    DebugLevel Level = iota - 1
    
    InfoLevel
    
    WarnLevel
    
    ErrorLevel
    
    DPanicLevel
    
    PanicLevel
    
    FatalLevel
)

func SetLevel(level Level) {
    logLevel.SetLevel(zapcore.Level(level))
}

func getCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Info(err)
    }
    fmt.Println(" dir:",dir,"os: ",os.Args)
    return strings.Replace(dir, "\\", "/", -1)
}

func getFilePath() string {
    logfile := getCurrentDirectory() + "/" + getAppName() + ".log"
    return logfile
}

func getAppName() string {
    full := os.Args[0]
    full = strings.Replace(full, "\\", "/", -1)
    splits := strings.Split(full, "/")
    if len(splits) >= 1 {
        name := splits[len(splits)-1]
        name = strings.TrimSuffix(name, ".exe")
        return name
    }
    return ""
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Infof(template string, args ...interface{}) {
    log.Infof(template, args...)
}

func Warn(args ...interface{}) {
    log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
    log.Warnf(template, args...)
}

func Error(args ...interface{}) {
    log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
    log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
    log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
    log.Panicf(template, args...)
}
