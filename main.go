package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"flag"

	"github.com/urkx/LZ77"
	"github.com/urkx/LZWCompress"
)

type Algorithm string

const (
	LZ77 Algorithm = "lz77"
	LZW Algorithm = "lzw"
)

var validAlgorithms = map[string]bool {
	string(LZ77): true,
	string(LZW): true,
}

func main() {
	var algorithmFlag = flag.String("a", string(LZ77), "Algorithm to apply")
	var isCompressFlag = flag.Bool("c", false, "Set if use compression")
	var isDecompressFlag = flag.Bool("d", false, "Set if use decompression")
	var filenameFlag = flag.String("f", "", "Name of the file")

	flag.Parse()

	if *filenameFlag == "" {
		panic("A filename must be provided")
	}

	if *isCompressFlag && *isDecompressFlag {
		panic("Can not compress and decompress at the same time!")
	}

	if !*isCompressFlag && !*isDecompressFlag {
		panic("Must choose compress or decompress")
	}

	if !(validAlgorithms[*algorithmFlag]) {
		panic("Algorithm not supported")
	}

	data, err := os.ReadFile(*filenameFlag)
	if err != nil {
		log.Fatalln("File not found.")
		return
	}

	switch(Algorithm(*algorithmFlag)) {
		case LZ77:
			if *isCompressFlag {
				c := lz77.Compress(string(data), 32000)
				err := lz77.WriteResultFile(*filenameFlag + ".lz77", c)
				if err != nil {
					panic("Error compressing in LZ77")
				}
			}
			if *isDecompressFlag {
				panic("TODO")
			}
		case LZW:
			if *isCompressFlag {
				compressed := lzwcompress.Compress(string(data))
				bytebyff := new(bytes.Buffer)
				binary.Write(bytebyff, binary.NativeEndian, compressed)
				of, err := os.Create(*filenameFlag + ".lzw")
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
			}
			if *isDecompressFlag {
				panic("TODO")
			}
	}
}