module github.com/mkenney/8bit-cpu/compiler/cmd

go 1.14

require (
	github.com/bdlm/log/v2 v2.0.3
	github.com/mkenney/8bit-cpu/compiler/pkg v0.0.0-20210118233019-84f1ef6c4e3b
)

replace github.com/mkenney/8bit-cpu/compiler/pkg => ../pkg
