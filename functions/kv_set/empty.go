package lib

import (
    "encoding/json"

    "github.com/taubyte/go-sdk/database"
    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
)

func fail(h http.Event, err error, code int) uint32 {
    h.Write([]byte(err.Error()))
    h.Return(code)
    return 1
}

type Req struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

//export set
func set(e event.Event) uint32 {
    h, err := e.HTTP()
    if err != nil {
        return 1
    }

    // (Create) & Open the database
    db, err := database.New("/example/kv")
    if err != nil {
        return fail(h, err, 500)
    }

    // Decode the request body
    reqDec := json.NewDecoder(h.Body())
    defer h.Body().Close()

    // Decode the request body
    var req Req
    err = reqDec.Decode(&req)
    if err != nil {
        return fail(h, err, 500)
    }

    // Put the key/value into the database
    err = db.Put(req.Key, []byte(req.Value))
    if err != nil {
        return fail(h, err, 500)
    }

    return 0
}