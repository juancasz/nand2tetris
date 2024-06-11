
//// SP = 256 ////
@256
D=A
@SP
M=D

//// FUNCTION ////
(OS)

//// CALL Sys.init 0 ////

//// push Sys.init$ret.0
@Sys.init$ret.0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

//// ARG = SP-5-num_args ////
@SP
D=M
@5
D=D-A
@0
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Sys.init
0;JMP

//// (return-address) ////
(Sys.init$ret.0)
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/StaticsTest/Sys.vm
// Tests that different functions, stored in two different 
// class files, manipulate the static segment correctly. 

//// FUNCTION ////
(Sys.init)

//// push constant 6 ////
@6
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 8 ////
@8
D=A
@SP
A=M
M=D
@SP
M=M+1

//// CALL Class1.set 2 ////

//// push Class1.set$ret.1
@Class1.set$ret.1
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

//// ARG = SP-5-num_args ////
@SP
D=M
@5
D=D-A
@2
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Class1.set
0;JMP

//// (return-address) ////
(Class1.set$ret.1)

//// pop temp 0 ////
@0
D=A
@5
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D	

//// push constant 23 ////
@23
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 15 ////
@15
D=A
@SP
A=M
M=D
@SP
M=M+1

//// CALL Class2.set 2 ////

//// push Class2.set$ret.2
@Class2.set$ret.2
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

//// ARG = SP-5-num_args ////
@SP
D=M
@5
D=D-A
@2
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Class2.set
0;JMP

//// (return-address) ////
(Class2.set$ret.2)

//// pop temp 0 ////
@0
D=A
@5
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D	

//// CALL Class1.get 0 ////

//// push Class1.get$ret.3
@Class1.get$ret.3
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

//// ARG = SP-5-num_args ////
@SP
D=M
@5
D=D-A
@0
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Class1.get
0;JMP

//// (return-address) ////
(Class1.get$ret.3)

//// CALL Class2.get 0 ////

//// push Class2.get$ret.4
@Class2.get$ret.4
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

//// ARG = SP-5-num_args ////
@SP
D=M
@5
D=D-A
@0
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Class2.get
0;JMP

//// (return-address) ////
(Class2.get$ret.4)
(Sys.init$WHILE)
@Sys.init$WHILE
0;JMP
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/StaticsTest/Class2.vm
// Stores two supplied arguments in static[0] and static[1].

//// FUNCTION ////
(Class2.set)

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

//// pop static 0 ////
@SP
AM=M-1
D=M
@StaticsTest.0
M=D

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

//// pop static 1 ////
@SP
AM=M-1
D=M
@StaticsTest.1
M=D

//// push constant 0 ////
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// RETURN ////
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

// Restore THAT of the caller
@R13
AM=M-1
D=M
@THAT
M=D

// Restore THIS of the caller
@R13
AM=M-1
D=M
@THIS
M=D

// Restore ARG of the caller
@R13
AM=M-1
D=M
@ARG
M=D

// Restore LCL of the caller
@R13
AM=M-1
D=M
@LCL
M=D

//// goto RET ////
@R14
A=M
0;JMP
// Returns static[0] - static[1].

//// FUNCTION ////
(Class2.get)

//// push static 0 ////
@StaticsTest.0
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push static 1 ////
@StaticsTest.1
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

//// RETURN ////
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

// Restore THAT of the caller
@R13
AM=M-1
D=M
@THAT
M=D

// Restore THIS of the caller
@R13
AM=M-1
D=M
@THIS
M=D

// Restore ARG of the caller
@R13
AM=M-1
D=M
@ARG
M=D

// Restore LCL of the caller
@R13
AM=M-1
D=M
@LCL
M=D

//// goto RET ////
@R14
A=M
0;JMP
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/08/FunctionCalls/StaticsTest/Class1.vm
// Stores two supplied arguments in static[0] and static[1].

//// FUNCTION ////
(Class1.set)

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

//// pop static 0 ////
@SP
AM=M-1
D=M
@StaticsTest.0
M=D

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

//// pop static 1 ////
@SP
AM=M-1
D=M
@StaticsTest.1
M=D

//// push constant 0 ////
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// RETURN ////
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

// Restore THAT of the caller
@R13
AM=M-1
D=M
@THAT
M=D

// Restore THIS of the caller
@R13
AM=M-1
D=M
@THIS
M=D

// Restore ARG of the caller
@R13
AM=M-1
D=M
@ARG
M=D

// Restore LCL of the caller
@R13
AM=M-1
D=M
@LCL
M=D

//// goto RET ////
@R14
A=M
0;JMP
// Returns static[0] - static[1].

//// FUNCTION ////
(Class1.get)

//// push static 0 ////
@StaticsTest.0
D=M
@SP
A=M
M=D
@SP
M=M+1

//// push static 1 ////
@StaticsTest.1
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

//// RETURN ////
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

// Restore THAT of the caller
@R13
AM=M-1
D=M
@THAT
M=D

// Restore THIS of the caller
@R13
AM=M-1
D=M
@THIS
M=D

// Restore ARG of the caller
@R13
AM=M-1
D=M
@ARG
M=D

// Restore LCL of the caller
@R13
AM=M-1
D=M
@LCL
M=D

//// goto RET ////
@R14
A=M
0;JMP
