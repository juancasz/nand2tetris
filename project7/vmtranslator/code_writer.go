package vmtranslator

import (
	"bufio"
	"fmt"
	"os"
)

type codeWriter struct {
	FileOS *os.File
	*bufio.Writer
}

func NewCodeWriter(fileName string) (*codeWriter, error) {
	fileOut, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &codeWriter{
		FileOS: fileOut,
		Writer: bufio.NewWriter(fileOut),
	}, nil
}

func (c *codeWriter) WriteArithmetic(command command) error {
	assembly, ok := arithmeticCommands[command.args[0]]
	if !ok {
		return fmt.Errorf("command not found")
	}
	if _, err := c.Writer.WriteString(assembly); err != nil {
		return err
	}
	return nil
}
