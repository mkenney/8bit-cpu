```ruby
# constants and subroutines must be defined before they are referenced.

# data is a $label plus a byte:
$d1 0x1C   # 28 - hex: 1c; bin: 11100 : 00 1c
$d2 0b1110 # 14 - hex: 0e; bin: 1110  : 00 0e

# subroutines are labels ending with a BRACE { and are delimited with a
# closing BRACE }
#
# `RUN <x>` should push the current PC value onto the call stack
# `}` should pop a value off the call stack and JMP to it

# initialize registers
reset {
    LDAV    0   # set register A to 0x00
    LDXV    0   # set register X to 0x00
    LDYV    0   # set register Y to 0x00
}

# calculate the next fibonacci number
nextfib {
    LDYA    # copy register A to register y
    ADDX    # add register A (always) + register X, store result in A (always)
    LDXY    # copy register Y to register X
}

# initialize. no label required.
    RUN     reset   # reset all data registers
    LDAV    1       # set register A to 0x01
    LDXV    0       # set register X to 0x01
    LDYV    1       # set register y to 0x01

# simple addition statement. the label is optional because it's never referenced
add
    LDAV $d1 # load 28 into register A (also the math register)
    ADDV $d2 # add 14 to the value in register A and store the result in register A
    OUTA     # load register A data into the output register

# this label is required, it is referenced in code below to create a loop
loop
    OUTA        # copy register A (rid 0) to the output register
    RUN nextfib # call Fibonacci subroutine
    JMP loop    # loop forever
```