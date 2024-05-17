package main

import (
	"errors"
	"flag"
	"log"
	"os"
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

	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("path does not exist")
	}
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	var outputName string
	if info.IsDir() {
		var err error
		files, err = listFilesInDir(path)
		if err != nil {
			log.Fatal(err)
		}
		parts := strings.Split(path, "/")
		outputName = filepath.Join(path, parts[len(parts)-1]+".asm")
	} else {
		files = append(files, path)
		dir, file := filepath.Split(path)
		file = strings.ReplaceAll(file, filepath.Ext(file), ".asm")
		outputName = filepath.Join(dir, file)
	}

	codeWriter, err := vmtranslator.NewCodeWriter(outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer codeWriter.Close()
	if info.IsDir() {
		if err := codeWriter.WriteInit(); err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range files {
		parser, err := vmtranslator.NewParser(file)
		if err != nil {
			log.Fatal(err)
		}

		if err := write(parser, codeWriter); err != nil {
			log.Fatal(err)
		}
		parser.Close()
	}
	codeWriter.Flush()
}

func write(parser *vmtranslator.Parser, codeWriter *vmtranslator.CodeWriter) error {
	for {
		err := parser.Advance()
		if errors.Is(err, vmtranslator.ErrNoMoreCommands) {
			break
		}
		if err != nil {
			return err
		}

		commandType, err := parser.CommandType()
		if err != nil {
			return err
		}

		switch commandType {
		case vmtranslator.C_ARITHMETIC:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteArithmetic(arg1); err != nil {
				return err
			}
		case vmtranslator.C_POP:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			arg2, err := parser.Arg2()
			if err != nil {
				return err
			}
			if err := codeWriter.WritePushPop(vmtranslator.Pop, arg1, arg2); err != nil {
				return err
			}
		case vmtranslator.C_PUSH:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			arg2, err := parser.Arg2()
			if err != nil {
				return err
			}
			if err := codeWriter.WritePushPop(vmtranslator.Push, arg1, arg2); err != nil {
				return err
			}
		case vmtranslator.C_COMMENT:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteComment(arg1); err != nil {
				return err
			}
		case vmtranslator.C_LABEL:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteLabel(arg1); err != nil {
				return err
			}
		case vmtranslator.C_GOTO:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteGoTo(arg1); err != nil {
				return err
			}
		case vmtranslator.C_IF:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteIf(arg1); err != nil {
				return err
			}
		case vmtranslator.C_FUNCTION:
			arg1, err := parser.Arg1()
			if err != nil {
				return err
			}
			arg2, err := parser.Arg2()
			if err != nil {
				return err
			}
			if err := codeWriter.WriteFunction(arg1, arg2); err != nil {
				return err
			}
		case vmtranslator.C_RETURN:
			if err := codeWriter.WriteReturn(); err != nil {
				return err
			}
		}
	}

	return nil
}

func listFilesInDir(dirPath string) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.EqualFold(filepath.Ext(fileInfo.Name()), ".vm") {
			fileNames = append(fileNames, filepath.Join(dirPath, fileInfo.Name()))
		}
	}
	return fileNames, nil
}
