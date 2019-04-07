package f

import (
	"fmt"
	"io"
	"os"
)

func OSStdin() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input=%s\n", size, string(buffer))
	}
}
