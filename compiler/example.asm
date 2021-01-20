# subroutine start label
name {

# subroutine end label
}

# data label $var [VAL]
# var is an address
# val is a byte
$var 1

# jump label
loop










# subroutine to calculate the next fibonacci number
#nextfib {
#    LDG,2   0       # copy register 0 to register 2
#    ADDG    1       # add register 0 (always) + register 1, store result in register 0 (always)
#    LDG,1   2       # copy register 2 to register 1
#}

# initialize
setup
    LD,0    1       # set register 0 to 0x01
    LD,1    0       # set register 1 to 0x00
    LD,2    1       # set register 2 to 0x01

# loop
loop
    OUTG    0       # copy register 0 to the output register
    # RUN nextfib     # will go here when there is a call stack
    LDG,2   0       # copy register 0 to register 2
    ADDG    1       # add register 0 (always) + register 1, store result in register 0 (always)
    LDG,1   2       # copy register 2 to register 1
    JMP     loop    # loop forever
