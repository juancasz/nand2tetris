// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/SimpleFunction/SimpleFunction.vm
// Performs a simple calculation and returns the result.
(SimpleFunction.test)
//// push constant 0 ////
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 0 ////
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push local 0 ////
@0
D=A
@LCL
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push local 1 ////
@1
D=A
@LCL
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

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

//// not ////
@SP
A=M-1
M=!M	

//// push argument 0 ////
@0
D=A
@ARG
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

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

//// push argument 1 ////
@1
D=A
@ARG
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

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

// FRAME = LCL
// FRAME is R13
@LCL
D=M
@R13
M=D

// RET = *(FRAME-5)
// RET is R14
@5
A=D-A
D=M
@R14
M=D

// *ARG = pop()
@SP
AM=M-1
D=M
@ARG
A=M
M=D

// SP = ARG+1
@ARG
D=M+1
@SP
M=D

@R13
AM=M-1
D=M
@THAT
M=D

@R13
AM=M-1
D=M
@THIS
M=D

@R13
AM=M-1
D=M
@ARG
M=D

@R13
AM=M-1
D=M
@LCL
M=D

@R14
A=M
0;JMP
