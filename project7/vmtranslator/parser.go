package vmtranslator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ErrNoMoreCommands = errors.New("no more commands")

type parser struct {
	path           string
	fileScanner    *bufio.Scanner
	FileOS         *os.File
	currentCommand command
	regexCommands  *regexp.Regexp
}

type command struct {
	line        string
	commandType CommandType
	args        []string
}

func NewParser(inputFile string) (*parser, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	regexCommands, err := regexp.Compile(`\S+`)
	if err != nil {
		return nil, err
	}

	return &parser{
		path:          inputFile,
		fileScanner:   fileScanner,
		FileOS:        file,
		regexCommands: regexCommands,
	}, nil
}

func (p *parser) hasMoreCommands() bool {
	return p.fileScanner.Scan()
}

func (p *parser) Advance() error {
	var line string
	for {
		if !p.hasMoreCommands() {
			return ErrNoMoreCommands
		}

		line = p.fileScanner.Text()
		if line == "" {
			continue
		}

		break
	}

	p.currentCommand.line = line
	return nil
}

// CommandType returns the type of the command and set the args and the type of the current command in memory
func (p *parser) CommandType() (CommandType, error) {
	var commandType CommandType
	defer func() {
		p.currentCommand.commandType = commandType
	}()

	if len(p.currentCommand.line) >= 2 && string(p.currentCommand.line[:2]) == "//" {
		commandType = C_COMMENT
		return commandType, nil
	}

	if strings.Contains(p.currentCommand.line, "//") {
		indexComment := strings.Index(p.currentCommand.line, "//")
		p.currentCommand.line = strings.ReplaceAll(p.currentCommand.line, string(p.currentCommand.line[indexComment:]), "")
	}

	args := p.regexCommands.FindAllString(p.currentCommand.line, -1)
	if args == nil {
		return 0, fmt.Errorf("command not found")
	}
	p.currentCommand.args = args

	if _, ok := arithmeticCommands[args[0]]; ok {
		commandType = C_ARITHMETIC
		return commandType, nil
	}
	if args[0] == push {
		commandType = C_PUSH
		return commandType, nil
	}
	if args[0] == pop {
		commandType = C_POP
		return commandType, nil
	}

	return 0, fmt.Errorf("command type not found")
}

func (p *parser) Arg1() (string, error) {
	switch p.currentCommand.commandType {
	case C_ARITHMETIC:
		return p.currentCommand.args[0], nil
	default:
		if len(p.currentCommand.args) < 2 {
			return "", fmt.Errorf("missing arg1")
		}
		return p.currentCommand.args[1], nil
	}
}

func (p *parser) Arg2() (int, error) {
	if p.currentCommand.commandType != C_POP && p.currentCommand.commandType != C_PUSH {
		return 0, fmt.Errorf("missing arg2")
	}
	if len(p.currentCommand.args) < 3 {
		return 0, fmt.Errorf("missing arg2")
	}
	arg2, err := strconv.Atoi(p.currentCommand.args[2])
	if err != nil {
		return 0, err
	}
	return arg2, nil
}
