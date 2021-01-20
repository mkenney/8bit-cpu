# data is a $label plus a byte:
$d1 0x1C   # hex 28
$d2 0b1110 # bin 14

    LDAV    $d1 # set register 0 to 28
    LDXV    $d2 # set register 1 to 14
    ADDX        # add register 0 (always) + register 1, store result in register 0 (always)
    OUTA        # copy register 0 to the output register
