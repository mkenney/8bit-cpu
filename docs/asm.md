# ASM

## Labels

## Subroutines

## Instructions

### LANG
Compiler syntax.
* `RUN	[string]` - execute the subroutine identified by `string`
  * push addr onto stack
  * run subroutine steps
  * `}` label signifies JMP to stack value

### SYS
System level commands.
* `HLT` - halt clock signal
* `RST` - system reset
* `NOP` - noop, use 1 instruction cycle
* `SLOP` - slow noop, use all instruction cycles

### LD* (load)
Data register load commands.

#### A register
* `LDA	[BYTE]` Load `BYTE` into register A
* `LDA	[@const]` Load `@const` into register A
* `LDA	[$var]` Load `$var` into register A
* `LDAX` Load data from register X into register A
* `LDAY` Load data from register Y into register A

#### X register
* `LDX	[BYTE]` Load `BYTE` into register X
* `LDX	[@const]` Load `@const` into register X
* `LDX	[$var]` Load `$var` into register X
* `LDXA` Load data from register A into register X
* `LDXY` Load data from register Y into register X

#### Y register
* `LDY	[BYTE]` Load `BYTE` into register Y
* `LDY	[@const]` Load `@const` into register Y
* `LDY	[$var]` Load `$var` into register Y
* `LDYA` Load data from register A into register Y
* `LDYX` Load data from register X into register Y


### OUT register
OUT register load commands.
* `OUT	[BYTE]` load BYTE into output register
* `OUT	[@const]` Load `@const` into output register
* `OUT	[$var]` Load `$var` into output register
* `OUTA` load register A into output register
* `OUTX` load register X into output register
* `OUTY` load register Y into output register

### Logical Operations
Various logical operations.

#### Math
_Math operations make direct use of register A._

* `ADD	[BYTE]` immed - add BYTE
* `ADDG	[RID]` reg - add data in register RID
* `ADDM	[ADDR]` rom - add data in ROM at address ADDR
* `ADDR	[ADDR]` ram - add data in RAM at address ADDR

* `SUB	[BYTE]` immed - subtract BYTE
* `SUBG	[RID]` reg - subtract data in register RID
* `SUBM	[ADDR]` rom - subtract data in ROM at address ADDR
* `SUBR	[ADDR]` ram - subtract data in RAM at address ADDR

#### Branch
* `JMP	[string]` jump to the label identified by `string`
