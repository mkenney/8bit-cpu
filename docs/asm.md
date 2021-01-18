# ASM

## LANG
Compiler syntax.
* `RUN` - execute a pre-defined subroutine

## SYS
System level commands.
* `HLT` - halt clock signal
* `RST` - system reset
* `NOP` - noop, use 1 instruction cycle
* `SLOP` - slow noop, use all instruction cycles
* `BUS [BYTE]` -

## LD (load)
Data register load commands.
* `LD,[RID] [BYTE]` load BYTE into register RID
* `LDG,[RID0] [RID1]` load data in register RID1 into register RID0
* `LDM,[RID] [BYTE]` load data in ROM at address BYTE into register RID
* `LDR,[RID] [BYTE]` load data in RAM at address BYTE into specified register

## OUT
Update BUS_OUT data.
* `OUT [BYTE]` load BYTE into output register
* `OUTG [RID]` load register RID into output register
* `OUTM [BYTE]` load ROM at address BYTE into output register
* `OUTR [BYTE]` load RAM at address BYTE into output register

## STACK
Push and pull from the system stack.

### PUSH
* `PSH [BYTE]` push BYTE onto stack
* `PSHG [RID]` push register RID onto stack
* `PSHM [BYTE]` push data in ROM at address BYTE onto stack
* `PSHR [BYTE]` push data in RAM at address BYTE onto stack

### PULL
* `PUL` - pull from stack, discard
* `PULG [RID]` pull from stack and store data in register BYTE
* `PULR [BYTE]` pull from stack and store data in RAM at address BYTE

## RAM
Write various register or ROM/RAM data into RAM.
* `ST,[ADDR] [DATA]` store DATA in RAM at address ADDR
* `STG,[ADDR] [RID]` store data from register RID in RAM at address ADDR
* `STM,[ADDR0] [ADDR1]` store data from ROM at address ADDR1 in RAM at address ADDR0
* `STR,[ADDR0] [ADDR1]` copy data from RAM at address ADDR1 to ADDR0

## Logical Operations
Various logical operations.

### Math
_Math operations make direct use of register A._

* `ADD [BYTE]` immed - add BYTE
* `ADDG [RID]` reg - add data in register RID
* `ADDM [ADDR]` rom - add data in ROM at address ADDR
* `ADDR [ADDR]` ram - add data in RAM at address ADDR

* `SUB [BYTE]` immed - subtract BYTE
* `SUBG [RID]` reg - subtract data in register RID
* `SUBM [ADDR]` rom - subtract data in ROM at address ADDR
* `SUBR [ADDR]` ram - subtract data in RAM at address ADDR

### Branch
* `JMP [BYTE]` set the program counter to BYTE
* `JMPG [RID]` set the program counter to the data in register RID
* `JMPM [ADDR]` set the program counter to the data in ROM at address ADDR
* `JMPR [ADDR]` set the program counter to the data in RAM at address ADDR
