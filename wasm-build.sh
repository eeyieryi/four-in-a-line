#!/bin/sh

env GOOS=js GOARCH=wasm go build -o ./wasm/four-in-a-row.wasm .
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./wasm/