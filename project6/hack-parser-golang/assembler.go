package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assembler struct {
	*parser
	*code
	*symbolTable
	outPath string
}

func NewAssembler(parser *parser, code *code, symbolTable *symbolTable, outPath string) *assembler {
	return &assembler{
		parser:      parser,
		code:        code,
		symbolTable: symbolTable,
		outPath:     outPath,
	}
}

func (a *assembler) Assemble() error {
	var err error
	if err = a.firtPass(); err != nil {
		return err
	}
	if err = a.secondPass(); err != nil {
		return err
	}
	if err = a.output(); err != nil {
		return err
	}
	return nil
}

func (a *assembler) firtPass() error {
	defer a.parser.FileOS.Close()
	counter := 0
	for {
		err := a.parser.Advance()
		if errors.Is(err, ErrNoMoreCommands) {
			break
		}
		if err != nil {
			return err
		}

		commandType, err := a.parser.CommandType()
		if err != nil {
			return err
		}

		switch commandType {
		case A_COMMAND:
			counter++
		case C_COMMAND:
			counter++
		case L_COMMAND:
			symbol, err := a.parser.Symbol()
			if err != nil {
				return err
			}
			a.symbolTable.AddEntry(symbol, counter)
		}
	}
	return nil
}

func (a *assembler) secondPass() error {
	fileOs, err := a.parser.Restore()
	if err != nil {
		return err
	}
	defer fileOs.Close()

	ramAdress := 16
	for {
		err := a.parser.Advance()
		if errors.Is(err, ErrNoMoreCommands) {
			break
		}
		if err != nil {
			return err
		}

		commandType, err := a.parser.CommandType()
		if err != nil {
			return err
		}

		if commandType == A_COMMAND {
			symbol, err := a.parser.Symbol()
			if err != nil {
				return err
			}

			isDigit := false
			if _, err = strconv.Atoi(symbol); err == nil {
				isDigit = true
			}

			if !isDigit && !a.symbolTable.Contains(symbol) {
				a.symbolTable.AddEntry(symbol, ramAdress)
				ramAdress++
			}
		}
	}

	return nil
}

func (a *assembler) output() error {
	fileOs, err := a.parser.Restore()
	if err != nil {
		return err
	}
	defer fileOs.Close()

	fileOut, err := os.Create(a.outPath)
	if err != nil {
		return err
	}

	writerOut := bufio.NewWriter(fileOut)

	for {
		err := a.parser.Advance()
		if errors.Is(err, ErrNoMoreCommands) {
			break
		}
		if err != nil {
			return err
		}

		commandType, err := a.parser.CommandType()
		if err != nil {
			return err
		}

		switch commandType {
		case A_COMMAND:
			symbol, err := a.parser.Symbol()
			if err != nil {
				return err
			}

			isDigit := false
			digit, err := strconv.Atoi(symbol)
			if err == nil {
				isDigit = true
			}

			var binSymbol string
			if isDigit {
				binSymbol = strconv.FormatInt(int64(digit), 2)
			} else {
				address := a.symbolTable.GetAddress(symbol)
				binSymbol = strconv.FormatInt(int64(address), 2)
			}

			aCommand := strings.Repeat("0", 16-len(binSymbol)) + binSymbol
			_, err = writerOut.WriteString(fmt.Sprintf("%s\n", aCommand))
			if err != nil {
				return err
			}
		case C_COMMAND:
			destMnem, err := a.parser.Dest()
			if err != nil {
				return err
			}
			dest, err := a.code.Dest(destMnem)
			if err != nil {
				return err
			}

			compMnem, err := a.parser.Comp()
			if err != nil {
				return err
			}
			comp, err := a.code.Comp(compMnem)
			if err != nil {
				return err
			}

			jumpMnem, err := a.parser.Jump()
			if err != nil {
				return err
			}
			jump, err := a.code.Jump(jumpMnem)
			if err != nil {
				return err
			}

			cCommand := "111" + comp + dest + jump
			_, err = writerOut.WriteString(fmt.Sprintf("%s\n", cCommand))
			if err != nil {
				return err
			}
		}

	}

	if err = writerOut.Flush(); err != nil {
		return err
	}

	return nil
}
