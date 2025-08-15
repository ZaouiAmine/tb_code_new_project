package lib

import (
    "io"

    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
    "github.com/taubyte/go-sdk/storage"
)

func failed(h http.Event, err error, code int) uint32 {
    h.Write([]byte(err.Error()))
    h.Return(code)
    return 1
}

//export get
func get(e event.Event) uint32 {
    h, err := e.HTTP()
    if err != nil {
        return 1
    }

    // Read the filename from the query string
    filename, err := h.Query().Get("filename")
    if err != nil {
        return failed(h, err, 400)
    }

    // Open/Create the storage
    sto, err := storage.New("/simple/storage")
    if err != nil {
        return failed(h, err, 500)
    }

    // Select file/object
    file := sto.File(filename)

    // Get a io.ReadCloser
    reader, err := file.GetFile()
    if err != nil {
        return failed(h, err, 500)
    }
    defer reader.Close()

    // Read from file and write to response
    _, err = io.Copy(h, reader)
    if err != nil {
        return failed(h, err, 500)
    }

    return 0
}