package f

import (
	"io"
	"os"
)

func FileCopy() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
