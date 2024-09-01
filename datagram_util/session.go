package datagram_util

import "net"

// Session is used to route datagrams and is passed to client handlers.
type Session struct {
    *Datagram
    Addr *net.UDPAddr
}
