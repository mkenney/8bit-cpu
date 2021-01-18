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
    LD,0    0       # set register A to 0b00000000
    LD,1    0       # set register B to 0b00000000
    LD,2    0       # set register X to 0b00000000
    LD,3    0       # set register Y to 0b00000000

    # todo...
}

# calculate the next fibonacci number
nextfib {
    LDG,2   0       # copy register A (rid 0) to register X (rid 2)
    ADDG    1       # add register A (always) + B (rid 1), store result in A (always)
    LDG,1   2       # copy register X (rid 2) to register B (rid 1)
}

# initialize
setup:
    RUN     reset   # reset all data registers
    LD,0    1       # set register A to 0b00000001
    LD,2    1       # set register X to 0b00000001

# loop
loop:
    OUTG    0       # copy register A (rid 0) to the output register
    RUN     nextfib # call Fibonacci subroutine
    JMP     loop    # loop forever
```
