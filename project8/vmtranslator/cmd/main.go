package main

import (
	"errors"
	"flag"
	"log"
	"path/filepath"
	"strings"
	"vmtranslator"
)

func main() {
	flag.Parse()
	var path string
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	} else {
		log.Fatal("no file provided")
	}

	parser, err := vmtranslator.NewParser(path)
	if err != nil {
		log.Fatal(err)
	}
	defer parser.Close()

	dir, file := filepath.Split(path)
	file = strings.ReplaceAll(file, filepath.Ext(file), ".asm")

	codeWriter, err := vmtranslator.NewCodeWriter(filepath.Join(dir, file))
	if err != nil {
		log.Fatal(err)
	}
	defer codeWriter.Close()

	if err := write(parser, codeWriter); err != nil {
		log.Fatal(err)
	}
}

func write(parser *vmtranslator.Parser, codeWriter *vmtranslator.CodeWriter) error {
	var commandType vmtranslator.CommandType
	var arg1 string
	var arg2 int
	var err error
	for {
		err = parser.Advance()
		if errors.Is(err, vmtranslator.ErrNoMoreCommands) {
			break
		}
		if err != nil {
			return err
		}

		commandType, err = parser.CommandType()
		if err != nil {
			return err
		}

		arg1, err = parser.Arg1()
		if err != nil {
			return err
		}

		switch commandType {
		case vmtranslator.C_ARITHMETIC:
			if err = codeWriter.WriteArithmetic(arg1); err != nil {
				return err
			}
		case vmtranslator.C_POP:
			arg2, err = parser.Arg2()
			if err != nil {
				return err
			}
			if err = codeWriter.WritePushPop(vmtranslator.Pop, arg1, arg2); err != nil {
				return err
			}
		case vmtranslator.C_PUSH:
			arg2, err = parser.Arg2()
			if err != nil {
				return err
			}
			if err = codeWriter.WritePushPop(vmtranslator.Push, arg1, arg2); err != nil {
				return err
			}
		case vmtranslator.C_COMMENT:
			if err = codeWriter.WriteComment(arg1); err != nil {
				return err
			}
		}
	}

	return codeWriter.Flush()
}
