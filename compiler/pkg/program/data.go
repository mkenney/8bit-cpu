package program

import "github.com/bdlm/errors/v2"

// bitMap
var bitMap []string = []string{

	// system
	"HLT",  // Halt system clock signal
	"RST",  // Reset all system registers
	"RUN",  // Execute a subroutine. Is encoded as a JMPV instruction
	"NOP",  // No-op, use 1 instruction cycle
	"SLOP", // Slow no-op, use 16 instruction cycles

	// math
	"ADDV", // Add a $const or literal to register A
	"ADDX", // Add register X to register A
	"ADDY", // Add register Y to register A
	"SUBV", // Subtract a $const or literal from register A
	"SUBX", // Subtract register X from register A
	"SUBY", // Subtract register Y from register A

	// branch
	"JMP",  // `JMP [label]`  - Jump to [label]:       Load a label index into the program counter
	"JMPV", // `JMPV [value]` - Jump to [value]:       Load a $const or literal value into the program counter
	"JMPA", // `JMPA`         - Jump to register A:    Load register A into the program counter
	"JMPX", // `JMPX`         - Jump to register X:    Load register X into the program counter
	"JMPY", // `JMPY`         - Jump to register Y:    Load register Y into the program counter
	"JMPS", // `JMPS`         - Jump to stack pointer: Load the last stack value into the program counter

	// data
	"LDAV", // Load a $const or literal value into register A
	"LDAX", // Load register X into register A
	"LDAY", // Load register Y into register A

	"LDXV", // Load a $const or literal value into register X
	"LDXA", // Load register A into register X
	"LDXY", // Load register Y into register X

	"LDYV", // Load a $const or literal value into register Y
	"LDYA", // Load register A into register Y
	"LDYX", // Load register X into register Y

	// stack
	"PSHV", // Push a $const or literal value onto the stack
	"PSHA", // Push register A onto the stack
	"PSHX", // Push register X onto the stack
	"PSHY", // Push register Y onto the stack
	"PSHP", // Push the current program counter onto the stack

	"POPA", // Pop a stack value into register A
	"POPX", // Pop a stack value into register X
	"POPY", // Pop a stack value into register Y
	"POPP", // Pop the last value from the stack into the program counter

	// output
	"OUTV", // Send a value to the output register
	"OUTA", // Send the A register to the output register
	"OUTX", // Send the X register to the output register
	"OUTY", // Send the Y register to the output register
}

// opCode
func opCode(token string) (byte, error) {
	for k, v := range bitMap {
		if v == token {
			return byte(k), nil
		}
	}
	return 0, errors.Errorf("unknown token '%s'", token)
}
