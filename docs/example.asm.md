# ASM

28 + 14
```ruby
# data is a $label plus a byte:
$d1: 0x1C    # hex 28
$d2: 0b1110  # bin 14

    LDA $d1 # load labeled data into register A (∑ register)
    ADD $d2 # add labeled data to register A (∑ register)
    JMP run # GOTO a label

# labels
run:
    # instructions
    LDA     $d1 # load 28 into register A
    LDB     $d2 # load 14 into register B
    ADD,reg B   # add register B to register A, store result in register A
    OUT,reg A   # load register A (∑ register) data into output register

callSubroutine:
    LDA $d1
    LDB $d2
    RUN addAndOutput

    # loop
    JMP callSubroutine

# subroutines
addAndOutput {
    ADD,reg B
    OUT,reg A
}
```

## SYS
System level commands.
* `HLT` - halt clock signal
* `RST` - system reset
* `NOP` - noop, use 1 instruction cycle
* `SLOP`- slow noop, use all instruction cycles

## REG
Data register load commands.
* `LD[A,X,Y] [BYTE]` load BYTE into specified register
* `LD[A,X,Y],ram [BYTE]` load data in RAM at address BYTE into specified register
* `LD[A,X,Y],reg [BYTE]` load data in register BYTE (0=a,1=x,2=y) into specified register
* `LD[A,X,Y],rom [BYTE]` load data in ROM at address BYTE into specified register

## OUT
Update BUS_OUT data.
* `OUT [BYTE]` send BYTE to output
* `OUT,ram [BYTE]` send data in RAM at address BYTE to output
* `OUT,reg [BYTE]` send data in register BYTE (0=a,1=x,2=y) to output
* `OUT,rom [BYTE]` send data in ROM at address BYTE to output

## STACK
Push and pull from the system stack.

### PUSH
* `PSH [BYTE]` push BYTE onto stack
* `PSH,ram [BYTE]` push data in RAM at address BYTE onto stack
* `PSH,reg [BYTE]` push register BYTE onto stack (0=a,1=x,2=y)
* `PSH,rom [BYTE]` push data in ROM at address BYTE onto stack

### PULL
* `PUL` - pull from stack, discard
* `PUL,ram [BYTE]` pull from stack and store data in RAM at address BYTE
* `PUL,reg [BYTE]` pull from stack and store data in register BYTE (0=a,1=x,2=y)


## RAM management
Write various register or ROM/RAM data into RAM.
* `STA [BYTE]` store data from register A in RAM at address BYTE
* `STB [BYTE]` store data from register B in RAM at address BYTE
* `STX [BYTE]` store data in register X in RAM at address BYTE
* `STY [BYTE]` store data in register Y in RAM at address BYTE
* `STM [BYTE]` store data from RAM in RAM at address BYTE
* `STR [BYTE]` store data from ROM in RAM at address BYTE

## Logical Operations
Various logical operations.

### Math
Math operations make direct use of register A.

* `ADD [BYTE]` immed - add BYTE
* `ADD,ram [BYTE]` ram - add data in RAM at address BYTE
* `ADD,reg [BYTE]` reg - add data in register BYTE (0=a,1=x,2=y)
* `ADD,rom [BYTE]` rom - add data in ROM at address BYTE

* `SUB [BYTE]` immed - subtract BYTE
* `SUB,ram [BYTE]` ram - subtract data in RAM at address BYTE
* `SUB,reg [BYTE]` reg - subtract data in register BYTE (0=a,1=x,2=y)
* `SUB,rom [BYTE]` rom - subtract data in ROM at address BYTE

### Branch
* `JMP [BYTE]` set the program counter to BYTE
* `JMP,ram [BYTE]` set the program counter to the data in RAM at address BYTE
* `JMP,reg [BYTE]` set the program counter to the data in register BYTE (0=a,1=x,2=y)
* `JMP,rom [BYTE]` set the program counter to the data in ROM at address BYTE
