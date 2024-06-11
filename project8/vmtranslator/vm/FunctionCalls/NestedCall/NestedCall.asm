
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
// Sys.vm for NestedCall test.
// Sys.init()
//
// Calls Sys.main() and stores return value in temp 1.
// Does not return.  (Enters infinite loop.)

//// FUNCTION ////
(Sys.init)

//// push constant 4000 ////
@4000
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

//// push constant 5000 ////
@5000
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

//// CALL Sys.main 0 ////

//// push Sys.main$ret.1
@Sys.main$ret.1
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
@Sys.main
0;JMP

//// (return-address) ////
(Sys.main$ret.1)

//// pop temp 1 ////
@1
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
(Sys.init$LOOP)
@Sys.init$LOOP
0;JMP
// Sys.main()
//
// Sets locals 1, 2 and 3, leaving locals 0 and 4 unchanged to test
// default local initialization to 0.  (RAM set to -1 by test setup.)
// Calls Sys.add12(123) and stores return value (135) in temp 0.
// Returns local 0 + local 1 + local 2 + local 3 + local 4 (456) to confirm
// that locals were not mangled by function call.

//// FUNCTION ////
(Sys.main)

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

//// push constant 0 ////
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

//// push constant 4001 ////
@4001
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

//// push constant 5001 ////
@5001
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

//// push constant 200 ////
@200
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop local 1 ////
@1
D=A
@LCL
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

//// push constant 40 ////
@40
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop local 2 ////
@2
D=A
@LCL
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

//// push constant 6 ////
@6
D=A
@SP
A=M
M=D
@SP
M=M+1

//// pop local 3 ////
@3
D=A
@LCL
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

//// push constant 123 ////
@123
D=A
@SP
A=M
M=D
@SP
M=M+1

//// CALL Sys.add12 1 ////

//// push Sys.add12$ret.2
@Sys.add12$ret.2
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
@1
D=D-A
@ARG
M=D

//// LCL = SP ////
@SP
D=M
@LCL
M=D

//// goto f ////
@Sys.add12
0;JMP

//// (return-address) ////
(Sys.add12$ret.2)

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

//// push local 2 ////
@2
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

//// push local 3 ////
@3
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

//// push local 4 ////
@4
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
// Sys.add12(int n)
//
// Returns n+12.

//// FUNCTION ////
(Sys.add12)

//// push constant 4002 ////
@4002
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

//// push constant 5002 ////
@5002
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

//// push constant 12 ////
@12
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
