# initialize
    LDAV 1 # set register A to 0x01
    LDXV 0 # set register X to 0x00
    LDYV 1 # set register Y to 0x01

# loop
loop
    OUTA     # copy register A to the output register
    LDYA     # copy register A to register Y
    ADDX     # add register A (always) + register X, store result in register A (always)
    LDXY     # copy register Y to register X
    JMP loop # loop forever
