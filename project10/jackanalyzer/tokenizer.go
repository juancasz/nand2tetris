package jackanalyzer

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

// file extensions
const (
	jack = ".jack"
)

var ErrNoMoreCommands = errors.New("no more commands")

type Tockenizer struct {
	fileScanner    *bufio.Scanner
	fileOS         *os.File
	indexCharacter int
	currentLine    []rune
	currentToken   token
}

type TokenType int

const (
	KEYWORD TokenType = iota
	SYMBOL
	IDENTIFIER
	INT_CONST
	STRING_CONST
)

func (t TokenType) String() string {
	return []string{"KEYWORD", "SYMBOL", "IDENTIFIER", "INT_CONST", "STRING_CONST"}[t]
}

var keywords = []string{
	"class",
	"method",
	"function",
	"constructor",
	"int",
	"boolean",
	"char",
	"void",
	"var",
	"static",
	"field",
	"let",
	"do",
	"if",
	"else",
	"while",
	"return",
	"true",
	"false",
	"null",
	"this",
}

var symbols = []string{
	"{",
	"}",
	"(",
	")",
	"[",
	"]",
	".",
	",",
	";",
	"+",
	"-",
	"*",
	"/",
	"&",
	"|",
	"<",
	">",
	"=",
	"~",
}

type token struct {
	character string
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
	var tokenType TokenType
	defer func() {
		t.currentToken.tokenType = tokenType
	}()
	if t.isKeyword() {
		tokenType = KEYWORD
		return tokenType, nil
	}
	if t.isSymbol() {
		tokenType = SYMBOL
		return tokenType, nil
	}
	if t.isIntConst() {
		tokenType = INT_CONST
		return tokenType, nil
	}
	if t.isStringConst() {
		tokenType = STRING_CONST
		return tokenType, nil
	}
	if t.isIdentifier() {
		tokenType = IDENTIFIER
		return tokenType, nil
	}
	return TokenType(0), fmt.Errorf("invalid token type")
}

func (t *Tockenizer) Advance() error {
	if t.currentLine == nil || t.indexCharacter >= len(t.currentLine) {
		if err := t.advanceLine(); err != nil {
			return err
		}
	}

	t.space()

	tokenType, err := t.TokenType()
	if err != nil {
		return err
	}

	switch tokenType {
	case KEYWORD:
		keyword, err := t.Keyword()
		if err != nil {
			return err
		}
		t.currentToken.character = keyword
	case SYMBOL:
		symbol, err := t.Symbol()
		if err != nil {
			return err
		}
		t.currentToken.character = symbol
	case INT_CONST:
		intConst, err := t.IntVal()
		if err != nil {
			return err
		}
		t.currentToken.character = intConst
	case STRING_CONST:
		stringConst, err := t.StringVal()
		if err != nil {
			return err
		}
		t.currentToken.character = stringConst
	case IDENTIFIER:
		identifier, err := t.Identifier()
		if err != nil {
			return err
		}
		t.currentToken.character = identifier
	}

	log.Printf("%s - %s", t.currentToken.character, t.currentToken.tokenType.String())

	return nil
}

func (t *Tockenizer) advanceLine() error {
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

		break
	}

	t.currentLine = []rune(line)
	t.indexCharacter = 0
	return nil
}

func (t *Tockenizer) Keyword() (string, error) {
	if t.currentToken.tokenType != KEYWORD {
		return "", fmt.Errorf("token type is not keyword")
	}
	init := t.indexCharacter
	for unicode.IsLetter(t.currentLine[t.indexCharacter]) {
		t.indexCharacter++
	}
	return string(t.currentLine[init:t.indexCharacter]), nil
}

func (t *Tockenizer) Symbol() (string, error) {
	if t.currentToken.tokenType != SYMBOL {
		return "", fmt.Errorf("token type is not symbol")
	}

	symbol := t.lookAhead(1)
	t.indexCharacter++
	return symbol, nil
}

func (t *Tockenizer) IntVal() (string, error) {
	if t.currentToken.tokenType != INT_CONST {
		return "", fmt.Errorf("token type is not int const")
	}

	init := t.indexCharacter
	for {
		if _, err := strconv.Atoi(string(t.currentLine[t.indexCharacter])); err != nil {
			break
		}
		t.indexCharacter++
	}

	return string(t.currentLine[init:t.indexCharacter]), nil
}

func (t *Tockenizer) StringVal() (string, error) {
	if t.currentToken.tokenType != STRING_CONST {
		return "", fmt.Errorf("token type is not string const")
	}

	t.indexCharacter++ // skip leading '"'
	init := t.indexCharacter
	for t.lookAhead(1) != `"` {
		t.indexCharacter++
	}
	t.indexCharacter++ // skip trailing '"'
	return string(t.currentLine[init : t.indexCharacter-1]), nil
}

func (t *Tockenizer) Identifier() (string, error) {
	if t.currentToken.tokenType != IDENTIFIER {
		return "", fmt.Errorf("token type is not identifier")
	}

	init := t.indexCharacter
	for unicode.IsLetter(t.currentLine[t.indexCharacter]) || string(t.currentLine[t.indexCharacter]) == "_" {
		t.indexCharacter++
	}
	return string(t.currentLine[init:t.indexCharacter]), nil
}

func (t *Tockenizer) space() {
	for t.lookAhead(1) == " " {
		t.indexCharacter++
	}
}

func (t *Tockenizer) isKeyword() bool {
	for _, keyword := range keywords {
		if t.lookAhead(len(keyword)) == keyword {
			return true
		}
	}

	return false
}

func (t *Tockenizer) isSymbol() bool {
	for _, symbol := range symbols {
		if t.lookAhead(1) == symbol {
			return true
		}
	}
	return false
}

func (t *Tockenizer) isIntConst() bool {
	if _, err := strconv.Atoi(t.lookAhead(1)); err == nil {
		return true
	}
	return false
}

func (t *Tockenizer) isStringConst() bool {
	if t.lookAhead(1) == `"` {
		return true
	}
	return false
}

func (t *Tockenizer) isIdentifier() bool {
	for _, character := range []rune(t.lookAhead(1)) {
		if unicode.IsLetter(character) || string(character) == "_" {
			return true
		}
	}
	return false
}

func (t *Tockenizer) lookAhead(n int) string {
	//log.Println("index", t.indexCharacter)
	if t.indexCharacter+n <= len(t.currentLine) {
		return string(t.currentLine[t.indexCharacter : t.indexCharacter+n])
	}
	return string(t.currentLine[t.indexCharacter:])
}

func (t *Tockenizer) Close() error {
	return t.fileOS.Close()
}
