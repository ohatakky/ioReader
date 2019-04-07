package main

import (
	"io"
	"os"
	"strings"
)

// io.Readerからio.Writerにそのまま値を渡したい
// ex.) ファイル読み込んでそのままHHTPで転送したい
func HelpCopy() {
	reader := strings.NewReader("sample data")
	writer, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}

	// writeSize, err := io.Copy(writer, reader)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(writeSize)

	// 8kbのバッファを使ってコピー
	buffer := make([]byte, 8*1024)
	io.CopyBuffer(writer, reader, buffer)
}
