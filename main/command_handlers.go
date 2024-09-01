package main

import (
    "ripple/datagram_util"
    "ripple/handlers/trustlines/client_trustlines"
    "ripple/handlers/trustlines/server_trustlines"
    "ripple/handlers/payments/client_payments"
    "ripple/handlers/payments/server_payments"
)

// Define function signatures for client and server command handlers
type ClientCommandHandler func(session datagram_util.Session)
type ServerCommandHandler func(datagram *datagram_util.Datagram)

// CommandHandlers for clients
var clientHandlers = [128]ClientCommandHandler{
    0: client_trustlines.SetTrustline,
    1: client_trustlines.SyncTrustlineIn,
    2: client_trustlines.SyncTrustlineOut,
    3: client_trustlines.GetTrustlineIn,
    4: client_trustlines.GetTrustlineOut,
    5: client_payments.NewPaymentOut,
    6: client_payments.NewPaymentIn,
    7: client_payments.GetPayment,
    8: client_payments.StartPayment,
    // Other client handlers...
}

// CommandHandlers for servers
var serverHandlers = [128]ServerCommandHandler{
    0: server_trustlines.SetTrustline,
    1: server_trustlines.GetTrustline,
    2: server_payments.FindPathOut,
    3: server_payments.FindPathIn,
    4: server_payments.PathRecurse,
    // Other server handlers...
}
