module github.com/mkenney/8bit-cpu/compiler/cmd

go 1.14

require (
	github.com/bdlm/log/v2 v2.0.3
	github.com/mkenney/8bit-cpu/cmp2/pkg v0.0.0-00010101000000-000000000000
)

replace github.com/mkenney/8bit-cpu/cmp2/pkg => ../pkg
