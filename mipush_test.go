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
)

func TestSend(t *testing.T) {
    for _, v := range os.Environ() {
        //输出系统所有环境变量的值
        fmt.Println(v)
        t.Log("os env", v)
    }
    
}

func TestSendMessage(t *testing.T) {
    timeStart := time.Now().UnixNano()
    fmt.Println("timeStart ", timeStart)
    
    for i := 0; i < 50; i++ {
        fileName := "MiPush" + strconv.Itoa(i%2)
        fmt.Println(fileName)
    }
    endTime := time.Now().UnixNano()
    fmt.Println("endTime ", endTime)
    fmt.Println("end - start time :", endTime-timeStart)
    fmt.Println(os.Getenv("MI_PUSH_LOGGING_MODE"))
}
