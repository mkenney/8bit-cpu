## Fibonacci

Fibonacci numbers output in an infinite loop using subroutine calls for initialization and to calculate the next number in each iteration.

```ruby
# data is a $label plus a byte:
$d1: 0x1C   # hex 28
$d2: 0b1110 # bin 14

# subroutines are labels ending with a BRACE { and are delimited with a
# closing BRACE }

# initialize registers
reset {
    LDA 0   # set register A to 0x00
    LDX 0   # set register X to 0x00
    LDY 0   # set register Y to 0x00

    # todo...
}

# calculate the next fibonacci number
nextfib {
    LDYA    # copy register A to register y
    ADDX    # add register A (always) + register X, store result in A (always)
    LDXY    # copy register Y to register X
}

# initialize
setup:
    RUN     reset   # reset all data registers

    LDA    1       # set register A to 0x01
    LDX    0       # set register X to 0x01
    LDY    1       # set register y to 0x01

# loop
loop:
    OUTA    # copy register A (rid 0) to the output register
    RUN nextfib # call Fibonacci subroutine
    JMP loop    # loop forever
```
