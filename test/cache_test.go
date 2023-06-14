package test

import (
	"fmt"
	"github.com/VINDA-98/memCache/cache"
	"log"
	"testing"
)

// @Title  test
// @Description  MyGO
// @Author  WeiDa  2023/6/14 17:17
// @Update  WeiDa  2023/6/14 17:17

func TestParseSize(t *testing.T) {
	testCase := []struct {
		input  string
		res    int64
		resStr string
	}{
		{"1KB", 1024, "1KB"},
		{"1kb", 1024, "1KB"},
		{"1MB", 1024 * 1024, "1MB"},
		{"1Mb", 1024 * 1024, "1MB"},
		{"1GB", 1024 * 1024 * 1024, "1GB"},
		{"3GB", 3 * 1024 * 1024 * 1024, "3GB"},
		{"1TB", 1024 * 1024 * 1024 * 1024, "1TB"},
		{"1PB", 1024 * 1024 * 1024 * 1024 * 1024, "1PB"},
		//无法识别，则返回100MB
		{"1P1B", 100 * 1024 * 1024, "100MB"},
	}
	for _, item := range testCase {
		res, resStr := cache.ParseSize(item.input)
		if res != item.res || resStr != item.resStr {
			t.Error(fmt.Sprintf("no pass input: %s res: %d resStr: %s", item.input, item.res, item.resStr))
		}
		log.Println(fmt.Sprintf("no pass input: %s res: %d resStr: %s", item.input, item.res, item.resStr))
	}
}
