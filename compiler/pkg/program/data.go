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

	// internal
	"LABEL", //

	// math
	"ADDV", //
	"ADDX", //
	"ADDY", //
	"SUBV", //
	"SUBX", //
	"SUBY", //

	// branch
	"JMP", //

	// data
	"LDAV", //
	"LDAX", //
	"LDAY", //
	"LDXV", //
	"LDXA", //
	"LDXY", //
	"LDYV", //
	"LDYA", //
	"LDYX", //

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
