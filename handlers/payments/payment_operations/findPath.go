package payment_operations

import (
    "log"
    "ripple/handlers/payments"
    "ripple/pathfinding"
    "ripple/types"
)

// FindPath handles the common logic for processing FindPath requests.
func FindPath(datagram *types.Datagram, inOrOut byte) {
    // Extract the path identifier and amount from datagram arguments
    pathIdentifier := types.BytesToArray32(datagram.Arguments[:32])
    pathAmount := types.BytesToUint32(datagram.Arguments[32:36])

    // Check if the trustline (incoming or outgoing) is sufficient for the path amount
    sufficient, err := CheckTrustlineSufficient(datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername, pathAmount, inOrOut)
    if err != nil {
        log.Printf("Error checking trustline: %v", err)
        return
    }
    if !sufficient {
        log.Printf("Insufficient trustline for user %s with peer %s at %s for amount: %d", datagram.Username, datagram.PeerUsername, datagram.PeerServerAddress, pathAmount)
        return
    }

    // Find the account using the username from the datagram
    account := pathfinding.GetPathManager().Find(datagram.Username)
    if account == nil {
        log.Printf("Account not found for user: %s", datagram.Username)
        return
    }

    // Retrieve the Path object using the identifier
    path := account.Find(pathIdentifier)
    if path == nil {
        // Path is not found, add the new path using the Add method
        newPeer := pathfinding.NewPeerAccount(datagram.PeerUsername, datagram.PeerServerAddress)
        if inOrOut == types.Outgoing {
            path = account.Add(pathIdentifier, pathAmount, newPeer, pathfinding.PeerAccount{})
        } else {
            path = account.Add(pathIdentifier, pathAmount, pathfinding.PeerAccount{}, newPeer)
        }
        log.Printf("Initialized new path for identifier: %x with amount: %d", pathIdentifier, pathAmount)

        // Send a PathFindingRecurse back to the appropriate peer
        PathRecurse(datagram, newPeer, 0)
        return
    }

    // Check if a Payment is already associated with this account and identifier
    if account.Payment != nil && account.Payment.Identifier == pathIdentifier {
        
    }
    if payments.CheckPathFound(path) {
        log.Printf("Path already found for path %x, ignoring find path command", pathIdentifier)
        return
    }

    pathDirection := payments.DeterminePathDirection(path)
    if pathDirection == payments.ReverseDirection(inOrOut) {
        log.Printf("Path found for %x", pathIdentifier)
    }

    // If the path is already present, forward the PathFinding request to peers
    log.Printf("Path already exists for identifier %x, forwarding to peers", pathIdentifier)
    ForwardFindPath(datagram, inOrOut)
}
