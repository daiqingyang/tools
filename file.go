package tools

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"strings"
)

// 确保某一行文本存在于指定的文件中
func LineInFile(content []byte, filePath string, mode os.FileMode) (err error) {
	// 检测文件，不存在就新建并写入内容
	// 文件存在的话，再过滤内容是否存在，内容不存在时进行追加
	var fInfo fs.FileInfo
	var contain bool
	var fs *os.File
	fInfo, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.WriteFile(filePath, content, mode)
			if err != nil {
				return
			} else {
			}
		} else {
			return
		}
	} else {
		if fInfo.IsDir() {
			err = errors.New(filePath + " is  directory")
			return
		}
	}
	contain, err = Grep(string(content), filePath)
	if err != nil {
		return
	}
	if !contain {
		fs, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0600)
		if err != nil {
			return
		}
		defer fs.Close()
		_, err = fs.Write(content)
		if err != nil {
			return
		}
	}
	os.Chmod(filePath, mode)
	return
}
func Grep(in, fileName string) (contain bool, err error) {
	in = strings.Trim(in, "\n")
	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	scaner := bufio.NewScanner(f)
	for scaner.Scan() {
		text := scaner.Text()
		if strings.Contains(text, in) {
			contain = true
			break
		}
	}
	return
}

// go arch : linux uname
var ArchMap = map[string]string{
	"arm64": "aarch64",
	"amd64": "x86_64",
}
