package f

import (
	"fmt"
	"io"
	"strings"
)

func HelpReader() {
	reader := strings.NewReader("sample data")
	// 全て読み込む
	// buffer, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	panic(err)
	// }

	// 4byte読み込む
	buffer := make([]byte, 4)
	size, err := io.ReadFull(reader, buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(size)
}
