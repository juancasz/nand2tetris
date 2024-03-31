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
	filename    string
}

func NewCodeWriter(filename string) (*codeWriter, error) {
	fileOut, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return &codeWriter{
		FileOS:      fileOut,
		Writer:      bufio.NewWriter(fileOut),
		lineCounter: 1,
		filename:    filename,
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
	defer func() { c.lineCounter++ }()
	switch command {
	case push:
		return c.writePush(segment, index)
	case pop:
		return c.writePop(segment, index)
	default:
		return fmt.Errorf("command not found")
	}
}

func (c *codeWriter) writePush(segment string, index int) error {
	if segment == constant {
		if _, err := c.Writer.WriteString(fmt.Sprintf(pushConstant, index)); err != nil {
			return err
		}
		return nil
	}
	if segment == static {
		if _, err := c.Writer.WriteString(fmt.Sprintf(pushStatic, index, c.filename)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{local, argument, this, that}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(fmt.Sprintf(pushMemoryDynamic, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{pointer, temp}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(fmt.Sprintf(pushPointerTemp, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("segment %s not found", segment)
}

func (c *codeWriter) writePop(segment string, index int) error {
	if segment == constant {
		return fmt.Errorf("can't pop on segment %s", segment)
	}
	if segment == static {
		if _, err := c.Writer.WriteString(fmt.Sprintf(popStatic, index, c.filename)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{local, argument, this, that}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(fmt.Sprintf(popMemoryDynamic, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{pointer, temp}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(fmt.Sprintf(popPointerTemp, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func elementInSlice(element string, slice []string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
