package vmtranslator

import (
	"bufio"
	"fmt"
	"os"
)

type codeWriter struct {
	FileOS *os.File
	*bufio.Writer
	lineCounter int
}

func NewCodeWriter(fileName string) (*codeWriter, error) {
	fileOut, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &codeWriter{
		FileOS:      fileOut,
		Writer:      bufio.NewWriter(fileOut),
		lineCounter: 1,
	}, nil
}

func (c *codeWriter) WriteArithmetic(command string) error {
	assembly, ok := arithmeticCommands[command]
	if !ok {
		return fmt.Errorf("command not found")
	}
	if _, err := c.Writer.WriteString(fmt.Sprintf(assembly, c.lineCounter)); err != nil {
		return err
	}
	c.lineCounter++
	return nil
}

func (c *codeWriter) WritePushPop(command, segment string, index int) error {
	return nil
}
