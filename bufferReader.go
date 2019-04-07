package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func BufferRead() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// 空のバッファ
	// var buffer1 bytes.Buffer
	// バイト列で初期化
	buffer2 := bytes.NewBuffer([]byte{})
	// 文字列で初期化
	// buffer3 := bytes.NewBufferString("初期文字列")

	io.Copy(buffer2, file)
	fmt.Println(buffer2)
}
