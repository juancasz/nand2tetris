// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/07/MemoryAccess/PointerTest/PointerTest.vm
// Executes pop and push commands using the 
// pointer, this, and that segments.

//// push constant 3030 ////
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop pointer 0 ////
@0
D=A
@3
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D	

//// push constant 3040 ////
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop pointer 1 ////
@1
D=A
@3
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D	

//// push constant 32 ////
@32
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop this 2 ////
@2
D=A
@THIS
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

//// push constant 46 ////
@46
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop that 6 ////
@6
D=A
@THAT
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

//// push pointer 0 ////
@0
D=A
@3
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push pointer 1 ////
@1
D=A
@3
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

//// push this 2 ////
@2
D=A
@THIS
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

//// push that 6 ////
@6
D=A
@THAT
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
