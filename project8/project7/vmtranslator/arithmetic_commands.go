package vmtranslator

// arithmetic commands
const (
	add = "add"
	sub = "sub"
	neg = "neg"
	eq  = "eq"
	gt  = "gt"
	lt  = "lt"
	and = "and"
	or  = "or"
	not = "not"
)

var arithmeticCommands = map[string]string{
	add: `
//// add ////
// D=y
@SP
AM=M-1
D=M

// x=x+y
@SP
AM=M-1
M=M+D

// SP++
@SP
M=M+1
`,
	sub: `
//// sub ////
// D=y
@SP
AM=M-1
D=M

// x=x+y
@SP
AM=M-1
M=M-D

// SP++
@SP
M=M+1
`,
	neg: `
//// neg ////
// -y
@SP
A=M-1
M=-M
`,
	eq: `
//// eq	////
// D=y
@SP
AM=M-1
D=M

// x==y
@SP
AM=M-1
D=M-D

@TRUE%[1]d
D;JEQ

@SP
A=M
M=0
@END%[1]d
0;JMP

(TRUE%[1]d)
@SP
A=M
M=-1

// SP++
(END%[1]d)
@SP
M=M+1
`,
	gt: `
//// gt	////
// D=y
@SP
AM=M-1
D=M

// x>y
@SP
AM=M-1
D=M-D

@TRUE%[1]d
D;JGT

@SP
A=M
M=0
@END%[1]d
0;JMP

(TRUE%[1]d)
@SP
A=M
M=-1

// SP++
(END%[1]d)
@SP
M=M+1
`,
	lt: `
//// lt	////
// D=y
@SP
AM=M-1
D=M

// x<y
@SP
AM=M-1
D=M-D

@TRUE%[1]d
D;JLT

@SP
A=M
M=0
@END%[1]d
0;JMP

(TRUE%[1]d)
@SP
A=M
M=-1

// SP++
(END%[1]d)
@SP
M=M+1	
`,
	and: `
//// and ////
// D=y
@SP
A=M-1
D=M

// x && y
A=A-1
M=M&D

// SP--
@SP
M=M-1
`,
	or: `
//// or ////
// D=y
@SP
A=M-1
D=M

// x | y
A=A-1
M=M|D

// SP--
@SP
M=M-1
`,
	not: `
//// not ////
@SP
A=M-1
M=!M	
`,
}
