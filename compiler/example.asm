# data is a $label plus a byte:
$d1 0x1C   # hex 28
$d2 0b1110 # bin 14

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

    # todo...
}

# calculate the next fibonacci number
nextfib {
    LDYA    # copy register A to register y
    ADDX    # add register A (always) + register X, store result in A (always)
    LDXY    # copy register Y to register X
}

# initialize
setup
    RUN     reset   # reset all data registers
    LDAV    1       # set register A to 0x01
    LDXV    0       # set register X to 0x01
    LDYV    1       # set register y to 0x01

# loop
loop
    OUTA        # copy register A (rid 0) to the output register
    RUN nextfib # call Fibonacci subroutine
    JMP loop    # loop forever
