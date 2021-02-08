package main

import (
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	registerHandlers()

	go generatePasswords()

	wasmhttp.Serve(nil)

	select {}
}

func saveCounter() {}
