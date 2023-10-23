package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	path := flag.String("path", "no path", "path hack code file")
	flag.Parse()

	parser, err := NewParser(*path)
	if err != nil {
		log.Fatal(err)
	}

	coder := NewCode()
	symbolTable := NewSymbolTable()

	dir, filePath := filepath.Split(*path)
	fileElements := strings.SplitAfterN(filePath, ".", 2)
	filePath = filepath.Join(dir, fileElements[0]+"hack")

	assembler := NewAssembler(parser, coder, symbolTable, filePath)
	if err = assembler.Assemble(); err != nil {
		log.Fatal(err)
	}
}
