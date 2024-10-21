package jackanalyzer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// file extensions
const (
	jack = ".jack"
)

type Tockenizer struct {
	fileScanner *bufio.Scanner
	fileOS      *os.File
}

func NewTockenizer(inputFile string) (*Tockenizer, error) {
	if ext := filepath.Ext(inputFile); !strings.EqualFold(ext, jack) {
		return nil, fmt.Errorf("unsupported extension %s", ext)
	}

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	return &Tockenizer{
		fileScanner: fileScanner,
		fileOS:      file,
	}, nil
}

func (t *Tockenizer) hasMoreCommands() bool {
	return t.fileScanner.Scan()
}

func (t *Tockenizer) Close() error {
	return t.fileOS.Close()
}
