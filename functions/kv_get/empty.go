package lib

import (
    "github.com/taubyte/go-sdk/database"
    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
)

func fail(h http.Event, err error, code int) uint32 {
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

    key, err := h.Query().Get("key")
    if err != nil {
        return fail(h, err, 400)
    }

    db, err := database.New("/example/kv")
    if err != nil {
        return fail(h, err, 500)
    }

    value, err := db.Get(key)
    if err != nil {
        return fail(h, err, 500)
    }

    h.Write(value)
    h.Return(200)

    return 0
}