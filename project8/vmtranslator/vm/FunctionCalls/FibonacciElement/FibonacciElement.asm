
@256
D=A
@SP
M=D
(OS)
@Sys.init$ret.1
D=A
@SP
A=M
M=D
@SP
M=M+1

@LCL
D=A
@SP
A=M
M=D
@SP
M=M+1

@ARG
D=A
@SP
A=M
M=D
@SP
M=M+1

@THIS
D=A
@SP
A=M
M=D
@SP
M=M+1

@THAT
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
D=M
@5
D=D-A
@0
D=D-A
@ARG
M=D

@SP
D=M
@LCL
M=D

@Sys.init
0;JMP
(Sys.init$ret.1)
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/FibonacciElement/Main.vm
// Computes the n'th element of the Fibonacci series, recursively.
// n is given in argument[0].  Called by the Sys.init function 
// (part of the Sys.vm file), which also pushes the argument[0] 
// parameter before this code starts running.
(Main.fibonacci)
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

//// push constant 2 ////
@2
D=A
@SP
A=M
M=D
@SP
M=M+1

//// lt	////
// D=y
@SP
AM=M-1
D=M

// x<y
@SP
AM=M-1
D=M-D

@TRUE4
D;JLT

@SP
A=M
M=0
@END4
0;JMP

(TRUE4)
@SP
A=M
M=-1

// SP++
(END4)
@SP
M=M+1	

@SP
AM=M-1
D=M
@Main.fibonacci$IF_TRUE
D;JNE

@Main.fibonacci$IF_FALSE
0;JMP
(Main.fibonacci$IF_TRUE)
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
(Main.fibonacci$IF_FALSE)
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

//// push constant 2 ////
@2
D=A
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

//// push constant 1 ////
@1
D=A
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
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/FibonacciElement/Sys.vm
// Pushes a constant, say n, onto the stack, and calls the Main.fibonacii
// function, which computes the n'th element of the Fibonacci series.
// Note that by convention, the Sys.init function is called "automatically" 
// by the bootstrap code.
(Sys.init)
//// push constant 4 ////
@4
D=A
@SP
A=M
M=D
@SP
M=M+1
(Sys.init$WHILE)
@Sys.init$WHILE
0;JMP
