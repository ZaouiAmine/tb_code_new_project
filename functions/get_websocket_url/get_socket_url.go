package lib

import (
    "crypto/md5"
    "encoding/hex"

    "github.com/taubyte/go-sdk/event"
    http "github.com/taubyte/go-sdk/http/event"
    pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

func fail(h http.Event, err error, code int) uint32 {
    h.Write([]byte(err.Error()))
    h.Return(code)
    return 1
}

//export getsocketurl
func getsocketurl(e event.Event) uint32 {
    h, err := e.HTTP()
    if err != nil {
        return 1
    }

    // get room from query
    room, err := h.Query().Get("room")
    if err != nil {
        return fail(h, err, 500)
    }

    // hash the room to create a channel name
    hash := md5.New()
    hash.Write([]byte(room))
    roomHash := hex.EncodeToString(hash.Sum(nil))

    // create/open a channel with the hash
    channel, err := pubsub.Channel("chat-" + roomHash)
    if err != nil {
        return fail(h, err, 500)
    }

    // get the websocket url
    url, err := channel.WebSocket().Url()
    if err != nil {
        return fail(h, err, 500)
    }

    // write the url to the response
    h.Write([]byte(url.Path))

    return 0
}