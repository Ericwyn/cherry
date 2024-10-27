package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

var (
	logFileName = "cherry-picup.log"
	saveLog     = true
	mu          sync.Mutex // to ensure thread safety for log writes
)

func SetRunnerPath(runPath string) {
	logFileName = path.Join(runPath, logFileName)
}

func E(v ...any) {
	logStr := fmt.Sprint(getTime(), " [ERRR] ", fmt.Sprint(v...))
	fmt.Println(logStr)
	saveLogToLocal(logStr)
}

func D(v ...any) {
	logStr := fmt.Sprint(getTime(), " [DBUG] ", fmt.Sprint(v...))
	fmt.Println(logStr)
	saveLogToLocal(logStr)
}

func I(v ...any) {
	logStr := fmt.Sprint(getTime(), " [INFO] ", fmt.Sprint(v...))
	fmt.Println(logStr)
	saveLogToLocal(logStr)
}

func getTime() string {
	// 格式化时间为指定格式
	return time.Now().Format(time.DateTime)
}

func saveLogToLocal(logStr string) {
	if !saveLog {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("[ERRR] Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	// Use a buffered writer for efficient writing
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(logStr + "\n")
	if err != nil {
		log.Printf("[ERRR] Failed to write log: %v", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("[ERRR] Failed to flush log: %v", err)
	}
}
