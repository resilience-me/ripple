package commands

const (
    ClientTrustlines_SetTrustline      = 0
    ClientTrustlines_SyncTrustlineIn   = 1
    ClientTrustlines_SyncTrustlineOut  = 2
    ClientTrustlines_GetTrustlineIn    = 3
    ClientTrustlines_GetTrustlineOut   = 4
    ClientPayments_NewPaymentOut       = 5
    ClientPayments_NewPaymentIn        = 6
    ClientPayments_GetPayment          = 7
    ClientPayments_StartPayment        = 8

    ServerTrustlines_SetTrustline      = 127
    ServerTrustlines_GetTrustline      = 128
    ServerPayments_FindPathOut         = 129
    ServerPayments_FindPathIn          = 130
    ServerPayments_PathRecurse         = 131
)

// commandNames maps command bytes to their names
var commandNames = [256]string{
    0:   "ClientTrustlines_SetTrustline",
    1:   "ClientTrustlines_SyncTrustlineIn",
    2:   "ClientTrustlines_SyncTrustlineOut",
    3:   "ClientTrustlines_GetTrustlineIn",
    4:   "ClientTrustlines_GetTrustlineOut",
    5:   "ClientPayments_NewPaymentOut",
    6:   "ClientPayments_NewPaymentIn",
    7:   "ClientPayments_GetPayment",
    8:   "ClientPayments_StartPayment",

    127: "ServerTrustlines_SetTrustline",
    128: "ServerTrustlines_GetTrustline",
    129: "ServerPayments_FindPathOut",
    130: "ServerPayments_FindPathIn",
    131: "ServerPayments_PathRecurse",
    // Other indices are empty strings by default
}

// GetCommandName returns the name of the command for a given byte value.
func GetCommandName(command byte) string {
    return commandNames[command]
}
