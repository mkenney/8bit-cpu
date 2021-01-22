package program

import "github.com/bdlm/errors/v2"

// bitMap
var bitMap []string = []string{

	// system
	"HLT",  // Halt system clock signal
	"RST",  // Reset all system registers
	"RUN",  // Execute a subroutine
	"NOP",  // No-op, use 1 instruction cycle
	"SLOP", // Slow no-op, use 16 instruction cycles

	// math
	"ADDV", //
	"ADDX", //
	"ADDY", //
	"SUBV", //
	"SUBX", //
	"SUBY", //

	// branch
	"JMP",  // Load a label index into the program counter
	"JMPV", // Load a value into the program counter
	"JMPA", // Load register A into the program counter
	"JMPX", // Load register X into the program counter
	"JMPY", // Load register Y into the program counter
	"JMPS", // Load the last stack value into the program counter

	// data
	"LDAV", // Load a value into register A
	"LDAX", // Load register X into register A
	"LDAY", // Load register Y into register A

	"LDXV", // Load a value into register X
	"LDXA", // Load register A into register X
	"LDXY", // Load register Y into register X

	"LDYV", // Load a value into register Y
	"LDYA", // Load register A into register Y
	"LDYX", // Load register X into register Y

	// stack
	"PSHV", // Push a value onto the stack
	"PSHA", // Push register A onto the stack
	"PSHX", // Push register X onto the stack
	"PSHY", // Push register Y onto the stack

	"PSHP", // Push the current program counter onto the stack
	"POPP", // Pull the last value from the stack into the program counter

	"PULA", // Pull a stack value into register A
	"PULX", // Pull a stack value into register X
	"PULY", // Pull a stack value into register Y

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
