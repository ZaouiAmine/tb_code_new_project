package lib

import (
    "encoding/json"
    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
    "github.com/taubyte/go-sdk/storage"
)

func failed(h http.Event, err error, code int) uint32 {
    h.Write([]byte(err.Error()))
    h.Return(code)
    return 1
}

type Req struct {
    Filename string `json:"filename"`
    Data     string `json:"data"`
}

//export store
func store(e event.Event) uint32 {
    h, err := e.HTTP()
    if err != nil {
        return 1
    }

    // Open/Create the storage
    sto, err := storage.New("/simple/storage")
    if err != nil {
        return failed(h, err, 500)
    }

    // Read the request body
    reqDec := json.NewDecoder(h.Body())
    defer h.Body().Close()

    var req Req
    err = reqDec.Decode(&req)
    if err != nil {
        return failed(h, err, 500)
    }

    // Select file/object
    file := sto.File(req.Filename)

    // Write data to the file
    _, err = file.Add([]byte(req.Data), true)
    if err != nil {
        return failed(h, err, 500)
    }

    return 0
}