package utils

import (
	"bufio"
	"os"
)

/*
NewWriter 使用buffer写文件. 类似可以创建读文件的方法

strings.Builder 高效构建字符串
buf := new(strings.Builder)
buf.WriteString("__")
str := buf.String()
*/
func NewWriter(path string) (*bufio.Writer, func(), error) {
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, func() {}, err
	}
	teardown := func() {
		fileObj.Close()
	}
	writer := bufio.NewWriterSize(fileObj, 4096)
	// writer.Write([]byte{})
	// writer.Flush()
	return writer, teardown, nil
}

// 锁: new(sync.RWMutex)
