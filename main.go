package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	doCommand(os.Args[1:])
}

// ConvertFileToLF 将文件内容从 CRLF 转换为 LF
func ConvertFileToLF(filePath string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	// 替换 CRLF 为 LF
	convertedContent := bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

	// 将转换后的内容写回文件
	err = ioutil.WriteFile(filePath, convertedContent, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	return nil
}

// IsShellScript 检查是否是 .sh 文件
func IsShellScript(filePath string) bool {
	return strings.HasSuffix(filePath, ".sh")
}

func doCommand(args []string) {
	if len(args) < 1 || args[0] == "-h" || len(args) >= 2 {
		showHelp()
		return
	}

	dir := os.Args[1]
	//fmt.Println(dir)
	// 遍历目录，查找所有 .sh 文件
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是 .sh 文件
		if !info.IsDir() && IsShellScript(path) {
			fmt.Printf("处理文件: %s      ", path)
			err := ConvertFileToLF(path)
			if err != nil {
				fmt.Printf("转换失败: %s\n", err)
			} else {
				fmt.Printf("转换成功: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %s\n", err)
	}
}

func showHelp() {
	fmt.Println("用法: $ transition [文件夹路径]")
}
