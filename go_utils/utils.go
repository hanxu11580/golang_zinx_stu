package go_utils

import (
	"bufio"
	"fmt"
	"os"
)

// 将文本一行一行写入
func writeLinetext(outputPath string, textSeq []string) {
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, lineStr := range textSeq {
		fmt.Fprintln(w, lineStr)
	}
	w.Flush()
}
