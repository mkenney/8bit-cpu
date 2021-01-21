# data is a $label plus a byte:
$d1 0x1C   # hex 28
$d2 0b1110 # bin 14

    LDAV    $d1 # set register A to 28
    LDXV    $d2 # set register X to 14
    ADDX        # add register A (always) + register X, store result in register A (always)
    OUTA        # copy register A to the output register
