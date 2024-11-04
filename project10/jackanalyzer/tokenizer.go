package jackanalyzer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// file extensions
const (
	jack = ".jack"
)

var ErrNoMoreCommands = errors.New("no more commands")

type Tockenizer struct {
	fileScanner  *bufio.Scanner
	fileOS       *os.File
	currentToken token
}

type TokenType int

const (
	KEYWORD TokenType = iota
	SYMBOL
	IDENTIFIER
	INT_CONST
	STRING_CONST
)

var keywords = map[string]struct{}{
	"class":       {},
	"method":      {},
	"function":    {},
	"constructor": {},
	"int":         {},
	"boolean":     {},
	"char":        {},
	"void":        {},
	"var":         {},
	"static":      {},
	"field":       {},
	"let":         {},
	"do":          {},
	"if":          {},
	"else":        {},
	"while":       {},
	"return":      {},
	"true":        {},
	"false":       {},
	"null":        {},
	"this":        {},
}

var symbols = map[string]struct{}{
	"{": {},
	"}": {},
	"(": {},
	")": {},
	"[": {},
	"]": {},
	".": {},
	",": {},
	";": {},
	"+": {},
	"-": {},
	"*": {},
	"/": {},
	"&": {},
	"|": {},
	"<": {},
	">": {},
	"=": {},
	"~": {},
}

type token struct {
	line      string
	tokenType TokenType
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

func (t *Tockenizer) hasMoreTokens() bool {
	return t.fileScanner.Scan()
}

func (t *Tockenizer) TokenType() (TokenType, error) {
	if t.isKeyword() {
		return KEYWORD, nil
	}
	return TokenType(0), nil
}

func (t *Tockenizer) Advance() error {
	var line string
	for {
		if !t.hasMoreTokens() {
			return ErrNoMoreCommands
		}

		line = t.fileScanner.Text()
		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}

		// Remove multi-line comments
		if startIdx := strings.Index(line, "/*"); startIdx != -1 {
			if endIdx := strings.Index(line[startIdx:], "*/"); endIdx != -1 {
				line = line[:startIdx] + line[startIdx+endIdx+2:]
			} else {
				line = line[:startIdx] // If no closing, remove from startIdx to end
			}
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for _, character := range strings.Fields(line) {
			t.currentToken.line = character
		}

		break
	}

	t.currentToken.line = line
	return nil
}

func (t *Tockenizer) isKeyword() bool {
	for keyword := range keywords {
		if strings.Contains(t.currentToken.line, keyword) {
			return true
		}
	}
	return false
}

func (t *Tockenizer) Close() error {
	return t.fileOS.Close()
}
