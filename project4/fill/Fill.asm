// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
(START)
    // 8K registers in screen
    @8192
    D=A
    @counterRegistersToPaint
    M=D

    // Set the value of var currentScreenRegister to the first register of the screen
    @SCREEN
    D=A
    @currentScreenRegister
    M=D

    // keyboard. 0=> No press // 1=> press
    @KBD
    D=M

    // If press => set paint color to black
    @SETBLACK
    D; JGT

    // If no press => set paint color to white
    @SETWHITE
    D; JEQ

// Set var color to 0
(SETWHITE)
    @color
    M=0

    @PAINT
    0;JMP

// Set var color to 1
(SETBLACK)
    @color
    M=-1

    @PAINT
    0;JMP

// Paint the screem
(PAINT)
    // Set D to value of color
    @color
    D=M

    
    // paint
    @currentScreenRegister
    A=M
    M=D
    
    //move to next register
    @currentScreenRegister
    M=M+1

    // count painted register
    @counterRegistersToPaint
    M=M-1
    D=M

    // if counterRegistersToPaint is 0 => go to start program
    @START
    D; JEQ

    // if counterRegistersToPaint is not 0, continue painting
    @PAINT
    0; JMP