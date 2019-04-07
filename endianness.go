package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// バイナリ解析ではエンディアン変換が必要
// ビッグエンディアン：ネットワーク (データの先頭byteを小さいアドレス)
// リトルエンディアン：CPU (データの先頭byteを大きいアドレス)
// 前から読むか、後ろから読むか
func EndiannessConvert() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// ビッグエンディアンのデータを現在のCPUに合わせて変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
