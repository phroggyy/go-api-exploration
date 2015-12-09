package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "time"
)

const (
    // The max time allowed to write a message
    writeWait = 10 * time.Second

    // Max time to read a message
    pongWait = 60 * time.Second

    // Ping the peer with this period. Must be < pongWait
    pingPeriod = pongWait * 9 / 10

    maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type connection struct {
    ws *websocket.Conn

    // Buffered channel of outbound messages
    send chan []byte
}

func (c *connection) readPump() {
    defer func() {
        h.unregister <- c
        c.ws.Close()
    }()
    c.ws.SetReadLimit(maxMessageSize)
    c.ws.SetReadDeadline(time.Now().Add(pongWait))
    c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
    for {
        _, message, err := c.ws.ReadMessage()
        if err != nil {
            break
        }
        h.broadcast <- message
    }
}

func (c *connection) write(mt int, payload []byte) error {
    c.ws.SetWriteDeadline(time.Now().Add(writeWait))
    return c.ws.WriteMessage(mt, payload)
}

func (c *connection) writePump() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        c.ws.Close()
    }()
    for {
        select {
        case message, ok := <-c.send:
            if !ok {
                // No more messages – write a close and then finish the 
                c.write(websocket.CloseMessage, []byte{})
                return
            }
            if err := c.write(websocket.TextMessage, message); err != nil {
                return
            }
        case <-ticker.C:
            if err := c.write(websocket.PingMessage, []byte{}); err != nil {
                return
            }
        }
    }
}

func (ctl *Controller) Stream(context *gin.Context) {
    ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }

    c := &connection{send: make(chan []byte, 256), ws: ws}
    h.register <- c
    go c.writePump()
    c.readPump()
}