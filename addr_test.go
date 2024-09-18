package addr

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSmart(t *testing.T) {
	data, err := os.ReadFile("data/test")
	if err != nil {
		fmt.Println("读取test文件失败，请检查文件！---", err)
		return
	}
	s := strings.Split(string(data), "\n")

	startT := time.Now() //计算当前时间
	for _, v := range s {
		marshal, err := json.Marshal(Smart(v))
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
	}
	tc := time.Since(startT) //计算解析耗时
	fmt.Printf("time cost = %v\n", tc)
}
