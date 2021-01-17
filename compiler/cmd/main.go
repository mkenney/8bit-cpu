package main

type prg [256][2]byte

var program prg

func main() {
	load()
}

func load() {
	program[0][0] = 0
}

var a = []string{
	"RUN",
	"HLT",
	"RST",
	"NOP",
	"SLOP",
	"LDA [BYTE]",
	"LDA,ram [BYTE]",
	"LDA,reg [BYTE]",
	"LDA,rom [BYTE]",
	"LDX [BYTE]",
	"LDX,ram [BYTE]",
	"LDX,reg [BYTE]",
	"LDX,rom [BYTE]",
	"LDY [BYTE]",
	"LDY,ram [BYTE]",
	"LDY,reg [BYTE]",
	"LDY,rom [BYTE]",
	"OUT [BYTE]",
	"OUT,ram [BYTE]",
	"OUT,reg [BYTE]",
	"OUT,rom [BYTE]",
	"PSH [BYTE]",
	"PSH,ram [BYTE]",
	"PSH,reg [BYTE]",
	"PSH,rom [BYTE]",
	"PUL",
	"PUL,ram [BYTE]",
	"PUL,reg [BYTE]",
	"STA [BYTE]",
	"STB [BYTE]",
	"STX [BYTE]",
	"STY [BYTE]",
	"STM [BYTE]",
	"STR [BYTE]",
	"ADD [BYTE]",
	"ADD,ram [BYTE]",
	"ADD,reg [BYTE]",
	"ADD,rom [BYTE]",
	"SUB [BYTE]",
	"SUB,ram [BYTE]",
	"SUB,reg [BYTE]",
	"SUB,rom [BYTE]",
	"JMP [BYTE]",
	"JMP,ram [BYTE]",
	"JMP,reg [BYTE]",
	"JMP,rom [BYTE]",
}
