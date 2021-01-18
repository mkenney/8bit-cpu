package program

type instructionType int

const (
	// Data stored on ROM
	TYPE_CONST instructionType = iota

	// Label for JMP instructions.
	TYPE_LABEL

	// Instruction
	TYPE_INSTR

	// Subroutine
	TYPE_SUB

	// Subroutine end
	TYPE_SUBEND
)

var callStack = []int{}
var subStack = []int{}

// data map
var datMap = map[string]int{}

// labels
var jmpMap = map[string]int{}

// step -> inst map
var stepMap = map[int]*Instruction{}

// subroutine map
var subMap = map[string]int{}

type bMap map[byte]string

var byt byte
