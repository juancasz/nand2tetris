package vmtranslator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CodeWriter struct {
	FileOS *os.File
	*bufio.Writer
	lineCounter int
	filename    string
}

func NewCodeWriter(filename string) (*CodeWriter, error) {
	if ext := filepath.Ext(filename); !strings.EqualFold(ext, ".asm") {
		return nil, fmt.Errorf("unsupported extension %s", ext)
	}

	fileOut, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	filename = strings.ReplaceAll(filepath.Base(filename), filepath.Ext(filename), "")

	return &CodeWriter{
		FileOS:      fileOut,
		Writer:      bufio.NewWriter(fileOut),
		lineCounter: 1,
		filename:    filename,
	}, nil
}

func (c *CodeWriter) WriteInit() error {
	if _, err := c.Writer.WriteString(`
@256
D=A
@SP
M=D
	`); err != nil {
		return err
	}

	if err := c.WriteFunction("OS", 0); err != nil {
		return err
	}

	if err := c.WriteCall("Sys.init", 0); err != nil {
		return err
	}

	return nil
}

func (c *CodeWriter) WriteArithmetic(command string) error {
	assembly, ok := arithmeticCommands[command]
	if !ok {
		return fmt.Errorf("command not found")
	}
	if _, err := c.Writer.WriteString(condSprintf(assembly, c.lineCounter)); err != nil {
		return err
	}
	c.lineCounter++
	return nil
}

func (c *CodeWriter) WritePushPop(command, segment string, index int) error {
	defer func() { c.lineCounter++ }()
	switch command {
	case Push:
		return c.writePush(segment, index)
	case Pop:
		return c.writePop(segment, index)
	default:
		return fmt.Errorf("command not found")
	}
}

func (c *CodeWriter) WriteComment(comment string) error {
	if _, err := c.Writer.WriteString(comment + "\n"); err != nil {
		return err
	}
	return nil
}

func (c *CodeWriter) WriteCall(functionName string, numArgs int) error {
	returnLabel := fmt.Sprintf("%s$ret.%d", functionName, c.lineCounter)

	// push return label
	if _, err := c.Writer.WriteString(fmt.Sprintf(pushStack, returnLabel)); err != nil {
		return err
	}

	// push LCL, ARG, THIS, and THAT
	for _, segment := range []string{memoryAccessDirections[local], memoryAccessDirections[argument], memoryAccessDirections[this], memoryAccessDirections[that]} {
		if _, err := c.Writer.WriteString(fmt.Sprintf(pushStack, segment)); err != nil {
			return err
		}
	}

	// ARG = SP-5-num_args
	if _, err := c.Writer.WriteString(fmt.Sprintf(`
@SP
D=M
@5
D=D-A
@%[1]d
D=D-A
@ARG
M=D
	`, numArgs)); err != nil {
		return err
	}

	// LCL = SP
	if _, err := c.Writer.WriteString(`
@SP
D=M
@LCL
M=D
	`); err != nil {
		return err
	}

	// goto f
	if _, err := c.Writer.WriteString(fmt.Sprintf(`
@%[1]s
0;JMP
	`, functionName)); err != nil {
		return err
	}

	// (return-address)
	if _, err := c.Writer.WriteString(fmt.Sprintf("(%s)\n", returnLabel)); err != nil {
		return err
	}

	c.lineCounter++
	return nil
}

func (c *CodeWriter) WriteFunction(functionName string, numLocals int) error {
	if _, err := c.WriteString(fmt.Sprintf("(%s)", functionName)); err != nil {
		return err
	}

	for i := 1; i <= numLocals; i++ {
		if err := c.writePush(constant, 0); err != nil {
			return err
		}
	}

	return nil
}

func (c *CodeWriter) Close() error {
	return c.FileOS.Close()
}

func (c *CodeWriter) Flush() error {
	return c.Writer.Flush()
}

func (c *CodeWriter) writePush(segment string, index int) error {
	if segment == constant {
		if _, err := c.Writer.WriteString(condSprintf(pushConstant, index)); err != nil {
			return err
		}
		return nil
	}
	if segment == static {
		if _, err := c.Writer.WriteString(condSprintf(pushStatic, index, c.filename)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{local, argument, this, that}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(condSprintf(pushMemoryDynamic, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{pointer, temp}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(condSprintf(pushPointerTemp, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("segment %s not found", segment)
}

func (c *CodeWriter) writePop(segment string, index int) error {
	if segment == constant {
		return fmt.Errorf("can't pop on segment %s", segment)
	}
	if segment == static {
		if _, err := c.Writer.WriteString(condSprintf(popStatic, index, c.filename)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{local, argument, this, that}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(condSprintf(popMemoryDynamic, segment, index, direction)); err != nil {
			return err
		}
		return nil
	}
	if elementInSlice(segment, []string{pointer, temp}) {
		direction, ok := memoryAccessDirections[segment]
		if !ok {
			return fmt.Errorf("segment %s without address", segment)
		}
		if _, err := c.Writer.WriteString(condSprintf(popPointerTemp, segment, index, direction)); err != nil {
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

func condSprintf(f string, args ...interface{}) string {
	// Count the number of placeholders in the format string
	n := strings.Count(f, "%") - (2 * strings.Count(f, "%%"))
	if n < len(args) {
		// If there are more arguments than placeholders, trim the extra arguments
		args = args[:n]
	}
	return fmt.Sprintf(f, args...)
}
