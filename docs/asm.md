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

### Logical Operations
Various logical operations.

#### Math
_Math operations make direct use of register A._

* `ADDV [value]` -  Add a $const or literal to register A
* `ADDX` -  Add register X to register A
* `ADDY` -  Add register Y to register A
* `SUBV [value]` -  Subtract a $const or literal from register A
* `SUBX` -  Subtract register X from register A
* `SUBY` -  Subtract register Y from register A

#### Branch
* `JMP  [string]` - Load a label index into the program counter
* `JMPV [value]` - Load a $const or literal value into the program counter
* `JMPA` - Load register A into the program counter
* `JMPX` - Load register X into the program counter
* `JMPY` - Load register Y into the program counter
* `JMPS` - Load the last stack value into the program counter

### LD* (load)
Data register load commands.

#### A register
* `LDAV [value]` - Load a $const or literal value into register A
* `LDAX` - Load register X into register A
* `LDAY` - Load register Y into register A

#### X register
* `LDXV [value]` - Load a $const or literal value into register X
* `LDXA` - Load register A into register X
* `LDXY` - Load register Y into register X

#### Y register
 `LDYV` - Load a $const or literal value into register Y
 `LDYA` - Load register A into register Y
 `LDYX` - Load register X into register Y

### Stack operations

#### Push operations
* `PSHV` - Push a $const or literal value onto the stack
* `PSHA` - Push register A onto the stack
* `PSHX` - Push register X onto the stack
* `PSHY` - Push register Y onto the stack

##### Call stack convenience method
* `PSHP` - Push the current program counter onto the stack

#### Pop operations
* `POPA` - Pop a stack value into register A
* `POPX` - Pop a stack value into register X
* `POPY` - Pop a stack value into register Y

##### Call stack convenience method
* `POPP` - Pop a stack value into the program counter

### OUT register
OUT register load commands.

* `OUTV` - Send a value to the output register
* `OUTA` - Send the A register to the output register
* `OUTX` - Send the X register to the output register
* `OUTY` - Send the Y register to the output register
