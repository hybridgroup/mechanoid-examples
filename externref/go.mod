module github.com/hybridgroup/mechanoid-examples/externref

go 1.22.0

replace github.com/tetratelabs/wazero => github.com/orsinium-forks/wazero v0.0.0-20240217173836-b12c024bcbe4

require (
	github.com/hybridgroup/mechanoid v0.0.0-20240305142025-4530e84844d6
	github.com/orsinium-labs/wypes v0.1.1
)

require (
	github.com/hybridgroup/wasman v0.0.0-20240304140329-ce1ea6b61834 // indirect
	github.com/tetratelabs/wazero v1.6.0 // indirect
)
