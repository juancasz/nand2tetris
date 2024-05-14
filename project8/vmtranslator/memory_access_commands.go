package vmtranslator

// memory access commands
const (
	Pop  = "pop"
	Push = "push"
)

// function commands and branching logic
const (
	Label    = "label"
	GoTo     = "goto"
	IfGoTo   = "if-goto"
	Function = "function"
	Return   = "return"
	Call     = "call"
)

// memory segments
const (
	local    = "local"
	argument = "argument"
	this     = "this"
	that     = "that"
	pointer  = "pointer"
	temp     = "temp"
	constant = "constant"
	static   = "static"
)

// memory segments directions
var memoryAccessDirections = map[string]string{
	local:    "LCL",
	argument: "ARG",
	this:     "THIS",
	that:     "THAT",
	pointer:  "3",
	temp:     "5",
}

// assembly memory access commands
const (
	pushStatic = `
//// push static %[1]d ////
@%[2]s.%[1]d
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushConstant = `
//// push constant %[1]d ////
@%[1]d
D=A
@SP
A=M
M=D
@SP
M=M+1
`
	pushMemoryDynamic = `
//// push %[1]s %[2]d ////
@%[2]d
D=A
@%[3]s
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
`

	pushPointerTemp = `
//// push %[1]s %[2]d ////
@%[2]d
D=A
@%[3]s
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushStack = `
@%[1]s
D=A
@SP
A=M
M=D
@SP
M=M+1
`
	popStatic = `
//// pop static %[1]d ////
@SP
AM=M-1
D=M
@%[2]s.%[1]d
M=D
`
	popMemoryDynamic = `
//// pop %[1]s %[2]d ////
@%[2]d
D=A
@%[3]s
A=M
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
`
	popPointerTemp = `
//// pop %[1]s %[2]d ////
@%[2]d
D=A
@%[3]s
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D	
`
)
