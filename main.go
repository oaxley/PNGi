package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	List "github.com/oaxley/PNGi/pkg/linkedlist"
	PNG "github.com/oaxley/PNGi/pkg/pngchunk"
)

func main() {
	// linkedlist
	var list List.LinkedList

	// read arguments from the command line
	flag.Parse()

	// read the original file
	fi, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatalf("Unable to open '%s' for reading.", flag.Args()[0])
	}
	defer fi.Close()

	// read the PNG Header
	var header PNG.Header
	header.Read(fi)

	// read all the chunks from the input file
	done := false
	for !done {
		var chunk PNG.Chunk
		err = chunk.Read(fi)
		if err != nil {
			done = true
		} else {
			list.Add(chunk)
		}
	}

	fmt.Printf("%d chunk(s) read from the source file.\n", list.Count())

	// write the data to the output file
	fo, err := os.Create(flag.Args()[1])
	if err != nil {
		log.Fatalf("Unable to create '%s' for writing.", flag.Args()[1])
	}
	defer fo.Close()

	// write the header
	header.Write(fo)

	// write all the chunks
	count := 0
	for node := range list.Elements() {
		chunk := node.Data.(PNG.Chunk)
		chunk.Write(fo)
		count = count + 1
	}
	fmt.Printf("%d chunk(s) written to the target file.\n", count)
}
