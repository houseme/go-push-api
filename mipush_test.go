/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: mipush_test
 * @Version: 1.0.0
 * @Date: 2020/1/31 21:55
 *
 */
package mipush

import (
    "fmt"
    "os"
    "strconv"
    "testing"
    "time"
    
    "go.uber.org/zap"
    
    "github.com/houseme/mipush/util"
)

func TestSend(t *testing.T)  {
    for _, v := range os.Environ() {
        //输出系统所有环境变量的值
        fmt.Println(v)
        t.Log("os env",v)
    }
    
}

func TestSendMessage(t *testing.T) {
    util.SetLevel(util.DebugLevel)
    util.Info("测测谁---")
    timeStart := time.Now().UnixNano()
    fmt.Println("timeStart ", timeStart)
    for i := 0; i < 1; i++ {
        fileName := "MiPush-" + strconv.Itoa(i%10)
        fileName = ""
        fmt.Println(fileName)
        logger := util.Logger(fileName)
        logger.Info("log 初始化成功")
        logger.Info("无法获取网址 ",
            zap.String("url", "http://www.baidu.com"),
            zap.Int("attempt", 3),
            zap.Duration("backoff", time.Second),
            zap.Int("I: ", i))
    }
    endTime := time.Now().UnixNano()
    fmt.Println("endTime ", endTime)
    fmt.Println("end - start time :", endTime-timeStart)
}