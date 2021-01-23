package program

// init populates opMap.
func init() {
	for pcid, op := range opTable {
		op.pcid = byte(pcid)
		opMap[op.name] = op
	}
}

type op struct {
	// Operation name as defined in `opTable`.
	name string
	// Whether this operation accepts param data.
	hasParam bool
	// Program counter Id.
	pcid byte
}

var opMap map[string]*op

var opTable []*op = []*op{
	// system
	&op{name: "HLT", hasParam: false},  // Halt system clock signal
	&op{name: "RST", hasParam: false},  // Reset all system registers
	&op{name: "NOP", hasParam: false},  // No-op, use 1 instruction cycle
	&op{name: "SLOP", hasParam: false}, // Slow no-op, use 16 instruction cycles

	// math
	&op{name: "ADDV", hasParam: true},  // Add a $const or literal to register A
	&op{name: "ADDX", hasParam: false}, // Add register X to register A
	&op{name: "ADDY", hasParam: false}, // Add register Y to register A

	&op{name: "SUBV", hasParam: true},  // Subtract a $const or literal from register A
	&op{name: "SUBX", hasParam: false}, // Subtract register X from register A
	&op{name: "SUBY", hasParam: false}, // Subtract register Y from register A

	// branching logic
	&op{name: "RUN", hasParam: true}, // Execute a subroutine. Is encoded as a JMPV operation

	&op{name: "JMP", hasParam: true},   // `JMP [label]`  - Jump to [label]:       Load a label index into the program counter
	&op{name: "JMPV", hasParam: true},  // `JMPV [value]` - Jump to [value]:       Load a $const or literal value into the program counter
	&op{name: "JMPA", hasParam: false}, // `JMPA`         - Jump to register A:    Load register A into the program counter
	&op{name: "JMPX", hasParam: false}, // `JMPX`         - Jump to register X:    Load register X into the program counter
	&op{name: "JMPY", hasParam: false}, // `JMPY`         - Jump to register Y:    Load register Y into the program counter
	&op{name: "JMPS", hasParam: false}, // `JMPS`         - Jump to stack pointer: Load the last stack value into the program counter

	// data
	&op{name: "LDAV", hasParam: true},  // Load a $const or literal value into register A
	&op{name: "LDAX", hasParam: false}, // Load register X into register A
	&op{name: "LDAY", hasParam: false}, // Load register Y into register A

	&op{name: "LDXV", hasParam: true},  // Load a $const or literal value into register X
	&op{name: "LDXA", hasParam: false}, // Load register A into register X
	&op{name: "LDXY", hasParam: false}, // Load register Y into register X

	&op{name: "LDYV", hasParam: true},  // Load a $const or literal value into register Y
	&op{name: "LDYA", hasParam: false}, // Load register A into register Y
	&op{name: "LDYX", hasParam: false}, // Load register X into register Y

	// stack
	&op{name: "PSHV", hasParam: true},  // Push a $const or literal value onto the stack
	&op{name: "PSHA", hasParam: false}, // Push register A onto the stack
	&op{name: "PSHX", hasParam: false}, // Push register X onto the stack
	&op{name: "PSHY", hasParam: false}, // Push register Y onto the stack
	&op{name: "PSHP", hasParam: false}, // Push the current program counter onto the stack

	&op{name: "POPA", hasParam: false}, // Pop a stack value into register A
	&op{name: "POPX", hasParam: false}, // Pop a stack value into register X
	&op{name: "POPY", hasParam: false}, // Pop a stack value into register Y
	&op{name: "POPP", hasParam: false}, // Pop the last value from the stack into the program counter

	// output
	&op{name: "OUTV", hasParam: true},  // Send a value to the output register
	&op{name: "OUTA", hasParam: false}, // Send the A register to the output register
	&op{name: "OUTX", hasParam: false}, // Send the X register to the output register
	&op{name: "OUTY", hasParam: false}, // Send the Y register to the output register
}
