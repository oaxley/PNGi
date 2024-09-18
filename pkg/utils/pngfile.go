package utils

import (
	"log"
	"os"

	"github.com/oaxley/PNGi/pkg/linkedlist"
	PNG "github.com/oaxley/PNGi/pkg/pngchunk"
)

type PNGFile struct {
	Header  PNG.Header
	ImgData linkedlist.LinkedList
}

// read a PNG file and create the associated linked list
func (file *PNGFile) Read(pngfile string) int {

	// open the PNG file
	fi, err := os.Open(pngfile)
	if err != nil {
		log.Fatalf("Unable to open '%s' for reading", pngfile)
	}
	defer fi.Close()

	// read the PNG header
	file.Header.Read(fi)

	// read all the chunks
	done := false
	for !done {
		var chunk PNG.Chunk
		err = chunk.Read(fi)
		if err != nil {
			done = true
		} else {
			file.ImgData.Add(chunk)
		}
	}

	return file.ImgData.Count()
}

// write a PNG file from the header and linked list
func (file *PNGFile) Write(pngfile string) int {
	// open the new file
	fo, err := os.Create(pngfile)
	if err != nil {
		log.Fatalf("Unable to open '%s' for writing", pngfile)
	}
	defer fo.Close()

	// write the header
	file.Header.Write(fo)

	// write all the chunks
	count := 0
	for node := range file.ImgData.Elements() {
		chunk := node.Data.(PNG.Chunk)
		chunk.Write(fo)
		count = count + 1
	}

	return count
}
