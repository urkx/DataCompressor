package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"strings"

	"github.com/urkx/LZWCompress"
)

func main() {
	args := os.Args[1:]
	filename := args[0]
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("File not found.")
		return
	}

	compressed := lzwcompress.Compress(string(data))
	bytebyff := new(bytes.Buffer)
	binary.Write(bytebyff, binary.NativeEndian, compressed)
	of, err := os.Create(strings.Split(filename, ".")[0] + "_compressed.txt")
	if err != nil {
		log.Fatalln("Could not create output file")
		return
	}
	defer of.Close()

	_, errW := of.Write(bytebyff.Bytes())
	if errW != nil {
		log.Fatalln("Error writting to output file")
		return 
	}

	log.Println("Compressed data written to output file")
}