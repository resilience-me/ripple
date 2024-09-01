package datagram_util

import (
    "net"
)

// Session represents a session and includes an embedded Datagram.
type Session struct {
    *Datagram          // Embedding Datagram directly
    Addr      *net.UDPAddr // The UDP address associated with this session
}
