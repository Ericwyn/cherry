package log

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	E("测试", "错误", errors.New("error"))
}
