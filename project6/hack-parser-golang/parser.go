package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrNoMoreCommands = errors.New("no more commands")
var ErrInvalidLCommand = errors.New("L command must end with )")
var ErrInvalidCCommand = errors.New("invalid C command")
var ErrCommandNoSymbols = errors.New("command does not have symbols")
var ErrNotCCommand = errors.New("not C command")
var ErrCommandNotAvailable = errors.New("command not available")

const (
	A_COMMAND string = "A_COMMAND"
	C_COMMAND string = "C_COMMAND"
	L_COMMAND string = "L_COMMAND"
)

type parser struct {
	path           string
	fileScanner    *bufio.Scanner
	FileOS         *os.File
	currentCommand command
}

type command struct {
	line        string
	commandType string
}

func NewParser(inputFile string) (*parser, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	p := &parser{
		path:        inputFile,
		fileScanner: fileScanner,
		FileOS:      file,
	}

	return p, nil
}

func (p *parser) Restore() (*os.File, error) {
	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	p.fileScanner = fileScanner
	p.fileScanner.Split(bufio.ScanLines)
	p.FileOS = file
	return file, nil
}

func (p *parser) Advance() error {
	var line string
	var hasMoreCommands bool
	for {
		if hasMoreCommands = p.fileScanner.Scan(); !hasMoreCommands {
			return ErrNoMoreCommands
		}

		line = p.fileScanner.Text()
		if line == "" {
			continue
		}

		if len(line) >= 2 && string(line[:2]) == "//" {
			continue
		}

		break
	}

	if strings.Contains(line, "//") {
		indexComment := strings.Index(line, "//")
		line = strings.ReplaceAll(line, string(line[indexComment:]), "")
	}

	line = strings.ReplaceAll(line, " ", "")
	p.currentCommand.line = line
	return nil
}

func (p *parser) CommandType() (commandType string, err error) {
	if isLine := p.checkIfLineAvailable(); !isLine {
		err = ErrCommandNotAvailable
		return
	}

	switch firstCharacter := string(p.currentCommand.line[0]); firstCharacter {
	case "@":
		commandType = A_COMMAND
		p.currentCommand.commandType = commandType
		return
	case "(":
		if err = p.parseLCommand(); err != nil {
			return
		}
		commandType = L_COMMAND
		p.currentCommand.commandType = commandType
		return
	default:
		if err = p.parseCCommand(); err != nil {
			return
		}
		commandType = C_COMMAND
		p.currentCommand.commandType = commandType
		return
	}
}

func (p *parser) Symbol() (string, error) {
	if isLine := p.checkIfLineAvailable(); !isLine {
		return "", ErrCommandNotAvailable
	}

	switch p.currentCommand.commandType {
	case A_COMMAND:
		return string(p.currentCommand.line[1:]), nil
	case L_COMMAND:
		return string(p.currentCommand.line[1 : len(p.currentCommand.line)-1]), nil
	default:
		return "", ErrCommandNoSymbols
	}
}

func (p *parser) Dest() (string, error) {
	if isLine := p.checkIfLineAvailable(); !isLine {
		return "", ErrCommandNotAvailable
	}

	if !strings.EqualFold(p.currentCommand.commandType, C_COMMAND) {
		return "", ErrNotCCommand
	}

	equalPosition := strings.Index(p.currentCommand.line, "=")
	if equalPosition == -1 {
		return "", nil
	}

	return string(p.currentCommand.line[:equalPosition]), nil
}

func (p *parser) Comp() (string, error) {
	if isLine := p.checkIfLineAvailable(); !isLine {
		return "", ErrCommandNotAvailable
	}

	if !strings.EqualFold(p.currentCommand.commandType, C_COMMAND) {
		return "", ErrNotCCommand
	}

	equalPosition := strings.Index(p.currentCommand.line, "=")
	endPosition := strings.Index(p.currentCommand.line, ";")

	if equalPosition != -1 && endPosition != -1 {
		return string(p.currentCommand.line[equalPosition+1 : endPosition]), nil
	} else if equalPosition != -1 && endPosition == -1 {
		return string(p.currentCommand.line[equalPosition+1:]), nil
	} else if equalPosition == -1 && endPosition != -1 {
		return string(p.currentCommand.line[:endPosition]), nil
	}

	return p.currentCommand.line, nil
}

func (p *parser) Jump() (string, error) {
	if isLine := p.checkIfLineAvailable(); !isLine {
		return "", ErrCommandNotAvailable
	}

	if !strings.EqualFold(p.currentCommand.commandType, C_COMMAND) {
		return "", ErrNotCCommand
	}

	endPosition := strings.Index(p.currentCommand.line, ";")
	if endPosition != -1 {
		return string(p.currentCommand.line[endPosition+1:]), nil
	}

	return "", nil
}

func (p *parser) checkIfLineAvailable() bool {
	if len(p.currentCommand.line) == 0 {
		return false
	}
	return true
}

func (p *parser) parseLCommand() error {
	if lastCharacter := string(p.currentCommand.line[len(p.currentCommand.line)-1]); lastCharacter != ")" {
		return ErrInvalidLCommand
	}
	return nil
}

func (p *parser) parseCCommand() error {
	if equalSigns := strings.Count(p.currentCommand.line, "="); equalSigns > 1 {
		return ErrInvalidCCommand
	}
	if endCommands := strings.Count(p.currentCommand.line, ";"); endCommands > 1 {
		return ErrInvalidCCommand
	}

	return nil
}
