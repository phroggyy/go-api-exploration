package main

type hub struct {
    // Contains all connections
    connections map[*connection]bool
    // Contains our messages to be broadcast
    broadcast chan []byte
    // Contains the most recently registered connection
    register chan *connection
    // Contains the most recently unregistered connection
    unregister chan *connection
}

var h = hub{
    broadcast:   make(chan []byte),
    register:    make(chan *connection),
    unregister:  make(chan *connection),
    connections: make(map[*connection]bool),
}

func (h *hub) run() {
    for {
        // Wait until one of the following happens:
        // 1. A new connection is registered – add it to our connections
        // 2. A connection is unregistered – remove the connection
        // 3. A new message is broadcast – send it to all connections
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            if _, ok := h.connections[c]; ok {
                delete(h.connections, c)
                close(c.send)
            }
        case m := <-h.broadcast:
            for c := range h.connections {
                select {
                case c.send <- m:
                default:
                    // If we can't send the message, we remove the connection
                    close(c.send)
                    delete(h.connections, c)
                }
            }
        }
    }
}