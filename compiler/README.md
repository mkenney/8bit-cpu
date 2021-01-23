```
$ cd cmd
$ go mod tidy
$ go mod vendor
$ go build -o ../bin/bcc .
$ cd ..
$ ./bin/bcc example.asm example.asm.img
$ hexdump -C example.asm.img
```
