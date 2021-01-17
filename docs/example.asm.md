# ASM

28 + 14
```ruby
# data is a $label plus a byte:
$d1: 0x1C   # hex 28
$d2: 0b1110 # bin 14

    LDA $d1 # load labeled data into register A (∑ register)
    ADD $d2 # add labeled data to register A (∑ register)
    JMP run # GOTO a label

# labels
run:
    # instructions
    LDA     $d1 # load 28 into register A
    LDB     $d2 # load 14 into register B
    ADD,reg B   # add register B to register A, store result in register A
    OUT,reg A   # load register A (∑ register) data into output register

callSubroutine:
    LDA $d1
    LDB $d2
    RUN addAndOutput

    # loop
    JMP callSubroutine

# subroutines
addAndOutput {
    ADD,reg B
    OUT,reg A
}
```
