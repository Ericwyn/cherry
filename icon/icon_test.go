package icon

import (
	"fmt"
	"os"
	"testing"
)

func TestPrintIconSlice(t *testing.T) {
	iconPath := "/home/ericwyn/dev/golang/cherry/icon/icon_64.png"
	iconBytes, err := os.ReadFile(iconPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 输出字节切片，每行 20 个字节
	fmt.Println("const CherryIcon = []byte{")
	fmt.Print("\t")
	for i, b := range iconBytes {
		if i%20 == 0 && i != 0 {
			fmt.Println()
			fmt.Print("\t")
		}
		fmt.Printf("0x%02x, ", b)
	}
	fmt.Println("\n}")
}
