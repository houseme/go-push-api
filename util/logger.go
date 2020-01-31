/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2020/1/31 15:45
 *
 */
package util

import (
    "fmt"
    "os"
    "strconv"
    "sync"
    "time"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var now time.Time
var loggerMap map[string]*zap.Logger

func Logger(fileName string) *zap.Logger {
    loggerMap = InitLoggerMap()
    logger, ok := loggerMap[fileName]
    fmt.Println("loggerMap ok ", ok)
    if !ok {
        logger = InitLogger(fileName)
    }
    now = time.Now()
    if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
        logger = InitLogger(fileName)
    }
    return logger
}

//Initialize object processing
func InitLoggerMap() map[string]*zap.Logger {
    if loggerMap == nil {
        once.Do(func() {
            loggerMap = make(map[string]*zap.Logger)
        })
    }
    return loggerMap
}

func InitLogger(fileName string) *zap.Logger {
    dayTime := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day())
    hook := lumberjack.Logger{
        Filename:   "./logs/" + dayTime + "/" + fileName + ".log", // 日志文件路径
        MaxSize:    128,                                           // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: 30,                                            // 日志文件最多保存多少个备份
        MaxAge:     7,                                             // 文件最多保存多少天
        Compress:   true,                                          // 是否压缩
    }
    
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "line",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }
    
    // 设置日志级别
    atomicLevel := zap.NewAtomicLevel()
    atomicLevel.SetLevel(zap.InfoLevel)
    
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
        atomicLevel,                                                                     // 日志级别
    )
    
    // 开启开发模式，堆栈跟踪
    caller := zap.AddCaller()
    // 开启文件及行号
    development := zap.Development()
    // 设置初始化字段
    filed := zap.Fields(zap.String("serviceName", "MiPush"))
    // 构造日志
    logger := zap.New(core, caller, development, filed)
    fmt.Println("Logger time: ", time.Now().Unix())
    loggerMap[fileName] = logger
    return logger
}
