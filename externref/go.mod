module github.com/hybridgroup/mechanoid-examples/externref

go 1.22.0

replace github.com/tetratelabs/wazero => github.com/orsinium-forks/wazero v0.0.0-20240305131633-28fdf656fe85

require (
	github.com/hybridgroup/mechanoid v0.0.0-20240306161038-16a074177aa7
	github.com/orsinium-labs/wypes v0.1.4
)

require github.com/tetratelabs/wazero v1.6.0 // indirect
