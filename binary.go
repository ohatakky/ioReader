package main

import (
	"io"
	"os"
	"strings"
)

// import "io"

// func ReadLimit() {
// 	lReader := io.LimitReader(reader, 16)
// }

func ReadPos() {
	reader := strings.NewReader("sample data\n")
	sectionReader := io.NewSectionReader(reader, 3, 5)
	io.Copy(os.Stdout, sectionReader)
}
