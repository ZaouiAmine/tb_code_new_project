package lib

import (
    "fmt"
    "strconv"

    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
)

// Import `add` the library
//
//go:wasmimport libraries/tauhow_example_library add
func add(a, b uint32) uint64

func getQueryVarAsUint32(h http.Event, varName string) uint32 {
    varStr, err := h.Query().Get(varName)
    if err != nil {
        panic(err)
    }

    varUint, err := strconv.ParseUint(varStr, 10, 32)
    if err != nil {
        panic(err)
    }

    return uint32(varUint)
}

//export doAdd
func doAdd(e event.Event) uint32 {
    h, err := e.HTTP()
    if err != nil {
        return 1
    }

    // call the library function
    sum := add(getQueryVarAsUint32(h, "a"), getQueryVarAsUint32(h, "b"))

    // send the result over http
    h.Write([]byte(fmt.Sprintf("%d", sum)))

    return 0
}
