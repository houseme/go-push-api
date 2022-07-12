package push

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/buger/jsonparser"
)

func TestSend(t *testing.T) {
	for _, v := range os.Environ() {
		// 输出系统所有环境变量的值
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
	t.Log("timeStart ", timeStart)
}

func TestToken(t *testing.T) {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      {
		"url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460",
		"type": "thumbnail"
		}
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	// You can specify key path by providing arguments to Get function
	t.Log(jsonparser.Get(data, "person", "name", "fullName"))

}
