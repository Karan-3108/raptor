package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultRaptorInstanceCost is initially set the same as in wasmd
	DefaultRaptorInstanceCost uint64 = 60_000
	// DefaultRaptorCompileCost set to a large number for testing
	DefaultRaptorCompileCost uint64 = 100
)

// RaptorGasRegisterConfig is defaults plus a custom compile amount
func RaptorGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultRaptorInstanceCost
	gasConfig.CompileCost = DefaultRaptorCompileCost

	return gasConfig
}

func NewRaptorWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(RaptorGasRegisterConfig())
}
