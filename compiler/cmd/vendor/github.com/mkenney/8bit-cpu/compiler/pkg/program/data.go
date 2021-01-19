package program

import "github.com/bdlm/errors/v2"

// bitMap
var bitMap []string = []string{
	// system
	"HLT",
	"RST",
	"RUN",
	"NOP",
	"SLOP",

	// math
	"ADDV",
	"ADDX",
	"ADDY",
	"SUB",
	"SUBX",
	"SUBY",

	// branch
	"JMP",

	// data
	"LDAV",
	"LDAX",
	"LDAY",
	"LDXV",
	"LDXA",
	"LDXY",
	"LDYV",
	"LDYA",
	"LDYX",

	// output
	"OUTV",
	"OUTA",
	"OUTX",
	"OUTY",
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
