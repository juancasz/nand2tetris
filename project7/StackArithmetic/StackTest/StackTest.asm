// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/07/StackArithmetic/StackTest/StackTest.vm
// Executes a sequence of arithmetic and logical operations
// on the stack. 

//// push constant 17 ////
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 17 ////
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

//// eq	////
// D=y
@SP
AM=M-1
D=M

// x==y
@SP
AM=M-1
D=M-D

@TRUE3
D;JEQ

@SP
A=M
M=0
@END3
0;JMP

(TRUE3)
@SP
A=M
M=-1

// SP++
(END3)
@SP
M=M+1

//// push constant 17 ////
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 16 ////
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

//// eq	////
// D=y
@SP
AM=M-1
D=M

// x==y
@SP
AM=M-1
D=M-D

@TRUE6
D;JEQ

@SP
A=M
M=0
@END6
0;JMP

(TRUE6)
@SP
A=M
M=-1

// SP++
(END6)
@SP
M=M+1

//// push constant 16 ////
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 17 ////
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

//// eq	////
// D=y
@SP
AM=M-1
D=M

// x==y
@SP
AM=M-1
D=M-D

@TRUE9
D;JEQ

@SP
A=M
M=0
@END9
0;JMP

(TRUE9)
@SP
A=M
M=-1

// SP++
(END9)
@SP
M=M+1

//// push constant 892 ////
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 891 ////
@891
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

@TRUE12
D;JLT

@SP
A=M
M=0
@END12
0;JMP

(TRUE12)
@SP
A=M
M=-1

// SP++
(END12)
@SP
M=M+1	

//// push constant 891 ////
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 892 ////
@892
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

@TRUE15
D;JLT

@SP
A=M
M=0
@END15
0;JMP

(TRUE15)
@SP
A=M
M=-1

// SP++
(END15)
@SP
M=M+1	

//// push constant 891 ////
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 891 ////
@891
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

@TRUE18
D;JLT

@SP
A=M
M=0
@END18
0;JMP

(TRUE18)
@SP
A=M
M=-1

// SP++
(END18)
@SP
M=M+1	

//// push constant 32767 ////
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 32766 ////
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

//// gt	////
// D=y
@SP
AM=M-1
D=M

// x>y
@SP
AM=M-1
D=M-D

@TRUE21
D;JGT

@SP
A=M
M=0
@END21
0;JMP

(TRUE21)
@SP
A=M
M=-1

// SP++
(END21)
@SP
M=M+1

//// push constant 32766 ////
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 32767 ////
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

//// gt	////
// D=y
@SP
AM=M-1
D=M

// x>y
@SP
AM=M-1
D=M-D

@TRUE24
D;JGT

@SP
A=M
M=0
@END24
0;JMP

(TRUE24)
@SP
A=M
M=-1

// SP++
(END24)
@SP
M=M+1

//// push constant 32766 ////
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 32766 ////
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

//// gt	////
// D=y
@SP
AM=M-1
D=M

// x>y
@SP
AM=M-1
D=M-D

@TRUE27
D;JGT

@SP
A=M
M=0
@END27
0;JMP

(TRUE27)
@SP
A=M
M=-1

// SP++
(END27)
@SP
M=M+1

//// push constant 57 ////
@57
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 31 ////
@31
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 53 ////
@53
D=A
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

//// push constant 112 ////
@112
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

//// neg ////
// -y
@SP
A=M-1
M=-M

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

//// push constant 82 ////
@82
D=A
@SP
A=M
M=D
@SP
M=M+1

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

//// not ////
@SP
A=M-1
M=!M	
