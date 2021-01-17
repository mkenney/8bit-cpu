# Opcodes

## Clock
* HLT: halt
* RST: system reset

## Program Counter
* PCE: program counter enable
* PCR: program counter reset
* JMP: program counter in
* PCO: program counter out

## RAM
* RARR: memory address register reset
* RARI: RAM address register in
  * set the RAM address to BUS
* RARO: RAM address register out
  * output the RAM address to BUS
* RAMI: RAM data in
  * store data on BUS in RAM_O at address RAMR
* RAMO: RAM data out
  * put data stored on RAM_O at address RAR on BUS

## ROM
* RORR: memory address register reset
* RORI: memory address register in
  * set the current memory address to BUS
* RORO: memory address register out
  * output the current memory address to BUS
* ROMO: memory address register data out
  * put data stored on ROM_1 at address MAR on BUS

## Instruction Register
* IR: reset; IR = 0000
* IE: end;   IR = 1111
* II: in;    IR = ROM_0

## ∑ Arithmetec Logic Unit
* AE: enable; REG_A = BUS_0 + REG_A on CLK
  * do not set if ARI is set
* SUB: subtract flag

## Data Registers
### A Register
* ARR: reset; REG_A = 00000000
* ARI: in;    REG_A = BUS_0
  * do not set if AE is set
* ARO: out;   BUS_0 = REG_A

### [X|Y] Registers
* [X|Y] Register
* [X|Y]RR: reset; REG_[X|Y] = 00000000
* [X|Y]RI: in;    REG_[X|Y] = BUS_0
* [X|Y]RO: out;   BUS_0 = REG_[X|Y]

## Output
* OUT: in; OUT_0 = BUS_0
  * send OUT_0 to BUS_OUT at all times
