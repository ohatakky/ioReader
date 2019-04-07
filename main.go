package main

import "os"

func main() {
	// HelpReader()
	// HelpCopy()
	// OSStdin()
	// FileCopy()
	// HTTPCopyStdin()
	// BufferRead()
	// ReadPos()
	// EndiannessConvert()

	file, err := os.Open("lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
