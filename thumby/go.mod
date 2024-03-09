module github.com/hybridgroup/mechanoid-examples/thumby

go 1.22.0

replace github.com/tetratelabs/wazero => github.com/orsinium-forks/wazero v0.0.0-20240305131633-28fdf656fe85

require (
	github.com/hybridgroup/mechanoid v0.0.0-20240309111213-758ddcc58e7a
	github.com/orsinium-labs/wypes v0.1.4
	tinygo.org/x/drivers v0.27.0
	tinygo.org/x/tinyfont v0.4.0
)

require (
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/hybridgroup/wasman v0.0.0-20240304140329-ce1ea6b61834 // indirect
	github.com/tetratelabs/wazero v1.6.0 // indirect
)
