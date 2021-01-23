package program

import (
	"github.com/bdlm/errors/v2"
)

type inst struct {
	// Instruction name as defined in `instructionSet`.
	Name string
	// Whether this instruction accepts param data.
	hasParam bool
}

var instTable []*inst = []*inst{
	// system
	&inst{Name: "HLT", hasParam: false},  // Halt system clock signal
	&inst{Name: "RST", hasParam: false},  // Reset all system registers
	&inst{Name: "NOP", hasParam: false},  // No-op, use 1 instruction cycle
	&inst{Name: "SLOP", hasParam: false}, // Slow no-op, use 16 instruction cycles

	// math
	&inst{Name: "ADDV", hasParam: true},  // Add a $const or literal to register A
	&inst{Name: "ADDX", hasParam: false}, // Add register X to register A
	&inst{Name: "ADDY", hasParam: false}, // Add register Y to register A

	&inst{Name: "SUBV", hasParam: true},  // Subtract a $const or literal from register A
	&inst{Name: "SUBX", hasParam: false}, // Subtract register X from register A
	&inst{Name: "SUBY", hasParam: false}, // Subtract register Y from register A

	// branching logic
	&inst{Name: "RUN", hasParam: true}, // Execute a subroutine. Is encoded as a JMPV instruction

	&inst{Name: "JMP", hasParam: true},   // `JMP [label]`  - Jump to [label]:       Load a label index into the program counter
	&inst{Name: "JMPV", hasParam: true},  // `JMPV [value]` - Jump to [value]:       Load a $const or literal value into the program counter
	&inst{Name: "JMPA", hasParam: false}, // `JMPA`         - Jump to register A:    Load register A into the program counter
	&inst{Name: "JMPX", hasParam: false}, // `JMPX`         - Jump to register X:    Load register X into the program counter
	&inst{Name: "JMPY", hasParam: false}, // `JMPY`         - Jump to register Y:    Load register Y into the program counter
	&inst{Name: "JMPS", hasParam: false}, // `JMPS`         - Jump to stack pointer: Load the last stack value into the program counter

	// data
	&inst{Name: "LDAV", hasParam: true},  // Load a $const or literal value into register A
	&inst{Name: "LDAX", hasParam: false}, // Load register X into register A
	&inst{Name: "LDAY", hasParam: false}, // Load register Y into register A

	&inst{Name: "LDXV", hasParam: true},  // Load a $const or literal value into register X
	&inst{Name: "LDXA", hasParam: false}, // Load register A into register X
	&inst{Name: "LDXY", hasParam: false}, // Load register Y into register X

	&inst{Name: "LDYV", hasParam: true},  // Load a $const or literal value into register Y
	&inst{Name: "LDYA", hasParam: false}, // Load register A into register Y
	&inst{Name: "LDYX", hasParam: false}, // Load register X into register Y

	// stack
	&inst{Name: "PSHV", hasParam: true},  // Push a $const or literal value onto the stack
	&inst{Name: "PSHA", hasParam: false}, // Push register A onto the stack
	&inst{Name: "PSHX", hasParam: false}, // Push register X onto the stack
	&inst{Name: "PSHY", hasParam: false}, // Push register Y onto the stack
	&inst{Name: "PSHP", hasParam: false}, // Push the current program counter onto the stack

	&inst{Name: "POPA", hasParam: false}, // Pop a stack value into register A
	&inst{Name: "POPX", hasParam: false}, // Pop a stack value into register X
	&inst{Name: "POPY", hasParam: false}, // Pop a stack value into register Y
	&inst{Name: "POPP", hasParam: false}, // Pop the last value from the stack into the program counter

	// output
	&inst{Name: "OUTV", hasParam: true},  // Send a value to the output register
	&inst{Name: "OUTA", hasParam: false}, // Send the A register to the output register
	&inst{Name: "OUTX", hasParam: false}, // Send the X register to the output register
	&inst{Name: "OUTY", hasParam: false}, // Send the Y register to the output register
}

func getInstruction(token string) (*inst, error) {
	for _, inst := range instTable {
		if token == inst.Name {
			return inst, nil
		}
	}
	return nil, errors.Errorf("unknown token '%s'", token)
}

// instructionCode
func instructionCode(token string) (byte, error) {
	for k, inst := range instTable {
		if token == inst.Name {
			return byte(k), nil
		}
	}
	return 0, errors.Errorf("unknown token '%s'", token)
}
