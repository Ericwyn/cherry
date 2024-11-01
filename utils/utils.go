package utils

import (
	"bytes"
	"cherry/log"
	"errors"
	"fmt"
	"golang.design/x/clipboard"
	"image"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GetClipboardImageData() ([]byte, error) {
	data := clipboard.Read(clipboard.FmtImage)
	if len(data) == 0 {
		return nil, errors.New("clipboard is empty")
	}
	return data, nil
}

func WriteUrlToClipboard(result string) {
	clipboard.Write(clipboard.FmtText, []byte(result))
}

func DetectImageFormat(data []byte) string {
	if len(data) < 4 {
		return "unknown"
	}

	// 使用标准库的 Decode 函数来判断格式
	imgType := "unknown"
	img, format, err := image.Decode(bytes.NewReader(data))
	if err == nil && img != nil {
		imgType = format
	}

	// 如果无法解码，则使用文件头判断
	if imgType == "unknown" {
		switch {
		case bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}):
			imgType = "jpeg"
		case bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47}):
			imgType = "png"
		case bytes.HasPrefix(data, []byte{0x47, 0x49, 0x46}):
			imgType = "gif"
		default:
			imgType = "png" // 默认返回 png
		}
	}

	return imgType
}

func OpenSysDirectory(path string) error {

	log.I("open runDir: " + path)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		path = filepath.FromSlash(path)
		cmd = exec.Command("explorer", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
